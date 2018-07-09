package campaignclient

import "time"

// campaignRequest is CreateCampaign's request object
type campaignRequest struct {
	Name           string    `json:"name"`
	DateEnd        time.Time `json:"date_end"`
	DateStart      time.Time `json:"date_start"`
	DestinationURL string    `json:"destination_url"`
	Source         string    `json:"source"`
}

// campaignResponse is CreateCampaign's response object
type campaignResponse struct {
	ID                     string      `json:"id"`
	AppID                  string      `json:"app_id"`
	Type                   string      `json:"type"`
	Source                 string      `json:"source"`
	Name                   string      `json:"name"`
	DestinationURL         string      `json:"destination_url"`
	DateCreated            int         `json:"date_created"`
	DateStart              int         `json:"date_start"`
	DateEnd                int         `json:"date_end"`
	DateTrackOutsideRange  bool        `json:"date_track_outside_range"`
	BudgetDaily            int         `json:"budget_daily"`
	BudgetWeekly           int         `json:"budget_weekly"`
	BudgetMax              int         `json:"budget_max"`
	TargetClicks           int         `json:"target_clicks"`
	TargetInstalls         int         `json:"target_installs"`
	Meta                   string      `json:"meta"`
	LegacyIoGUID           string      `json:"legacy_io_guid"`
	SmartLinkID            interface{} `json:"smart_link_id"`
	WhatIfParentCampaignID interface{} `json:"what_if_parent_campaign_id"`
}

// segmentRequest is CreateSegment's request Object
type segmentRequest struct {
	Name       string `json:"name"`
	Source     string `json:"source"`
	CampaignID string
}

// segmentResponse is CreateSegment's response Object
type segmentResponse struct {
	ID                 string      `json:"id"`
	AppID              string      `json:"app_id"`
	CampaignID         string      `json:"campaign_id"`
	Source             string      `json:"source"`
	Name               string      `json:"name"`
	TargetGeo          string      `json:"target_geo"`
	DateCreated        int         `json:"date_created"`
	WhatIfParentTierID interface{} `json:"what_if_parent_tier_id"`
}

// trackerRequest is CreateTracker's request object
// TODO: Currently removing reengagement functionality; it needs to be implemented
// for version 1.0 (production version)
// NOTE: all potential request parameters are included below, however
//	unnecessary parameters have been commented out
type trackerRequest struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	NetworkID      string `json:"network_id"`
	DestinationURL string `json:"destination_url"`
	// DestinationURLReengagement string        `json:"destination_url_reengagement"`
	// NetworkPricing             string        `json:"network_pricing"`
	// NetworkPrice               int           `json:"network_price"`
	// PermPublisherAllowView     bool          `json:"perm_publisher_allow_view"`
	// ClickURLCustomParams       []interface{} `json:"click_url_custom_params"`
	TierID     string `json:"tier_id"`
	CampaignID string `json:"campaign_id"`
	// DestinationData            struct {
	// 	Type    string `json:"type,omitempty"`
	// 	TypeObj string `json:"typeObj,omitempty"`
	// } `json:"destination_data,omitempty"`
	// Events []string `json:"events"`
}

// trackerResponse is CreateTracker's response object
type trackerResponse struct {
	ID                         string        `json:"id"`
	TierID                     string        `json:"tier_id"`
	CampaignID                 string        `json:"campaign_id"`
	AppID                      string        `json:"app_id"`
	DateCreated                time.Time     `json:"date_created"`
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
	NetworkPrice               int           `json:"network_price"`
	BudgetDaily                int           `json:"budget_daily"`
	BudgetWeekly               int           `json:"budget_weekly"`
	BudgetMax                  int           `json:"budget_max"`
	RtbID                      interface{}   `json:"rtb_id"`
	RtbDefinition              interface{}   `json:"rtb_definition"`
	Meta                       string        `json:"meta"`
	LegacyCampaignID           string        `json:"legacy_campaign_id"`
	LegacyPostID               string        `json:"legacy_post_id"`
	PermPublisherAllowView     bool          `json:"perm_publisher_allow_view"`
	IsActive                   bool          `json:"is_active"`
	CreativeIds                interface{}   `json:"creative_ids"`
	ClickURLCustomParams       []interface{} `json:"click_url_custom_params"`
	DestinationData            struct {
		Type    string `json:"type"`
		TypeObj string `json:"typeObj"`
	} `json:"destination_data"`
	RtbUpdateStatus            string      `json:"rtb_update_status"`
	RtbUpdateResponse          interface{} `json:"rtb_update_response"`
	RtbUpdatePid               interface{} `json:"rtb_update_pid"`
	S2SDestination             interface{} `json:"s2s_destination"`
	PostbackURL                interface{} `json:"postback_url"`
	VerificationRules          interface{} `json:"verification_rules"`
	SmartLinkID                interface{} `json:"smart_link_id"`
	WhatIfParentTrackerID      interface{} `json:"what_if_parent_tracker_id"`
	NetworkName                string      `json:"network_name"`
	NetworkIsSelfAttributing   bool        `json:"network_is_self_attributing"`
	NetworkSupportsImpressions bool        `json:"network_supports_impressions"`
	CampaignName               string      `json:"campaign_name"`
	TierName                   string      `json:"tier_name"`
	Events                     interface{} `json:"events"`
	AppGUID                    string      `json:"app_guid"`
	AgencyTrackerID            interface{} `json:"agency_tracker_id"`
	TwttterEventGUID           string      `json:"twttter_event_guid"`
	GoogleAndroidPostbackURL   string      `json:"google_android_postback_url"`
	GoogleIosPostbackURL       string      `json:"google_ios_postback_url"`
}
