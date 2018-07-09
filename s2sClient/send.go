package s2sclient

const (
	endpoint = `http://control.kochava.com/track/json`
)

// SendInstall to Kochava
func (c *Client) SendInstall(name string, installData struct{}) error {

	return nil
}

// SendEvent to Kochava
func (c *Client) SendEvent(name string, eventData struct{}) error {

	return nil
}

// SendIdentity to Kochava
func (c *Client) SendIdentity(identityName, identityValue string) error {

	return nil
}

// Request is a generalized request object to send an install or event to Kochava
type Request struct {
	Action       string `json:"action"`
	KochavaAppID string `json:"kochava_app_id"`
	AppVer       string `json:"app_ver"`
	Data         struct {
		OriginationIP string `json:"origination_ip,omit"`
		DeviceUa      string `json:"device_ua,omit"`
		DeviceVer     string `json:"device_ver,omit"`
		DeviceIds     struct {
			Idfa      string `json:"idfa,omit"`
			Idfv      string `json:"idfv,omit"`
			Imei      string `json:"imei,omit"`
			Adid      string `json:"adid,omit"`
			AndroidID string `json:"android_id,omit"`
		} `json:"device_ids"`
		EventName string `json:"event_name,omit"`
		Currency  string `json:"currency,omit"`
		EventData struct {
			ID   string `json:"id,omit"`
			Name string `json:"name,omit"`
			Sum  int    `json:"sum,omit"`
		} `json:"event_data,omit"`
	} `json:"data"`
}

func createInstallRequest() {

}

func createEventRequest() {

}

func createIdentityRequest() {

}
