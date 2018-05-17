package fraudclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Response struct
type response struct {
	Status string `json:"status"`
}

// SendList sends a blackList to Kochava
func SendList(logger *log.Logger, list BlackList, api string, debug bool, action string) error {

	for i := range list.BlackListDevices {

		reqBody, err := json.Marshal(list.BlackListDevices[i])
		if err != nil {
			logger.Println(err)
			return err
		}
		if debug == false {
			err = send(logger, reqBody, api, action)
			if err != nil {
				logger.Println(err)
				return err
			}
		} else {
			logger.Printf("DEBUG: %#v\n", reqBody)
		}
	}

	for i := range list.BlackListSiteIDs {

		reqBody, err := json.Marshal(list.BlackListSiteIDs[i])
		if err != nil {
			logger.Println(err)
			return err
		}
		if debug == false {
			err = send(logger, reqBody, api, action)
			if err != nil {
				logger.Println(err)
				return err
			}
		} else {
			logger.Printf("DEBUG: %#v\n", reqBody)
		}

	}

	for i := range list.BlackListIPs {

		reqBody, err := json.Marshal(list.BlackListIPs[i])
		if err != nil {
			logger.Println(err)
			return err
		}

		if debug == false {
			err = send(logger, reqBody, api, action)
			if err != nil {
				logger.Println(err)
				return err
			}
		} else {
			logger.Printf("DEBUG: %#v\n", reqBody)
		}
	}
	return nil

}

func send(logger *log.Logger, reqBody []byte, api string, action string) error {

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
		logger.Println(err)
		return err
	}

	method := "POST"

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		logger.Println(err)
		return err
	}

	req.Header.Add("Authentication-Key", api)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		logger.Println(err)
		return err
	}

	body, _ := ioutil.ReadAll(res.Body)
	var status response
	err = json.Unmarshal(body, &status)
	if err != nil {
		logger.Println(err)
		return err
	}

	switch res.StatusCode {
	case 404:
		logger.Println(status.Status)
		return errors.New(status.Status)
	case 403:
		if action == "addupdate" {
			send(logger, reqBody, api, "update")
			logger.Println("entry already found; updating instead")
			return nil
		}
		logger.Println(status.Status)
		return errors.New(status.Status)
	case 200:
		return nil
	default:
		logger.Println(status.Status)
		return errors.New(status.Status)
	}

}
