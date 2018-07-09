package fraudclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendRequest(accountID, view, startDate, endDate, format, fraudType, endpoint, authKey string, filters []FilterObject) (interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Send Request One: ", format)
		}
	}()

	var req request
	req.AccountID = accountID
	req.View = "account"
	req.StartDate = startDate
	req.EndDate = endDate
	req.Format = format
	req.FraudType = fraudType
	req.Filters = filters

	reqBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Send Request Two: ", err)
		return nil, err
	}
	fmt.Println(string(reqBody))
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Send Request Three: ", err)
		return nil, err
	}

	// fmt.Printf("%#v\n", request)
	// fmt.Println(string(reqBody))
	request.Header.Add("Authentication-Key", authKey)
	machine := &http.Client{}

	res, err := machine.Do(request)
	if err != nil {
		fmt.Println("Send Request Four: ", err)
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Send Request Five: ", err)
		return nil, err
	}

	// fmt.Println(string(resBody))

	var resp interface{}

	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		fmt.Println("Send Request Six: ", err)
		return nil, err
	}

	return resp, nil

}
