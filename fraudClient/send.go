package fraudclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Response struct
type blacklistResponse struct {
	Status string `json:"status"`
}

// SiteIDs adds siteids to the account blacklist
func (a Add) SiteIDs(sites ...siteID) error {

	for i := range sites {
		reqBody, err := json.Marshal(sites[i])
		if err != nil {
			return err
		}
		if a.debug == false {
			err = send(reqBody, a.APIKey, "add")
			if err != nil {
				return err
			}
		} else {
			return nil
		}
	}
	return nil
}

// DeviceIDs adds device ids to the account blacklist
func (a Add) DeviceIDs(devices ...deviceID) error {
	for i := range devices {
		reqBody, err := json.Marshal(devices[i])
		if err != nil {
			return err
		}
		if a.debug == false {
			err = send(reqBody, a.APIKey, "add")
			if err != nil {
				return err
			}
		} else {
			return nil
		}
	}
	return nil
}

// IPAddresses adds ips to the account blacklist
func (a Add) IPAddresses(ips ...ipAddress) error {
	for i := range ips {
		reqBody, err := json.Marshal(ips[i])
		if err != nil {
			return err
		}
		if a.debug == false {
			err = send(reqBody, a.APIKey, "add")
			if err != nil {
				return err
			}
		} else {
			return nil
		}
	}
	return nil
}

func send(reqBody []byte, api string, action string) error {

	// Slow it down so it doesn't hit the API too quickly
	time.Sleep(50 * time.Millisecond)

	// Kochava Fraud Endpoint to Hit

	var endpoint string

	switch {
	case action == "add" || action == "addupdate":
		endpoint = "https://fraud.api.kochava.com/fraud/blacklist/add"
	case action == "update":
		endpoint = "https://fraud.api.kochava.com/fraud/blacklist/update"
	case action == "remove":
		endpoint = "https://fraud.api.kochava.com/fraud/blacklist/remove"
	default:
		err := fmt.Errorf("switch case for action in send() should never reach default; action %s invalid", action)

		return err
	}

	method := "POST"

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(reqBody))
	if err != nil {

		return err
	}

	req.Header.Add("Authentication-Key", api)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {

		return err
	}

	body, _ := ioutil.ReadAll(res.Body)
	var status blacklistResponse
	err = json.Unmarshal(body, &status)
	if err != nil {

		return err
	}

	switch res.StatusCode {
	case 404:

		return errors.New(status.Status)
	case 403:
		// if action == "addupdate" {
		// 	send(logger, reqBody, api, "update")
		// 	logger.Println("entry already found; updating instead")
		// 	return nil
		// }

		return errors.New(status.Status)
	case 200:
		return nil
	default:

		return errors.New(status.Status)
	}

}
