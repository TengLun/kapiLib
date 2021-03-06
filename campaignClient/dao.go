package campaignclient

// APIAccessor interface allows for interfacing with the Kochava Premium Publisher
// API; it creates apps and campaigns in the Kochava system. aPIAccessor has
// two structs that satisfy it: aPIA is a working accessor; aPIA_Fake is a test
// accessor
type APIAccessor interface {

	// GetCampaigns API provides the ability to retrieve the entire list of campaigns
	// from the numerical App ID provided in the URL.
	GetCampaigns(stats string) ([]Campaign, error)

	// CreateCampaign API is used to create a new campaign by providing a JSON
	// definition of the campaign.
	CreateCampaign(name, destination string) (Campaign, error)

	// UpdateCampaign API is used to update an existing campaign by providing a JSON
	// definition of the campaign with the modifications. If the campaign is
	// successfully updated an HTTP 200 code and response, as shown below, is
	// returned.
	UpdateCampaign(id, name string) (Campaign, error)

	// GetCampaigns API provides the ability to retrieve a single campaign for the
	// numerical Campaign ID provided in the URL.
	GetCampaign(campaignID string) (Campaign, error)

	// This API provides the ability to retrieve the segments for the numerical
	// Campaign ID provided in the URL.
	GetSegments(campaignID string) ([]Segment, error)

	// This API is used to create a new segment by providing a JSON definition of the segment.
	CreateSegment(name, campaignID string) (Segment, error)

	// This API is used to update an existing segment by providing a JSON definition
	// of the segment with the modifications. If the segment is successfully updated
	// an HTTP 200 code and response, as shown below, is returned.
	UpdateSegment(name, campaignID, segmentID string) (Segment, error)

	// This API provides the ability to retrieve a single segment for the numerical
	// Segment ID provided in the URL.
	GetSegment(campaignID, segmentID string) (Segment, error)

	// This API provides the ability to retrieve the entire list of trackers for the
	// numerical App ID provided in the URL.
	GetTrackers(query string) ([]Tracker, error)

	// This API is used to update an existing tracker by providing a JSON definition
	// of the tracker with modifications. If the tracker is successfully updated an
	// HTTP 200 code and response, as shown below, is returned.
	UpdateTracker(updates Tracker) (Tracker, error)

	// This API is used to delete an existing tracker by providing the numerical
	// Tracker ID. If the tracker is deleted an HTTP 200 response will be returned,
	// otherwise another HTTP code and message detailing the error will be returned.
	DeleteTracker(trackerID string) error

	// This API is used to create a new tracker by providing a JSON definition of
	// the tracker.
	CreateTracker(name, trackerType, networkID, destinationURL, deeplinkURL, campaignID, segmentID, priceType string, priceValue float32, allowPublisherView bool, events []string, clickURLCustomParams []interface{}) (Tracker, error)

	// This API provides the ability to retrieve the tracker overrides for the
	// numerical Override ID provided in the URL.
	GetTrackerOverrides(trackerID string) (GetOverridesResponse, error)

	// This API provides the ability to create tracker overrides for the numerical
	// Tracker ID provided in the URL.
	PostTrackerOverrides(trackerID string, overrides PostOverridesRequest) error
}
