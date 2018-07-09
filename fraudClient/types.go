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
	View      string         `json:"view"`
	FraudType string         `json:"fraudType"`
	AccountID string         `json:"accountId"`
	StartDate string         `json:"startDate"`
	EndDate   string         `json:"endDate"`
	Format    string         `json:"format"`
	Filters   []FilterObject `json:"filters,omitempty"`
}

// FilterObject for modifying requests
type FilterObject struct {
	Dimension string   `json:"dimension"`
	Values    []string `json:"values"`
	Modifier  string   `json:"modifier"`
}

const (
	// AdStacking Constant to ensure proper spelling per API Requirements
	AdStacking = "adStacking"
	// AnonymousInstall Constant to ensure proper spelling per API Requirements
	AnonymousInstall = "anonymousInstall"
	// DeviceHighClickVolume Constant to ensure proper spelling per API Requirements
	DeviceHighClickVolume = "deviceHighClickVolume"
	// DoubleHashAttribution Constant to ensure proper spelling per API Requirements
	DoubleHashAttribution = "doubleHashAttribution"
	// GeoOutliers Constant to ensure proper spelling per API Requirements
	GeoOutliers = "geoOutliers"
	// InstallReceiptVerification Constant to ensure proper spelling per API Requirements
	InstallReceiptVerification = "installReceiptVerification"
	// FraudSummary Constant to ensure proper spelling per API Requirements
	FraudSummary = "fraudSummary"
	// IPHighClick Constant to ensure proper spelling per API Requirements
	IPHighClick = "ipHighClick"
	// MTTIOutliers Constant to ensure proper spelling per API Requirements
	MTTIOutliers = "mttiOutliers"
	// PlatformDifference Constant to ensure proper spelling per API Requirements
	PlatformDifference = "platformDifference"
	// PurchaseReceiptVerification Constant to ensure proper spelling per API Requirements
	PurchaseReceiptVerification = "purchaseReceiptVerification"
	// TTIOutlier Constant to ensure proper spelling per API Requirements
	TTIOutlier = "ttiOutlier"
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

// KFResponse is a generic response struct
type KFResponse struct {
	MetaData struct {
		Headers []string `json:"headers"`
	} `json:"metaData"`
	Data []struct {
		AppName         string `json:"appName,omitempty"`
		AppID           string `json:"appId,omitempty"`
		NetworkName     string `json:"networkName,omitempty"`
		NetworkID       string `json:"networkId,omitempty"`
		ClickCt         int    `json:"clickCt,omitempty"`
		SameAcctClickCt int    `json:"sameAcctClickCt,omitempty"`
		DiffAcctClickCt int    `json:"diffAcctClickCt,omitempty"`
		InstallCt       int    `json:"installCt,omitempty"`
	} `json:"data"`
}
