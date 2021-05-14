package client

import (
	"fmt"
	"net/http"

	"github.com/CiscoDevNet/sdwan-go-client/container"
	"github.com/CiscoDevNet/sdwan-go-client/models"
)

//GetViaURL used for GET API call for objects
func (client *Client) GetViaURL(endpoint string) (*container.Container, error) {
	req, err := client.MakeRequest("GET", endpoint, nil, nil, true)
	if err != nil {
		return nil, err
	}

	cont, resp, err := client.Do(req)
	if err != nil && err == fmt.Errorf("Invalid Session") {
		client.authClient.valid = false
		cont, resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}
	}

	if cont == nil {
		return nil, fmt.Errorf("Empty Response Body")
	}

	return cont, checkForErrors(cont, resp)
}

//Save used for POST API call for objects
func (client *Client) Save(endpoint string, obj models.Model) (*container.Container, error) {
	jsonPayload, err := client.prepareModel(obj)
	if err != nil {
		return nil, err
	}

	req, err := client.MakeRequest("POST", endpoint, jsonPayload, nil, true)
	if err != nil {
		return nil, err
	}

	cont, resp, err := client.Do(req)
	if err != nil && err == fmt.Errorf("Invalid Session") {
		client.authClient.valid = false
		cont, resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}
	}
	return cont, checkForErrors(cont, resp)
}

//Update used for PUT API call for objects
func (client *Client) Update(endpoint string, obj models.Model) (*container.Container, error) {
	jsonPayload, err := client.prepareModel(obj)
	if err != nil {
		return nil, err
	}

	req, err := client.MakeRequest("PUT", endpoint, jsonPayload, nil, true)
	if err != nil {
		return nil, err
	}

	cont, resp, err := client.Do(req)
	if err != nil && err == fmt.Errorf("Invalid Session") {
		client.authClient.valid = false
		cont, resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}
	}
	return cont, checkForErrors(cont, resp)
}

//Delete used for DELETE API call for objects
func (client *Client) Delete(endpoint string) (*container.Container, error) {
	req, err := client.MakeRequest("DELETE", endpoint, nil, nil, true)
	if err != nil {
		return nil, err
	}

	cont, resp, err := client.Do(req)
	if err != nil && err == fmt.Errorf("Invalid Session") {
		client.authClient.valid = false
		cont, resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}
	}
	return cont, checkForErrors(cont, resp)
}

func (client *Client) prepareModel(obj models.Model) (*container.Container, error) {
	con, err := obj.ToMap()
	if err != nil {
		return nil, err
	}

	payload := &container.Container{}

	for key, value := range con {
		payload.Set(value, key)
	}
	return payload, nil
}

func checkForErrors(cont *container.Container, resp *http.Response) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	msg := cont.S("error", "message").String()
	details := cont.S("error", "details").String()

	return fmt.Errorf("Status: %v, messege: %s, details: %s", resp.StatusCode, msg, details)
}
