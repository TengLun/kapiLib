package fraudclient

const (
	Dimension_AccountId   = "accountId"
	Dimension_AppId       = "appId"
	Dimension_NetworkId   = "networkId"
	Dimension_SiteId      = "siteId"
	Dimension_TrackerId   = "trackerId"
	Modifier_In           = "IN"
	Modifier_NotIn        = "NOT IN"
	Modifier_GreaterThan  = "greaterThan"
	Modifier_LessThan     = "lessThan"
	Modifier_InstallCount = "installct"
	Modifier_ClickCount   = "clickct"
	Modifier_Auto         = "auto"
)

func CreateFilter(dimension string, modifier string, values ...string) filter {

	filter := filter{
		Dimension: dimension,
		Modifier:  modifier,
		Values:    values,
	}

	return filter
}
