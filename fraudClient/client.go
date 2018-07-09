package fraudclient

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Client is the base struct that allows for access to all of the API methods
type Client struct {
	List   List
	Data   Data
	Add    Add
	Update Update
	Remove Remove

	debug bool
}

// List is a struct contained in Client that allows access to all of the API List
// methods
type List struct {
	AccountID string
	Format    string
	View      string
	APIKey    string

	BaseURL    string
	httpClient *http.Client

	debug bool
}

// Data is a struct contained in the client that allows access to all of the Data
// methods
type Data struct {
	AccountID string
	Format    string
	View      string
	APIKey    string

	BaseURL    string
	httpClient *http.Client

	debug bool
}

// Update is a struct allowing for updating entries in the Account Blacklist
type Update struct {
	AccountID string
	Format    string
	View      string
	APIKey    string

	BaseURL    string
	httpClient *http.Client

	debug bool
}

// Remove is a struct allowing for the removal of entries in the account blacklist
type Remove struct {
	AccountID string
	Format    string
	View      string
	APIKey    string

	BaseURL    string
	httpClient *http.Client

	debug bool
}

// Add is a struct allowing for the addition of entries to the account blacklist
type Add struct {
	AccountID string
	Format    string
	View      string
	APIKey    string

	BaseURL    string
	httpClient *http.Client

	debug bool
}

// CreateClient creates a client with the necessary information to access the
// methods
func CreateClient(apiKey string, accountID string, options ...func(*Client)) (*Client, error) {

	var client Client

	// Different accounts types have access to different views. The getView function
	// makes a quick call to establish which view the current client has access to
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
	client.List.BaseURL = `https://fraud.api.kochava.com:8320/fraud/`

	client.Data.AccountID = accountID
	client.Data.APIKey = apiKey
	client.Data.View = view
	client.Data.BaseURL = `https://fraud.api.kochava.com:8320/fraud/`

	for _, option := range options {
		option(&client)
	}

	if client.List.httpClient == nil {
		client.List.httpClient = &http.Client{}
		client.Data.httpClient = &http.Client{}
	}

	if client.List.BaseURL == "" {
		client.List.BaseURL = "https://fraud.api.kochava.com:8320/fraud/"
		client.Data.BaseURL = "https://fraud.api.kochava.com:8320/fraud/"
	}

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

	endpoint := l.BaseURL + fraudEndpointMap[fraudType] + `/list/apps`

	return sendRequest(l.AccountID, l.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, l.APIKey, filters)

}

// Networks lists networks with fraudulent data
func (l List) Networks(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := l.BaseURL + fraudEndpointMap[fraudType] + `/list/networks`

	return sendRequest(l.AccountID, l.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, l.APIKey, filters)

}

// Accounts lists Accounts with fraudulent data
func (l List) Accounts(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := l.BaseURL + fraudEndpointMap[fraudType] + `/list/accounts`

	return sendRequest(l.AccountID, l.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, l.APIKey, filters)

}

// Accounts returns data from accounts with fraudulent data
func (d Data) Accounts(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := d.BaseURL + fraudEndpointMap[fraudType] + `/data`

	return sendRequest(d.AccountID, d.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, d.APIKey, filters)

}

// Apps returns data from apps with fraudulent data
func (d Data) Apps(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := d.BaseURL + fraudEndpointMap[fraudType] + `/app/data`

	return sendRequest(d.AccountID, d.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, d.APIKey, filters)

}

// SiteIds returns data from siteIds with fraudulent data
func (d Data) SiteIds(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := d.BaseURL + fraudEndpointMap[fraudType] + `/siteid/data`

	return sendRequest(d.AccountID, d.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, d.APIKey, filters)

}

// Trackers returns data from trackers with fraudulent data
func (d Data) Trackers(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := d.BaseURL + fraudEndpointMap[fraudType] + `/tracker/data`

	return sendRequest(d.AccountID, d.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, d.APIKey, filters)

}

// Networks returns data from networks with fraudulent data
func (d Data) Networks(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (interface{}, error) {

	endpoint := d.BaseURL + fraudEndpointMap[fraudType] + `/network/data`

	return sendRequest(d.AccountID, d.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, d.APIKey, filters)

}
