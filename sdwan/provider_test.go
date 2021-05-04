package sdwan

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"sdwan": testAccProvider,
	}
}
func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	// We will use this function later on to make sure our test environment is valid.
	// For example, you can make sure here that some environment variables are set.
	if v := os.Getenv("SDWAN_USERNAME"); v == "" {
		t.Fatal("SDWAN_USERNAME env variable must be set for acceptance tests")
	}

	if v := os.Getenv("SDWAN_PASSWORD"); v == "" {
		t.Fatal("Either of SDWAN_PASSWORD env variables must be set for acceptance tests")
	}

	if v := os.Getenv("SDWAN_URL"); v == "" {
		t.Fatal("SDWAN_URL env variable must be set for acceptance tests")
	}
}
