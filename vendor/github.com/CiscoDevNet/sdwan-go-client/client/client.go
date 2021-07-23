package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/CiscoDevNet/sdwan-go-client/container"
	"golang.org/x/time/rate"
)

//Client structure is the singleton implementation of the client
type Client struct {
	baseURL     *url.URL
	httpClient  *http.Client
	authClient  *auth
	username    string
	password    string
	insecure    bool
	skipLogging bool
	proxyURL    string
	proxyCreds  string
	rateLimiter *rate.Limiter
	mutex       *sync.Mutex
}

var clientImpl *Client

func (client *Client) configProxy(transport *http.Transport) *http.Transport {
	proxy, err := url.Parse(client.proxyURL)
	if err != nil {
		log.Fatal(err)
	}
	transport.Proxy = http.ProxyURL(proxy)

	if client.proxyCreds != "" {
		basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(client.proxyCreds))
		transport.ProxyConnectHeader = http.Header{}
		transport.ProxyConnectHeader.Add("Proxy-Authorization", basicAuth)
	}

	return transport
}

func (client *Client) useInsecureHTTPClient() *http.Transport {
	transport := http.DefaultTransport.(*http.Transport)

	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: client.insecure,
	}
	// transport := &http.Transport{
	// 	TLSClientConfig: &tls.Config{
	// 		InsecureSkipVerify: client.insecure,
	// 	},
	// }

	return transport
}

func initClient(clientURL, username, password, proxyURL, proxyCreds string, rateLimit int, insecure bool) *Client {
	baseURL, err := url.Parse(clientURL)
	if err != nil {
		log.Fatal(err)
	}

	client := &Client{
		baseURL:     baseURL,
		username:    username,
		password:    password,
		insecure:    insecure,
		proxyURL:    proxyURL,
		proxyCreds:  proxyCreds,
		rateLimiter: rate.NewLimiter(rate.Limit(float64(rateLimit)/float64(1)), 1),
		httpClient:  http.DefaultClient,
		authClient:  &auth{},
		mutex:       &sync.Mutex{},
	}

	transport := client.useInsecureHTTPClient()

	if client.proxyURL != "" {
		transport = client.configProxy(transport)
	}

	client.httpClient = &http.Client{
		Transport: transport,
	}

	return client
}

//GetClient constructor for Client implementation
//
//e.g. GetClient("https://my-company.com:port", "userName", "password", "https://my-company-proxy.com:port", "username:password", 500, true)
func GetClient(clientURL, username, password, proxy, proxyCreds string, rateLimit int, insecure bool) *Client {
	if clientImpl == nil {
		clientImpl = initClient(clientURL, username, password, proxy, proxyCreds, rateLimit, insecure)
	}
	return clientImpl
}

//MakeRequest generate the HTTP request with required headers
func (client *Client) MakeRequest(method, path string, body *container.Container, bodyBytes []byte, authenticated bool) (*http.Request, error) {
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	reqURL := client.baseURL.ResolveReference(url)

	var req *http.Request
	if body != nil {
		req, err = http.NewRequest(method, reqURL.String(), bytes.NewBuffer(body.Bytes()))
	} else if bodyBytes != nil {
		req, err = http.NewRequest(method, reqURL.String(), bytes.NewBuffer(bodyBytes))
	} else {
		req, err = http.NewRequest(method, reqURL.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	if client.skipLogging {
		log.Println("HTTP request ", method, path)
	} else {
		log.Println("HTTP request ", method, path, req)
	}

	if authenticated {
		req, err = client.injectAuthenticationHeader(req, path)
		if err != nil {
			return req, err
		}
		log.Println("HTTP request after injection ", method, path, req)
	}
	return req, nil
}

func (client *Client) authenticate() error {
	client.skipLogging = true

	jsessionID, err := client.getSessionID()
	if err != nil {
		return err
	}

	token, err := client.getToken(jsessionID)
	if err != nil {
		return err
	}

	client.authClient.jsessionID = jsessionID
	client.authClient.token = token
	client.authClient.valid = true

	client.skipLogging = false
	return nil
}

//Do performs the appropriate API call to the server and respond back with
//
//processed API response
func (client *Client) Do(req *http.Request) (*container.Container, *http.Response, error) {
	log.Println("[DEBUG] Begining Do method ", req.URL.String())

	err := client.rateLimiter.Wait(context.Background())
	if err != nil {
		return nil, nil, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	log.Println("[DEBUG] HTTP Request ", req.Method, req.URL.String())
	log.Println("[DEBUG] HTTP Response ", resp.StatusCode, resp)

	err = checkResponse(resp)
	if err != nil {
		return nil, resp, err
	}

	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	bodystrings := string(bodybytes)
	// resp.Body.Close()
	log.Println("[DEBUG] HTTP Response unique string ", req.Method, req.URL.String(), bodystrings)

	obj, err := container.ParseJSON(bodybytes)
	if err != nil && resp.StatusCode != 200 {
		return nil, resp, fmt.Errorf(bodystrings)
	}

	log.Println("[DEBUG] Ending Do method ", req.URL.String())
	return obj, resp, nil
}

func checkResponse(resp *http.Response) error {
	if resp.StatusCode == 200 {
		if resp.Header.Get("Content-Type") == "text/html;charset=UTF-8" {
			return fmt.Errorf("Invalid Session")
		}
		return nil
	}
	return nil
}

func (client *Client) getSessionID() (string, error) {
	method := "POST"
	path := "/j_security_check"
	body := []byte(fmt.Sprintf("j_username=%s&j_password=%s", client.username, client.password))

	req, err := client.MakeRequest(method, path, nil, body, false)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	err = checkSessionResponse(resp)
	if err != nil {
		return "", err
	}

	cookies := resp.Header.Get("Set-Cookie")
	if cookies == "" {
		return "", fmt.Errorf("No valid JSESSION ID returned")
	}
	session := strings.Split(cookies, ";")[0]
	jsessionID := strings.Split(session, "=")[1]

	return jsessionID, nil
}

func (client *Client) getToken(jsessionID string) (string, error) {
	method := "GET"
	path := "/dataservice/client/token"
	req, err := client.MakeRequest(method, path, nil, nil, false)
	if err != nil {
		return "", err
	}
	req.Header.Set("Cookie", fmt.Sprintf("JSESSIONID=%s;", jsessionID))
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	err = checkTokenResponse(resp, jsessionID)
	if err != nil {
		return "", err
	}

	var token string
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		token = ""
	} else {
		token = string(bodyBytes)
	}

	return token, nil
}

func checkTokenResponse(resp *http.Response, jsessionID string) error {
	if resp.StatusCode == 200 {
		if resp.Header.Get("Content-Type") == "text/html;charset=UTF-8" {
			return fmt.Errorf("User name or password was invalid or If username and password are valid user account is locked Wait for some time and try again or contact Administrator or If username and password are valid, password has expired")
		}
	}
	return nil
}

func checkSessionResponse(resp *http.Response) error {
	if resp.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("Invalid Session")
}
