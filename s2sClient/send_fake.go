package s2sclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// SendInstall to Kochava
func (c *ClientFake) SendInstall(installData S2SRequest, settings ...func(s *S2SRequest) error) error {

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

	return nil

}

// SendEvent to Kochava
func (c *ClientFake) SendEvent(eventData S2SRequest, settings ...func(s *S2SRequest) error) error {

	for _, setting := range settings {
		setting(&eventData)
	}

	eventData.Action = "event"

	body, err := json.Marshal(eventData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	return nil

}

// SendIdentity to Kochava
func (c *ClientFake) SendIdentity(identityRequest IdentityRequest) error {

	identityEndpoint := "http://control.kochava.com/v1/cpi/identityLink.php"

	body, err := json.Marshal(identityRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", identityEndpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	return nil

}
