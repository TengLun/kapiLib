// Package campaignclient is a library for the Kochava Campaign Management API
// The Campaign Management API Integration provides the programmatic tools
// to create and maintain campaigns, segments and trackers in the Kochava system.
package campaignclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client is the functional APIAccessor struct to communicate with the Kochava
// Premium Publisher API
type Client struct {
	appID   string
	authKey string

	debug bool

	client *http.Client
}

// Campaign struct defines the information for a campaign object
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
func (c Client) GetCampaigns(stats string) ([]Campaign, error) {
	endpoint := fmt.Sprintf(`https://campaign.api.kochava.com/campaign/%v?stats=%v`, c.appID, stats)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return []Campaign{}, err
	}

	req.Header.Add("Authentication-Key", c.authKey)

	res, err := c.client.Do(req)
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

// CreateCampaignRequest contains necessary information to create a campaign
type CreateCampaignRequest struct {
	Name           string `json:"name"`
	DateEnd        string `json:"date_end"`
	DateStart      string `json:"date_start"`
	DestinationURL string `json:"destination_url"`
	Source         string `json:"source"`
}

// CreateCampaign API is used to create a new campaign by providing a JSON
// definition of the campaign.
func (c Client) CreateCampaign(name, destination string) (Campaign, error) {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/campaign/%v", c.appID)

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

	req.Header.Add("Authentication-Key", c.authKey)

	res, err := c.client.Do(req)
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

// UpdateCampaignRequest contains the necessary information to update a campaign
type UpdateCampaignRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UpdateCampaign API is used to update an existing campaign by providing a JSON
// definition of the campaign with the modifications. If the campaign is
// successfully updated an HTTP 200 code and response, as shown below, is
// returned.
func (c Client) UpdateCampaign(id, name string) (Campaign, error) {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/campaign/%v", c.appID)

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

	req.Header.Add("Authentication-Key", c.authKey)

	res, err := c.client.Do(req)
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

// GetCampaign API provides the ability to retrieve a single campaign for the
// numerical Campaign ID provided in the URL.
func (c Client) GetCampaign(campaignID string) (Campaign, error) {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/campaign/%v/%v", c.appID, campaignID)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return Campaign{}, err
	}

	req.Header.Add("Authentication-Key", c.authKey)

	res, err := c.client.Do(req)
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

// Segment struct contains all of the information for a segment object
type Segment struct {
	ID                 string `json:"id"`
	AppID              string `json:"app_id"`
	CampaignID         string `json:"campaign_id"`
	Source             string `json:"source"`
	Name               string `json:"name"`
	TargetGeo          string `json:"target_geo"`
	DateCreated        int    `json:"date_created"`
	WhatIfParentTierID string `json:"what_if_parent_tier_id"`
}

// GetSegments API provides the ability to retrieve the segments for the numerical
// Campaign ID provided in the URL.
func (c Client) GetSegments(campaignID string) ([]Segment, error) {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/tier/%v", campaignID)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return []Segment{}, err
	}

	req.Header.Add("Authentication-Key", c.authKey)

	res, err := c.client.Do(req)
	if err != nil {
		return []Segment{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []Segment{}, err
	}

	switch {
	case res.StatusCode < 300 && res.StatusCode > 199:

		segmentList := make([]Segment, 0)

		err = json.Unmarshal(body, &segmentList)
		if err != nil {
			fmt.Println(err)
			return []Segment{}, err
		}

		return segmentList, nil

	default:
		fmt.Println(string(body))
		return []Segment{}, err
	}

}

// CreateSegmentRequest contains the necessary information to create a segment
type CreateSegmentRequest struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}

// CreateSegment API is used to create a new segment by providing a JSON definition of the segment.
func (c Client) CreateSegment(name, campaignID string) (Segment, error) {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/tier/%v", campaignID)

	reqBody := CreateSegmentRequest{
		Name:   name,
		Source: "api",
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return Segment{}, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return Segment{}, err
	}

	req.Header.Add("Authentication-Key", c.authKey)

	res, err := c.client.Do(req)
	if err != nil {
		return Segment{}, err
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return Segment{}, fmt.Errorf("api returned non-200 response:\nresponse_code: %v\nresponse_body: %v", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Segment{}, err
	}

	var resBody Segment

	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return Segment{}, err
	}

	return resBody, nil
}

// UpdateSegmentRequest contains the necessary information to update a segment
type UpdateSegmentRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UpdateSegment API is used to update an existing segment by providing a JSON definition
// of the segment with the modifications. If the segment is successfully updated
// an HTTP 200 code and response, as shown below, is returned.
func (c Client) UpdateSegment(name, campaignID, segmentID string) (Segment, error) {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/tier/%v", campaignID)

	reqBody := UpdateSegmentRequest{
		ID:   segmentID,
		Name: name,
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return Segment{}, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return Segment{}, err
	}

	req.Header.Add("Authentication-Key", c.authKey)

	res, err := c.client.Do(req)
	if err != nil {
		return Segment{}, err
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return Segment{}, fmt.Errorf("api returned non-200 response:\nresponse_code: %v\nresponse_body: %v", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Segment{}, err
	}

	var resBody Segment

	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return Segment{}, err
	}

	return resBody, nil
}

// GetSegment API provides the ability to retrieve a single segment for the numerical
// Segment ID provided in the URL.
func (c Client) GetSegment(campaignID, segmentID string) (Segment, error) {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/tier/%v/%v", campaignID, segmentID)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return Segment{}, err
	}

	req.Header.Add("Authentication-Key", c.authKey)

	res, err := c.client.Do(req)
	if err != nil {
		return Segment{}, err
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return Segment{}, fmt.Errorf("api returned non-200 response:\nresponse_code: %v\nresponse_body: %v", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Segment{}, err
	}

	var resBody Segment

	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return Segment{}, err
	}

	return resBody, nil
}

// Tracker struct contains all of the information for a tracker object
type Tracker struct {
	ID                         string        `json:"id"`
	TierID                     string        `json:"tier_id"`
	CampaignID                 string        `json:"campaign_id"`
	AppID                      string        `json:"app_id"`
	DateCreated                string        `json:"date_created"`
	Source                     string        `json:"source"`
	GUID                       string        `json:"guid"`
	Name                       string        `json:"name"`
	Type                       string        `json:"type"`
	ClickTrackingURL           string        `json:"click_tracking_url"`
	ImpTrackingURL             string        `json:"imp_tracking_url"`
	DestinationURL             string        `json:"destination_url"`
	DestinationURLReengagement string        `json:"destination_url_reengagement"`
	NetworkID                  string        `json:"network_id"`
	NetworkPricing             string        `json:"network_pricing"`
	NetworkPrice               float64       `json:"network_price"`
	BudgetDaily                float64       `json:"budget_daily"`
	BudgetWeekly               float64       `json:"budget_weekly"`
	BudgetMax                  float64       `json:"budget_max"`
	RtbID                      string        `json:"rtb_id"`
	RtbDefinitions             string        `json:"rtb_definitions"`
	Meta                       string        `json:"meta"`
	LegacyCampaignID           string        `json:"legacy_campaign_id"`
	LegacyPostID               string        `json:"legacy_post_id"`
	PermPublisherAllowView     bool          `json:"perm_publisher_allow_view"`
	IsActive                   bool          `json:"is_active"`
	CreativeIds                []interface{} `json:"creative_ids"`
	ClickURLCustomParams       []interface{} `json:"click_url_custom_params"`
	DestinationData            struct {
		Type    string `json:"type,omitempty"`
		TypeObj string `json:"typeObj,omitempty"`
	} `json:"destination_data"`
	RtbUpdateStatus          string      `json:"rtb_update_status"`
	RtbUpdateResponse        interface{} `json:"rtb_update_response"`
	RtbUpdatePid             interface{} `json:"rtb_update_pid"`
	S2SDestination           interface{} `json:"s2s_destination"`
	PostbackURL              interface{} `json:"postback_url"`
	VerificationRules        interface{} `json:"verification_rules"`
	SmartLinkID              interface{} `json:"smart_link_id"`
	WhatIfParentTrackerID    interface{} `json:"what_if_parent_tracker_id"`
	NetworkName              string      `json:"network_name"`
	NetworkIsSelfAttributing bool        `json:"network_is_self_attributing"`
	CampaignName             string      `json:"campaign_name"`
	TierName                 string      `json:"tier_name"`
	AppGUID                  string      `json:"app_guid"`
	AgencyTrackerID          interface{} `json:"agency_tracker_id"`
	GoogleAndroidPostbackURL string      `json:"google_android_postback_url"`
	GoogleIosPostbackURL     string      `json:"google_ios_postback_url"`
	TwitterEventGUID         string      `json:"twttter_event_guid,omitempty"`
	AgencyNetworkID          string      `json:"agency_network_id,omitempty"`
	Events                   []string    `json:"events,omitempty"`
}

// GetTrackers API provides the ability to retrieve the entire list of trackers for the
// numerical App ID provided in the URL. NOTE: This function does not currently
// have support for a querystring to filter trackers, even though that ability
// exists in the API. Support will be added for that eventually.
func (c Client) GetTrackers(query string) ([]Tracker, error) {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/tracker/%v?%v", c.appID, query)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return []Tracker{}, err
	}

	req.Header.Add("Authentication-Key", c.authKey)

	res, err := c.client.Do(req)
	if err != nil {
		return []Tracker{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []Tracker{}, err
	}

	switch {
	case res.StatusCode < 300 && res.StatusCode > 199:

		trackerList := make([]Tracker, 0)

		err = json.Unmarshal(body, &trackerList)
		if err != nil {
			fmt.Println(err)
			return []Tracker{}, err
		}

		return trackerList, nil

	default:
		fmt.Println(string(body))
		return []Tracker{Tracker{}, Tracker{}}, err
	}

}

// UpdateTrackerRequest contains all of the necessary information to update a tracker
type UpdateTrackerRequest struct {
	ID                         string        `json:"id"`
	Name                       string        `json:"name,omitempty"`
	DestinationURL             string        `json:"destination_url,omitempty"`
	DestinationURLReengagement string        `json:"destination_url_reengagement,omitempty"`
	LandingPageID              string        `json:"landing_page_id,omitempty"`
	NetworkPricing             string        `json:"network_pricing,omitempty"`
	NetworkPrice               float64       `json:"network_price,omitempty"`
	PermPublisherAllowView     bool          `json:"perm_publisher_allow_view,omitempty"`
	ClickURLCustomParams       []interface{} `json:"click_url_custom_params,omitempty"`
	S2SDestination             interface{}   `json:"s2s_destination,omitempty"`
	VerificationRules          interface{}   `json:"verification_rules,omitempty"`
	TierID                     string        `json:"tier_id,omitempty"`
	CampaignID                 string        `json:"campaign_id,omitempty"`
	DestinationData            struct {
		Type    string `json:"type,omitempty"`
		TypeObj string `json:"typeObj,omitempty"`
	} `json:"destination_data,omitempty"`
	AgencyNetworkID string   `json:"agency_network_id,omitempty"`
	Events          []string `json:"events,omitempty"`
}

// UpdateTracker API is used to update an existing tracker by providing a JSON definition
// of the tracker with modifications. If the tracker is successfully updated an
// HTTP 200 code and response, as shown below, is returned.
func (c Client) UpdateTracker(updates Tracker) (Tracker, error) {

	updateRequest := UpdateTrackerRequest{
		ID:                         updates.ID,
		Name:                       updates.Name,
		DestinationURL:             updates.DestinationURL,
		DestinationURLReengagement: updates.DestinationURLReengagement,
		LandingPageID:              "",
		NetworkPricing:             updates.NetworkPricing,
		NetworkPrice:               updates.NetworkPrice,
		PermPublisherAllowView:     updates.PermPublisherAllowView,
		ClickURLCustomParams:       updates.ClickURLCustomParams,
		S2SDestination:             updates.S2SDestination,
		VerificationRules:          updates.VerificationRules,
		TierID:                     updates.TierID,
		CampaignID:                 updates.CampaignID,
		DestinationData:            updates.DestinationData,
		AgencyNetworkID:            updates.AgencyNetworkID,
		Events:                     updates.Events,
	}

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/tracker/%v", c.appID)

	reqBodyBytes, err := json.Marshal(updateRequest)
	if err != nil {
		return Tracker{}, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return Tracker{}, err
	}

	req.Header.Add("Authentication-Key", c.authKey)

	res, err := c.client.Do(req)
	if err != nil {
		return Tracker{}, err
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return Tracker{}, fmt.Errorf("api returned non-200 response:\nresponse_code: %v\nresponse_body: %v", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Tracker{}, err
	}

	var resBody Tracker

	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return Tracker{}, err
	}

	return resBody, nil
}

// DeleteTracker API is used to delete an existing tracker by providing the numerical
// Tracker ID. If the tracker is deleted an HTTP 200 response will be returned,
// otherwise another HTTP code and message detailing the error will be returned.
func (c Client) DeleteTracker(trackerID string) error {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/tracker/%v/delete/%v", c.appID, trackerID)

	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authentication-Key", c.authKey)

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode > 199 && res.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("non-200 response: %v", res.Status)
}

// CreateTrackerRequest contains all of the information to create a tracker
type CreateTrackerRequest struct {
	Name                       string        `json:"name"`
	Type                       string        `json:"type"`
	NetworkID                  string        `json:"network_id"`
	DestinationURL             string        `json:"destination_url"`
	DestinationURLReengagement string        `json:"destination_url_reengagement"`
	NetworkPricing             string        `json:"network_pricing"`
	NetworkPrice               float32       `json:"network_price"`
	PermPublisherAllowView     bool          `json:"perm_publisher_allow_view"`
	ClickURLCustomParams       []interface{} `json:"click_url_custom_params"`
	TierID                     string        `json:"tier_id"`
	CampaignID                 string        `json:"campaign_id"`
	DestinationData            struct {
		Type    string `json:"type"`
		TypeObj string `json:"typeObj"`
	} `json:"destination_data"`
	Events []string `json:"events"`
}

// CreateTracker API is used to create a new tracker by providing a JSON definition of
// the tracker.
func (c Client) CreateTracker(name, trackerType, networkID, destinationURL, deeplinkURL, campaignID, segmentID, priceType string, priceValue float32, allowPublisherView bool, events []string, clickURLCustomParams []interface{}) (Tracker, error) {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/tracker/%v/create", c.appID)

	reqBody := CreateTrackerRequest{
		Name:                       name,
		Type:                       trackerType,
		NetworkID:                  networkID,
		DestinationURL:             destinationURL,
		DestinationURLReengagement: deeplinkURL,
		NetworkPricing:             priceType,
		NetworkPrice:               priceValue,
		PermPublisherAllowView:     allowPublisherView,
		ClickURLCustomParams:       clickURLCustomParams,
		TierID:                     segmentID,
		CampaignID:                 campaignID,
		Events:                     events,
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return Tracker{}, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return Tracker{}, err
	}

	req.Header.Add("Authentication-Key", c.authKey)

	res, err := c.client.Do(req)
	if err != nil {
		return Tracker{}, err
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return Tracker{}, fmt.Errorf("api returned non-200 response:\nresponse_code: %v\nresponse_body: %v", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Tracker{}, err
	}

	var resBody Tracker

	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return Tracker{}, err
	}

	return resBody, nil
}

// GetOverridesResponse Contains the informations for an overrides object
type GetOverridesResponse struct {
	Success          bool   `json:"success"`
	RequestID        string `json:"request_id"`
	TrackerOverrides []struct {
		DeviceIDLookbackWindow struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"device_id_lookback_window"`
		EnableVerification struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"enable_verification"`
		FingerprintLookbackWindow struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"fingerprint_lookback_window"`
		ImpressionLookbackDevice struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"impression_lookback_device"`
		ImpressionLookbackFingerprint struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"impression_lookback_fingerprint"`
		ImpressionLookbackIP struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"impression_lookback_ip"`
		ImpressionLookbackPartialIP struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"impression_lookback_partial_ip"`
		ImpressionsDisableAttribution struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"impressions_disable_attribution"`
		IPLookbackWindow struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"ip_lookback_window"`
	} `json:"tracker_overrides"`
}

// GetTrackerOverrides API provides the ability to retrieve the tracker overrides for the
// numerical Override ID provided in the URL.
func (c Client) GetTrackerOverrides(trackerID string) (GetOverridesResponse, error) {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/tracker/override?id=%v", trackerID)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Request Creation Failed: ")
		return GetOverridesResponse{}, err
	}

	req.Header.Add("Authentication-Key", c.authKey)

	res, err := c.client.Do(req)
	if err != nil {
		fmt.Println("Request Failed: ")
		return GetOverridesResponse{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Response Body ReadAll Failed: ")
		return GetOverridesResponse{}, err
	}

	var resBody GetOverridesResponse

	err = json.Unmarshal(body, &resBody)
	if err != nil {
		fmt.Println("JSON UnMarshal Failed: ")
		fmt.Printf("%#v\n", string(body))
		return GetOverridesResponse{}, err
	}

	return resBody, nil
}

// PostOverridesRequest contains the necessary information to configure overrides
// on a tracker
type PostOverridesRequest struct {
	TrackerOverrides struct {
		DeviceIDLookbackWindow struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"device_id_lookback_window"`
		EnableVerification struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"enable_verification"`
		FingerprintLookbackWindow struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"fingerprint_lookback_window"`
		ImpressionLookbackDevice struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"impression_lookback_device"`
		ImpressionLookbackFingerprint struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"impression_lookback_fingerprint"`
		ImpressionLookbackIP struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"impression_lookback_ip"`
		ImpressionLookbackPartialIP struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"impression_lookback_partial_ip"`
		ImpressionsDisableAttribution struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"impressions_disable_attribution"`
		IPLookbackWindow struct {
			Name  string `json:"name"`
			Key   string `json:"key"`
			Group string `json:"group"`
			Type  string `json:"type"`
			Value int    `json:"value"`
			Admin int    `json:"admin"`
		} `json:"ip_lookback_window"`
	} `json:"tracker_overrides"`
}

// PostTrackerOverrides API provides the ability to create tracker overrides for the numerical
// Tracker ID provided in the URL.
func (c Client) PostTrackerOverrides(trackerID string, overrides PostOverridesRequest) error {

	endpoint := fmt.Sprintf("https://campaign.api.kochava.com/tracker/override?id=%v", trackerID)

	reqBody, err := json.Marshal(overrides)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Add("Authentication-Key", c.authKey)

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode > 199 && res.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("non-200 response: %v", res.Status)

}
