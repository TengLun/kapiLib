package fraudclient

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Client is the base struct that allows for access to all of the API methods
type Client struct {
	List List
	Data Data

	BaseURL    *url.URL
	httpClient *http.Client
}

// List is a struct contained in Client that allows access to all of the API List
// methods
type List struct {
	AccountID string
	Format    string
	View      string
	APIKey    string
}

// Data is a struct contained in the client that allows access to all of the Data
// methods
type Data struct {
	AccountID string
	Format    string
	View      string
	APIKey    string
}

// CreateClient creates a client with the necessary information to access the
// methods
func CreateClient(apiKey string, accountID string) (*Client, error) {

	var client Client

	view, err := getView(apiKey)

	if err != nil {
		fmt.Println(err)
		return &Client{}, err
	}

	if err != nil {
		fmt.Println(err)
		return &Client{}, err
	}

	client.List.AccountID = accountID
	client.List.APIKey = apiKey
	client.List.View = view

	client.Data.AccountID = accountID
	client.Data.APIKey = apiKey
	client.Data.View = view

	return &client, nil
}

func getView(apiKey string) (string, error) {
	endpoint := "https://fraud.api.kochava.com:8320/fraud/installreceipt/tracker/data"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(`
{
  "view": "network",
  "fraudType": "installReceiptVerification",
  "accountId": "XXX",
  "startDate": "2016-11-13",
  "endDate": "2017-1-11",
  "format": "JSON",
  "filters": []
}`)))
	if err != nil {
		fmt.Println(err)
	}

	machine := &http.Client{}

	res, err := machine.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	switch res.StatusCode {
	case 200:
		return "network", nil
	case 403:
		return "account", nil
	default:
		return "", errors.New(res.Status)
	}

}

// Apps lists apps with fraudulent data
func (l List) Apps(fraudType string, startDate, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/list/apps`

	return sendRequest(l.AccountID, l.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, l.APIKey, filters)

}

// Networks lists networks with fraudulent data
func (l List) Networks(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/list/networks`

	return sendRequest(l.AccountID, l.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, l.APIKey, filters)

}

// Accounts lists Accounts with fraudulent data
func (l List) Accounts(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/list/accounts`

	return sendRequest(l.AccountID, l.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, l.APIKey, filters)

}

// Accounts returns data from accounts with fraudulent data
func (g Data) Accounts(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/data`

	return sendRequest(g.AccountID, g.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, g.APIKey, filters)

}

// Apps returns data from apps with fraudulent data
func (g Data) Apps(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/app/data`

	return sendRequest(g.AccountID, g.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, g.APIKey, filters)

}

// SiteIds returns data from siteIds with fraudulent data
func (g Data) SiteIds(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/siteid/data`

	return sendRequest(g.AccountID, g.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, g.APIKey, filters)

}

// Trackers returns data from trackers with fraudulent data
func (g Data) Trackers(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/tracker/data`

	return sendRequest(g.AccountID, g.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, g.APIKey, filters)

}

// Networks returns data from networks with fraudulent data
func (g Data) Networks(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/network/data`

	return sendRequest(g.AccountID, g.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, g.APIKey, filters)

}

// KFResponse is a generic response struct
type KFResponse struct {
	MetaData struct {
		Headers []string `json:"headers"`
	} `json:"metaData"`
	Data []struct {
		AppName         string `json:"appName,omitempty"`
		AppID           string `json:"appId,omitempty"`
		NetworkName     string `json:"networkName,omitempty"`
		NetworkID       string `json:"networkId,omitempty"`
		ClickCt         int    `json:"clickCt,omitempty"`
		SameAcctClickCt int    `json:"sameAcctClickCt,omitempty"`
		DiffAcctClickCt int    `json:"diffAcctClickCt,omitempty"`
		InstallCt       int    `json:"installCt,omitempty"`
	} `json:"data"`
}
