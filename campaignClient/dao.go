package campaignclient

// aPIAccessor interface allows for interfacing with the Kochava Premium Publisher
// API; it creates apps and campaigns in the Kochava system. aPIAccessor has
// two structs that satisfy it: aPIA is a working accessor; aPIA_Fake is a test
// accessor
type aPIAccessor interface {

	// CreateApp creates an app
	createApp() error

	// CreateCampaign creates one or more campaigns
	createCampaign(r ...campaignRequest) (campaignResponse, error)

	// CreateSegment creates a segment for the tracker
	createSegment(campaignID string, r ...segmentRequest) (segmentResponse, error)

	// CreateTracker creates a tracker in the UI; this is the primary item to create
	// Creating a campaign and segment are necessary in order to create a tracker
	createTracker(r ...trackerRequest) (trackerResponse, error)

	// RetrieveTrackers retrieves a list of campaigns
	retrieveTrackers() ([]trackerResponse, error)

	// VerifyCreated verifies the campaigns created match the templates retrieved
	verifyCreated() error
}
