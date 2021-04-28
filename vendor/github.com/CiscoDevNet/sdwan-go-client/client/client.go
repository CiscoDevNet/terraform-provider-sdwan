package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

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
	rateLimiter *rate.Limiter
}

var clientImpl *Client

func (client *Client) configProxy(transport *http.Transport) *http.Transport {
	proxy, err := url.Parse(client.proxyURL)
	if err != nil {
		log.Fatal(err)
	}
	transport.Proxy = http.ProxyURL(proxy)
	return transport

}

func (client *Client) useInsecureHTTPClient() *http.Transport {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: client.insecure,
		},
	}

	return transport
}

func initClient(clientURL, username, password, proxyURL string, rateLimit int, insecure bool) *Client {
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
		rateLimiter: rate.NewLimiter(rate.Limit(float64(rateLimit)/float64(1)), 1),
		httpClient:  http.DefaultClient,
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
//e.g. GetClient("https://my-company.com:port", "userName", "password", "https://my-company-proxy.com:port", true)
func GetClient(clientURL, username, password, proxy string, rateLimit int, insecure bool) *Client {
	if clientImpl == nil {
		clientImpl = initClient(clientURL, username, password, proxy, rateLimit, insecure)
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
	method := "POST"
	path := "/j_security_check"
	body := []byte(fmt.Sprintf("j_username=%s&j_password=%s", client.username, client.password))

	client.skipLogging = true

	req, err := client.MakeRequest(method, path, nil, body, false)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, resp, err := client.Do(req)
	if err != nil {
		return err
	}

	cookies := resp.Header.Get("Set-Cookie")
	if cookies == "" {
		return fmt.Errorf("No valid JSESSION ID returned")
	}
	jsessionID := strings.Split(cookies, ";")[0]

	method = "GET"
	path = "/dataservice/client/token"
	req, err = client.MakeRequest(method, path, nil, nil, false)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", fmt.Sprintf("JESSIONID=%s", jsessionID))

	_, resp, err = client.Do(req)
	if err != nil {
		return err
	}

	var token string
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		token = ""
	} else {
		token = string(bodyBytes)
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
	if !client.skipLogging {
		log.Println("[DEBUG] HTTP Request ", req.Method, req.URL.String())
		log.Println("[DEBUG] HTTP Response ", resp.StatusCode, resp)
	}

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
	if !client.skipLogging {
		log.Println("[DEBUG] HTTP Response unique string ", req.Method, req.URL.String(), bodystrings)
	}

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
	} else if resp.StatusCode == 405 {
		return fmt.Errorf("Invalid Session")
	}
	return nil
}
