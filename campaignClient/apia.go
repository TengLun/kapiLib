package campaignclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// aPIA is the functional aPIAccessor struct to communicate with the Kochava
// Premium Publisher API
type aPIA struct {
	appID   string
	authKey string
	client  *http.Client
	logger  *log.Logger
}

type Campaign struct {
	ID                     string `json:"id"`
	AppID                  string `json:"app_id"`
	Type                   string `json:"type"`
	Source                 string `json:"source"`
	Name                   string `json:"name"`
	DestinationURL         string `json:"destination_url"`
	DateCreated            int    `json:"date_created"`
	DateStart              int    `json:"date_start"`
	DateEnd                int    `json:"date_end"`
	DateTrackOutsideRange  bool   `json:"date_track_outside_range"`
	BudgetDaily            int    `json:"budget_daily"`
	BudgetWeekly           int    `json:"budget_weekly"`
	BudgetMax              int    `json:"budget_max"`
	TargetClicks           int    `json:"target_clicks"`
	TargetInstalls         int    `json:"target_installs"`
	Meta                   string `json:"meta"`
	LegacyIoGUID           string `json:"legacy_io_guid"`
	SmartLinkID            string `json:"smart_link_id"`
	WhatIfParentCampaignID string `json:"what_if_parent_campaign_id"`
}

// GetCampaigns API provides the ability to retrieve the entire list of campaigns
// from the numerical App ID provided in the URL.
func (a aPIA) GetCampaigns(stats string) ([]Campaign, error) {
	endpoint := fmt.Sprintf(`https://campaign.api.kochava.com/campaign/%v?stats=%v`, a.appID, stats)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return []Campaign{}, err
	}

	req.Header.Add("Authentication-Key", a.authKey)

	res, err := a.client.Do(req)
	if err != nil {
		return []Campaign{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []Campaign{}, err
	}

	switch {
	case res.StatusCode < 300 && res.StatusCode > 199:

		campaignList := make([]Campaign, 0)

		err = json.Unmarshal(body, &campaignList)
		if err != nil {
			fmt.Println(err)
			return []Campaign{}, err
		}
		fmt.Println(campaignList)

		return campaignList, nil

	default:
		fmt.Println(string(body))
		return []Campaign{}, err
	}

}

type CreateCampaignRequest struct {
	Name           string `json:"name"`
	DateEnd        string `json:"date_end"`
	DateStart      string `json:"date_start"`
	DestinationURL string `json:"destination_url"`
	Source         string `json:"source"`
}

// CreateCampaign API is used to create a new campaign by providing a JSON
// definition of the campaign.
func (a aPIA) CreateCampaign(name, destination string) (Campaign, error) {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/campaign/%v", a.appID)

	reqBody := CreateCampaignRequest{
		Name:           name,
		DestinationURL: destination,
		Source:         "api",
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return Campaign{}, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return Campaign{}, err
	}

	req.Header.Add("Authentication-Key", a.authKey)

	res, err := a.client.Do(req)
	if err != nil {
		return Campaign{}, err
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return Campaign{}, fmt.Errorf("api returned non-200 response:\nresponse_code: %v\nresponse_body: %v", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Campaign{}, err
	}

	var resBody Campaign

	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return Campaign{}, err
	}

	return resBody, nil
}

type UpdateCampaignRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UpdateCampaign API is used to update an existing campaign by providing a JSON
// definition of the campaign with the modifications. If the campaign is
// successfully updated an HTTP 200 code and response, as shown below, is
// returned.
func (a aPIA) UpdateCampaign(id, name string) (Campaign, error) {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/campaign/%v", a.appID)

	reqBody := UpdateCampaignRequest{
		ID:   id,
		Name: name,
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return Campaign{}, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return Campaign{}, err
	}

	req.Header.Add("Authentication-Key", a.authKey)

	res, err := a.client.Do(req)
	if err != nil {
		return Campaign{}, err
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return Campaign{}, fmt.Errorf("api returned non-200 response:\nresponse_code: %v\nresponse_body: %v", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Campaign{}, err
	}

	var resBody Campaign

	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return Campaign{}, err
	}

	return resBody, nil
}

// GetCampaigns API provides the ability to retrieve a single campaign for the
// numerical Campaign ID provided in the URL.
func (a aPIA) GetCampaign(campaignID string) (Campaign, error) {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/campaign/%v/%v", a.appID, campaignID)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return Campaign{}, err
	}

	req.Header.Add("Authentication-Key", a.authKey)

	res, err := a.client.Do(req)
	if err != nil {
		return Campaign{}, err
	}

	return campaign, nil
}

// This API provides the ability to retrieve the segments for the numerical
// Campaign ID provided in the URL.
func (a aPIA) GetSegments() {

}

// This API is used to create a new segment by providing a JSON definition of the segment.
func (a aPIA) CreateSegment() {

}

// This API is used to update an existing segment by providing a JSON definition
// of the segment with the modifications. If the segment is successfully updated
// an HTTP 200 code and response, as shown below, is returned.
func (a aPIA) UpdateSegment() {

}

// This API provides the ability to retrieve a single segment for the numerical
// Segment ID provided in the URL.
func (a aPIA) GetSegment() {

}

// This API provides the ability to retrieve the entire list of trackers for the
// numerical App ID provided in the URL.
func (a aPIA) GetTrackers() {

}

// This API is used to update an existing tracker by providing a JSON definition
// of the tracker with modifications. If the tracker is successfully updated an
// HTTP 200 code and response, as shown below, is returned.
func (a aPIA) UpdateTracker() {

}

// This API is used to delete an existing tracker by providing the numerical
// Tracker ID. If the tracker is deleted an HTTP 200 response will be returned,
// otherwise another HTTP code and message detailing the error will be returned.
func (a aPIA) DeleteTracker() {

}

// This API is used to create a new tracker by providing a JSON definition of
// the tracker.
func (a aPIA) CreateTracker() {

}

// This API provides the ability to retrieve the tracker overrides for the
// numerical Override ID provided in the URL.
func (a aPIA) GetTrackerOverrides() {

}

// This API provides the ability to create tracker overrides for the numerical
// Tracker ID provided in the URL.
func (a aPIA) PostTrackerOverrides() {

}
