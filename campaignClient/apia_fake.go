package campaignclient

// aPIA_Fake is the functional aPIA_Fakeccessor struct to communicate with the Kochava
// Premium Publisher API
type aPIA_Fake struct {
	appID   string
	authKey string
}

// GetCampaigns API provides the ability to retrieve the entire list of campaigns
// from the numerical App ID provided in the URL.
func (a aPIA_Fake) GetCampaigns(stats string) ([]GetCampaignsResponse, error) {
	return []GetCampaignsResponse{}, nil
}

// CreateCampaign API is used to create a new campaign by providing a JSON
// definition of the campaign.
func (a aPIA_Fake) CreateCampaign() {

}

// UpdateCampaign API is used to update an existing campaign by providing a JSON
// definition of the campaign with the modifications. If the campaign is
// successfully updated an HTTP 200 code and response, as shown below, is
// returned.
func (a aPIA_Fake) UpdateCampaign() {

}

// GetCampaigns API provides the ability to retrieve a single campaign for the
// numerical Campaign ID provided in the URL.
func (a aPIA_Fake) GetCampaign() {

}

// This API provides the ability to retrieve the segments for the numerical
// Campaign ID provided in the URL.
func (a aPIA_Fake) GetSegments() {

}

// This API is used to create a new segment by providing a JSON definition of the segment.
func (a aPIA_Fake) CreateSegment() {

}

// This API is used to update an existing segment by providing a JSON definition
// of the segment with the modifications. If the segment is successfully updated
// an HTTP 200 code and response, as shown below, is returned.
func (a aPIA_Fake) UpdateSegment() {

}

// This API provides the ability to retrieve a single segment for the numerical
// Segment ID provided in the URL.
func (a aPIA_Fake) GetSegment() {

}

// This API provides the ability to retrieve the entire list of trackers for the
// numerical App ID provided in the URL.
func (a aPIA_Fake) GetTrackers() {

}

// This API is used to update an existing tracker by providing a JSON definition
// of the tracker with modifications. If the tracker is successfully updated an
// HTTP 200 code and response, as shown below, is returned.
func (a aPIA_Fake) UpdateTracker() {

}

// This API is used to delete an existing tracker by providing the numerical
// Tracker ID. If the tracker is deleted an HTTP 200 response will be returned,
// otherwise another HTTP code and message detailing the error will be returned.
func (a aPIA_Fake) DeleteTracker() {

}

// This API is used to create a new tracker by providing a JSON definition of
// the tracker.
func (a aPIA_Fake) CreateTracker() {

}

// This API provides the ability to retrieve the tracker overrides for the
// numerical Override ID provided in the URL.
func (a aPIA_Fake) GetTrackerOverrides() {

}

// This API provides the ability to create tracker overrides for the numerical
// Tracker ID provided in the URL.
func (a aPIA_Fake) PostTrackerOverrides() {

}
