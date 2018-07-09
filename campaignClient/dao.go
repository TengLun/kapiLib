package campaignclient

// aPIAccessor interface allows for interfacing with the Kochava Premium Publisher
// API; it creates apps and campaigns in the Kochava system. aPIAccessor has
// two structs that satisfy it: aPIA is a working accessor; aPIA_Fake is a test
// accessor
type aPIAccessor interface {

	// GetCampaigns API provides the ability to retrieve the entire list of campaigns
	// from the numerical App ID provided in the URL.
	GetCampaigns(stats string) ([]GetCampaignsResponse, error)

	// CreateCampaign API is used to create a new campaign by providing a JSON
	// definition of the campaign.
	CreateCampaign()

	// UpdateCampaign API is used to update an existing campaign by providing a JSON
	// definition of the campaign with the modifications. If the campaign is
	// successfully updated an HTTP 200 code and response, as shown below, is
	// returned.
	UpdateCampaign()

	// GetCampaigns API provides the ability to retrieve a single campaign for the
	// numerical Campaign ID provided in the URL.
	GetCampaign()

	// This API provides the ability to retrieve the segments for the numerical
	// Campaign ID provided in the URL.
	GetSegments()

	// This API is used to create a new segment by providing a JSON definition of the segment.
	CreateSegment()

	// This API is used to update an existing segment by providing a JSON definition
	// of the segment with the modifications. If the segment is successfully updated
	// an HTTP 200 code and response, as shown below, is returned.
	UpdateSegment()

	// This API provides the ability to retrieve a single segment for the numerical
	// Segment ID provided in the URL.
	GetSegment()

	// This API provides the ability to retrieve the entire list of trackers for the
	// numerical App ID provided in the URL.
	GetTrackers()

	// This API is used to update an existing tracker by providing a JSON definition
	// of the tracker with modifications. If the tracker is successfully updated an
	// HTTP 200 code and response, as shown below, is returned.
	UpdateTracker()

	// This API is used to delete an existing tracker by providing the numerical
	// Tracker ID. If the tracker is deleted an HTTP 200 response will be returned,
	// otherwise another HTTP code and message detailing the error will be returned.
	DeleteTracker()

	// This API is used to create a new tracker by providing a JSON definition of
	// the tracker.
	CreateTracker()

	// This API provides the ability to retrieve the tracker overrides for the
	// numerical Override ID provided in the URL.
	GetTrackerOverrides()

	// This API provides the ability to create tracker overrides for the numerical
	// Tracker ID provided in the URL.
	PostTrackerOverrides()
}
