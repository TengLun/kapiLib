package s2sclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	endpoint = `http://control.kochava.com/track/json`
)

// SendInstall to Kochava
func (c *Client) SendInstall(installData S2SRequest, settings ...func(s *S2SRequest) error) error {

	for _, setting := range settings {
		setting(&installData)
	}

	installData.Action = "install"

	body, err := json.Marshal(installData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode > 199 && res.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("non-200 response: statuscode: %v, status: %v", res.StatusCode, res.Status)
}

// SendEvent to Kochava
func (c *Client) SendEvent(eventData S2SRequest) error {

	if err != nil {
		return err
	}
	err := c.SendInstall(
		SetDeviceUA("Generic Android 6.0"),
		SetDeviceID(map[string]string{
			"idfa":   "12345-6789-0123",
			"custom": "my_custom_id",
		}),
		SetIPAddress("127.0.0.2"),
		SetDeviceCurrency("USD"),
	)

	eventData.Action = "event"

	body, err := json.Marshal(eventData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode > 199 && res.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("non-200 response: statuscode: %v, status: %v", res.StatusCode, res.Status)
}

// IdentityRequest sends user_id or account_id, etc, to the Kochava system for a specific device
type IdentityRequest struct {
	KochavaAppID string `json:"kochava_app_id"`
	DeviceID     struct {
		Idfa string `json:"idfa,omitempty"`
		Adid string `json:"adid,omitempty"`
	} `json:"device_id"`
	DeviceHashMethod struct {
		Mac string `json:"mac,omitempty"`
	} `json:"device_hash_method,omitempty"`
	Data map[string]interface {
	} `json:"data"`
}

// SendIdentity to Kochava
func (c *Client) SendIdentity(identityRequest IdentityRequest) error {

	identityEndpoint := "http://control.kochava.com/v1/cpi/identityLink.php"

	body, err := json.Marshal(identityRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", identityEndpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode > 199 && res.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("non-200 response: statuscode: %v, status: %v", res.StatusCode, res.Status)
}

// S2SRequest is a generalized request object to send an install or event to Kochava
type S2SRequest struct {
	Action       string `json:"action"`
	KochavaAppID string `json:"kochava_app_id"`
	AppVer       string `json:"app_ver"`
	Data         struct {
		OriginationIP         string            `json:"origination_ip,omitempty"`
		DeviceUa              string            `json:"device_ua,omitempty"`
		DeviceVer             string            `json:"device_ver,omitempty"`
		DeviceIds             map[string]string `json:"device_ids"`
		EventName             string            `json:"event_name,omitempty"`
		Currency              string            `json:"currency,omitempty"`
		EventData             map[string]string `json:"event_data,omitempty"`
		AppVer                string            `json:"app_ver"`
		IadAttributionDetails struct {
			Version31 struct {
				IadAttribution    string `json:"iad-attribution,omitempty"`
				IadLineitemID     string `json:"iad-lineitem-id,omitempty"`
				IadKeyword        string `json:"iad-keyword,omitempty"`
				IadOrgName        string `json:"iad-org-name,omitempty"`
				IadClickDate      string `json:"iad-click-date,omitempty"`
				IadAdgroupName    string `json:"iad-adgroup-name,omitempty"`
				IadCampaignID     string `json:"iad-campaign-id,omitempty"`
				IadAdgroupID      string `json:"iad-adgroup-id,omitempty"`
				IadLineitemName   string `json:"iad-lineitem-name,omitempty"`
				IadCampaignName   string `json:"iad-campaign-name,omitempty"`
				IadConversionDate string `json:"iad-conversion-date,omitempty"`
			} `json:"Version3.1,omitempty"`
		} `json:"iad_attribution_details,omitempty"`
	} `json:"data"`
}

// SetDeviceUA configures the requests device UA
func SetDeviceUA(deviceUA string) func(s *S2SRequest) error {
	return func(s *S2SRequest) error {
		s.Data.DeviceUa = deviceUA
		return nil
	}
}

// SetDeviceID configures the requests device ID
func SetDeviceID(deviceIds map[string]string) func(s *S2SRequest) error {
	return func(s *S2SRequest) error {
		s.Data.DeviceIds = deviceIds
		return nil
	}
}

// SetEventData configures the event data for the request
func SetEventData(eventData map[string]string) func(s *S2SRequest) error {
	return func(s *S2SRequest) error {
		s.Data.EventData = eventData
		return nil
	}
}

// SetIPAddress configures the requests IP Address
func SetIPAddress(ip string) func(s *S2SRequest) error {
	return func(s *S2SRequest) error {
		s.Data.OriginationIP = ip
		return nil
	}
}

// SetDeviceCurrency sets the devices local currency
func SetDeviceCurrency(currency string) func(s *S2SRequest) error {
	return func(s *S2SRequest) error {
		s.Data.Currency = currency
		return nil
	}

}
