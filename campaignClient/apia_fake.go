package campaignclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// aPIA_Fake is the functional aPIA_Fakeccessor struct to communicate with the Kochava
// Premium Publisher API
type aPIA_Fake struct {
	appID   string
	authKey string
}

// CreateApp creates an app
// TODO: Finish
func (a aPIA_Fake) createApp() error {
	return nil
}

// CreateCampaign creates one or more campaigns
// TODO: Finish
func (a aPIA_Fake) createCampaign(r ...campaignRequest) (campaignResponse, error) {

	cr := campaignRequest{
		Name:           "this campaign",
		DateEnd:        time.Unix(1490727343, 0),
		DateStart:      time.Unix(1490727343, 0),
		DestinationURL: "http://",
		Source:         "api",
	}

	body, err := json.Marshal(cr)
	if err != nil {
		fmt.Println(err)
	}

	endpoint := `https://campaign.api.kochava.com/campaign/` + a.appID
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authentication-Key", a.authKey)

	return campaignResponse{}, nil
}

// CreateSegment creates a segment for the tracker
func (a aPIA_Fake) createSegment(campaignID string, r ...segmentRequest) (segmentResponse, error) {

	// endpoint := `https://campaign.api.kochava.com/tier/` + campaignID

	return segmentResponse{}, nil
}

// CreateTracker creates a tracker in the UI; this is the primary item to create
// Creating a campaign and segment are necessary in order to create a tracker
func (a aPIA_Fake) createTracker(r ...trackerRequest) (trackerResponse, error) {

	// endpoint := `https://campaign.api.kochava.com/tracker/` + a.appID + `/create`

	return trackerResponse{}, nil
}

// RetrieveTrackers retrieves a list of campaigns
// TODO: Finish
func (a aPIA_Fake) retrieveTrackers() ([]trackerResponse, error) {

	// endpoint := `https://campaign.api.kochava.com/tracker/` + a.appID

	return []trackerResponse{}, nil
}

// VerifyCreated verifies the campaigns created match the templates retrieved
func (a aPIA_Fake) verifyCreated() error {
	return nil
}
