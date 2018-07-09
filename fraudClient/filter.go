package fraudclient

const (
	// AccountID dimension for filter creation
	AccountID = "accountId"
	// AppID dimension for filter creation
	AppID = "appId"
	// NetworkID dimension for filter creation
	NetworkID = "networkId"
	// SiteID dimension for filter creation
	SiteID = "siteId"
	// TrackerID dimension for filter creation
	TrackerID = "trackerId"
	// In modifier for filter creation
	In = "IN"
	// NotIn modifier for filter creation
	NotIn = "NOT IN"
	// Modifier_GreaterThan  = "greaterThan"
	// Modifier_LessThan     = "lessThan"
	// Modifier_InstallCount = "installct"
	// Modifier_ClickCount   = "clickct"
	// Modifier_Auto         = "auto"
)

// Filter returns a filter for a request
func Filter(dimension string, modifier string, values ...string) FilterObject {

	f := FilterObject{
		Dimension: dimension,
		Modifier:  modifier,
		Values:    values,
	}

	return f
}
