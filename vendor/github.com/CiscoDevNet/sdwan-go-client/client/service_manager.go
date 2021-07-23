package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/CiscoDevNet/sdwan-go-client/container"
	"github.com/CiscoDevNet/sdwan-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
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

//SaveviaBytes used for POST API call for bodyBytes
func (client *Client) SaveviaBytes(endpoint string, body []byte) (*container.Container, error) {
	req, err := client.MakeRequest("POST", endpoint, nil, body, true)
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

var contentTypeforFileUpload string

//SaveWithFile used for POST API call for file upload
func (client *Client) SaveWithFile(endpoint string, obj models.Model) (*container.Container, error) {
	bytesPayload, err := client.prepareModelWithFile(obj)
	if err != nil {
		return nil, err
	}

	req, err := client.MakeRequest("POST", endpoint, nil, bytesPayload, true)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentTypeforFileUpload)

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

//UpdateviaBytes used for PUT API call for bodyBytes
func (client *Client) UpdateviaBytes(endpoint string, body []byte) (*container.Container, error) {
	req, err := client.MakeRequest("PUT", endpoint, nil, body, true)
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

func (client *Client) prepareModelWithFile(obj models.Model) ([]byte, error) {
	con, err := obj.ToMap()
	if err != nil {
		return nil, err
	}

	jsonData, err := structure.FlattenJsonToString(con["data"].(map[string]interface{}))
	if err != nil {
		return nil, err
	}

	payload, err := newfileUploadRequest(jsonData, "file", con["file"].(string))
	if err != nil {
		return nil, err
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

func newfileUploadRequest(data string, keyName string, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(keyName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	_ = writer.WriteField("data", data)

	contentTypeforFileUpload = writer.FormDataContentType()

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return body.Bytes(), nil
}
