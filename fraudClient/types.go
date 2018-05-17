package fraudclient

// siteID is the request structure for adding a site_id to the blackList
type siteID struct {
	Type            string `json:"type"`
	BlackListSiteID struct {
		SiteID    string `json:"site_id"`    // the id of the site
		NetworkID string `json:"network_id"` // the network id of belonging to the site
		Source    int    `json:"source"`     // should always be 2
		AccountID string `json:"accountId"`  // account_id of the user
		Reason    string `json:"reason"`     // why the site has been put into the blackList
		Score     int    `json:"score"`      // the % that this site_id is true fraud
	} `json:"blacklistSiteId"`
}

// DeviceID is the request structure for adding a device_id to the blackList
type deviceID struct {
	Type            string `json:"type"`
	BlackListDevice struct {
		DeviceIDValue string `json:"deviceIdValue"` // id of the device
		DeviceIDType  string `json:"deviceIdType"`  // type of the device
		Source        int    `json:"source"`        // should always be 2
		AccountID     string `json:"accountId"`     // account_id of the user
		Reason        string `json:"reason"`        // why the site has been put into the blackList
		Score         int    `json:"score"`         // the % that this site_id is true fraud
	} `json:"blacklistDevice"`
}

// IPAddress is the request structure for adding an IP Address to the blackList
type ipAddress struct {
	Type        string `json:"type"`
	BlackListIP struct {
		IPAddress string `json:"ipAddress"` // ip address that will be blackListed
		Source    int    `json:"source"`    // should always be 2
		AccountID string `json:"accountId"` // account_id of the user
		Reason    string `json:"reason"`    // why the ip has been put on the blackList
		Score     int    `json:"score"`     // the % that this ip is true fraud
	} `json:"blacklistIp"`
}

// BlackList of all entries to be sent
type BlackList struct {
	BlackListSiteIDs []siteID
	BlackListDevices []deviceID
	BlackListIPs     []ipAddress
}

// response struct
type addResponse struct {
	Status string `json:"status"`
}

type request struct {
	View      string   `json:"view"`
	FraudType string   `json:"fraudType"`
	AccountID string   `json:"accountId"`
	StartDate string   `json:"startDate"`
	EndDate   string   `json:"endDate"`
	Format    string   `json:"format"`
	Filters   []filter `json:"filters,omitempty"`
}

// Filter object for modifying requests
type filter struct {
	Dimension string   `json:"dimension"`
	Values    []string `json:"values"`
	Modifier  string   `json:"modifier"`
}

const (
	AdStacking                  = "adStacking"
	AnonymousInstall            = "anonymousInstall"
	DeviceHighClickVolume       = "deviceHighClickVolume"
	DoubleHashAttribution       = "doubleHashAttribution"
	GeoOutliers                 = "geoOutliers"
	InstallReceiptVerification  = "installReceiptVerification"
	FraudSummary                = "fraudSummary"
	IPHighClick                 = "ipHighClick"
	MTTIOutliers                = "mttiOutliers"
	PlatformDifference          = "platformDifference"
	PurchaseReceiptVerification = "purchaseReceiptVerification"
	TTIOutlier                  = "ttiOutlier"
)

var fraudEndpointMap = map[string]string{
	"adStacking":                  "adstacking",
	"anonymousInstall":            "anonymousinstall",
	"deviceHighClickVolume":       "devicehighclick",
	"doubleHashAttribution":       "doublehashattrib",
	"geoOutliers":                 "geooutliers",
	"installReceiptVerification":  "installreceipt",
	"fraudSummary":                "fraudsummary",
	"ipHighClick":                 "iphighclick",
	"mttiOutliers":                "mttioutlier",
	"platformDifference":          "platformdiff",
	"purchaseReceiptVerification": "purchasereceipt",
	"ttiOutlier":                  "ttioutlier",
}

/*
  client.Account.List.Apps()
  account.list.Networks()
  account.Data.Apps()
  account.Data.Networks()
  account.Data.SiteIDs()
  account.Data.Trackers()
*/
