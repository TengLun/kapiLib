package campaignclient

import "fmt"

// aPIAFake is the functional aPIAFakeccessor struct to communicate with the Kochava
// Premium Publisher API
type aPIAFake struct {
	appID   string
	authKey string
}

// GetCampaigns API provides the ability to retrieve the entire list of campaigns
// from the numerical App ID provided in the URL.
func (a aPIAFake) GetCampaigns(stats string) ([]Campaign, error) {
	return []Campaign{}, nil
}

// CreateCampaign API is used to create a new campaign by providing a JSON
// definition of the campaign.
func (a aPIAFake) CreateCampaign(name, destination string) (Campaign, error) {
	return Campaign{}, nil
}

// UpdateCampaign API is used to update an existing campaign by providing a JSON
// definition of the campaign with the modifications. If the campaign is
// successfully updated an HTTP 200 code and response, as shown below, is
// returned.``
func (a aPIAFake) UpdateCampaign(id, name string) (Campaign, error) {
	return Campaign{}, nil
}

// GetCampaign API provides the ability to retrieve a single campaign for the
// numerical Campaign ID provided in the URL.
func (a aPIAFake) GetCampaign(campaignID string) (Campaign, error) {
	return Campaign{}, nil
}

// GetSegments API provides the ability to retrieve the segments for the numerical
// Campaign ID provided in the URL.
func (a aPIAFake) GetSegments(campaignID string) ([]Segment, error) {
	return []Segment{}, nil
}

// CreateSegment API is used to create a new segment by providing a JSON definition of the segment.
func (a aPIAFake) CreateSegment(name, campaignID string) (Segment, error) {
	return Segment{}, nil
}

// UpdateSegment API is used to update an existing segment by providing a JSON definition
// of the segment with the modifications. If the segment is successfully updated
// an HTTP 200 code and response, as shown below, is returned.
func (a aPIAFake) UpdateSegment(name, campaignID, segmentID string) (Segment, error) {
	return Segment{}, nil
}

// GetSegment API provides the ability to retrieve a single segment for the numerical
// Segment ID provided in the URL.
func (a aPIAFake) GetSegment(campaignID, segmentID string) (Segment, error) {
	return Segment{}, nil
}

// GetTrackers API provides the ability to retrieve the entire list of trackers for the
// numerical App ID provided in the URL.
func (a aPIAFake) GetTrackers(query string) ([]Tracker, error) {
	return []Tracker{Tracker{}, Tracker{}}, nil
}

// UpdateTracker API is used to update an existing tracker by providing a JSON definition
// of the tracker with modifications. If the tracker is successfully updated an
// HTTP 200 code and response, as shown below, is returned.
func (a aPIAFake) UpdateTracker(updates Tracker) (Tracker, error) {
	fmt.Printf("%#v", updates)
	return Tracker{}, nil
}

// DeleteTracker API is used to delete an existing tracker by providing the numerical
// Tracker ID. If the tracker is deleted an HTTP 200 response will be returned,
// otherwise another HTTP code and message detailing the error will be returned.
func (a aPIAFake) DeleteTracker(trackerID string) error {
	return nil
}

// CreateTracker API is used to create a new tracker by providing a JSON definition of
// the tracker.
func (a aPIAFake) CreateTracker(name, trackerType, networkID, destinationURL, deeplinkURL, campaignID, segmentID, priceType string, priceValue float32, allowPublisherView bool, events []string, clickURLCustomParams []interface{}) (Tracker, error) {
	return Tracker{}, nil
}

// GetTrackerOverrides API provides the ability to retrieve the tracker overrides for the
// numerical Override ID provided in the URL.
func (a aPIAFake) GetTrackerOverrides(trackerID string) (GetOverridesResponse, error) {
	return GetOverridesResponse{}, nil
}

// PostTrackerOverrides API provides the ability to create tracker overrides for the numerical
// Tracker ID provided in the URL.
func (a aPIAFake) PostTrackerOverrides(trackerID string, overrides PostOverridesRequest) error {
	return nil
}
