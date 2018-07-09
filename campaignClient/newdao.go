package campaignclient

import "net/http"

// AccountAccessor defines the account information
type AccountAccessor struct {
	AppID   string
	AuthKey string
}

// CreateClient returns an accessor object. If debug flag is true, a aPIA_Fake is
// returned for debugging purposes. Otherwise, an aPIA struct is returned. Options
// are passed using functional options.
func CreateClient(a AccountAccessor, options ...func(a *APIA) error) (APIAccessor, error) {
	var dao APIA

	for _, option := range options {
		option(&dao)
	}

	if dao.debug == true {
		var apiaFake aPIA_Fake
		return apiaFake, nil
	}

	var apia APIA
	apia.appID = a.AppID
	apia.authKey = a.AuthKey
	apia.client = &http.Client{}
	return apia, nil
}
