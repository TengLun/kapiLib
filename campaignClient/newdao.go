package campaignclient

import "net/http"

// AccountAccessor defines the account information
type AccountAccessor struct {
	AppID   string
	AuthKey string
}

// CreateClient returns an accessor object. If debug flag is true, a Client_Fake is
// returned for debugging purposes. Otherwise, an Client struct is returned. Options
// are passed using functional options.
func CreateClient(a AccountAccessor, options ...func(c *Client) error) (APIAccessor, error) {
	var dao Client

	for _, option := range options {
		option(&dao)
	}

	if dao.debug == true {
		var clientFake ClientFake
		return clientFake, nil
	}

	var client Client
	client.appID = a.AppID
	client.authKey = a.AuthKey
	client.client = &http.Client{}
	return client, nil
}
