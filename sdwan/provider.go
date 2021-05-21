package sdwan

import (
	"fmt"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

//Provider is the main entry point for the SDWAN Terraform Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SDWAN_USERNAME", nil),
				Description: "Username for the SDWAN vManage server",
			},

			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SDWAN_PASSWORD", nil),
				Description: "Password for the SDWAN vManage server",
			},

			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SDWAN_URL", nil),
				Description: "URL for the SDWAN vManage server",
			},

			"proxy_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SDWAN_PROXY_URL", nil),
				Description: "Proxy Server URL with port number",
			},

			"proxy_creds": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SDWAN_PROXY_CREDS", nil),
				Description: "Proxy Server Credentials in the form of username:password",
			},

			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Allow insecure HTTPS client",
			},

			"rate_limit": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "API rate limit per second",
				Default:      500,
				ValidateFunc: validation.IntBetween(0, 500),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"sdwan_device_template":         resourceSDWANDeviceTemplate(),
			"sdwan_system_feature_template": resourceSDWANSystemFeatureTemplate(),
			"sdwan_ntp_feature_template":    resourceSDWANNtpFeatureTemplate(),
			"sdwan_vpn_feature_template":    resouceSDWANVPNFeatureTemplate(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"sdwan_device_template":         datasourceSDWANDeviceTemplate(),
			"sdwan_system_feature_template": datasourceSDWANSystemFeatureTemplate(),
			"sdwan_ntp_feature_template":    datasourceSDWANNtpFeatureTemplate(),
			"sdwan_vpn_feature_template":    datasourceSDWANVPNFeatureTemplate(),
		},

		ConfigureFunc: configClient,
	}
}

func configClient(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		Username:   d.Get("username").(string),
		Password:   d.Get("password").(string),
		URL:        d.Get("url").(string),
		IsInsecure: d.Get("insecure").(bool),
		ProxyURL:   d.Get("proxy_url").(string),
		ProxyCreds: d.Get("proxy_creds").(string),
		RateLimit:  d.Get("rate_limit").(int),
	}

	if err := config.Valid(); err != nil {
		return nil, err
	}

	return config.getClient(), nil
}

//Valid is responsible for the validation of the required parameter of terraform block schema
func (c Config) Valid() error {

	if c.Username == "" {
		return fmt.Errorf("Username must be provided for the SDWAN provider")
	}

	if c.Password == "" {
		return fmt.Errorf("Password must be provided for the SDWAN provider")
	}

	if c.URL == "" {
		return fmt.Errorf("The URL must be provided for the SDWAN provider")
	}

	return nil
}

func (c Config) getClient() interface{} {
	return client.GetClient(c.URL, c.Username, c.Password, c.ProxyURL, c.ProxyCreds, c.RateLimit, c.IsInsecure)
}

//Config the basic structure used for the management of the Terraform block schema attributes
type Config struct {
	Username   string
	Password   string
	URL        string
	IsInsecure bool
	ProxyURL   string
	ProxyCreds string
	RateLimit  int
}
