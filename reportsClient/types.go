package reportsClient

type ReportRequest struct {
	APIKey           string   `json:"api_key"`
	AppGUID          string   `json:"app_guid"`
	TimeStart        string   `json:"time_start"`
	TimeEnd          string   `json:"time_end"`
	Traffic          []string `json:"traffic"`
	TrafficGrouping  []string `json:"traffic_grouping"`
	TimeSeries       string   `json:"time_series"`
	TimeZone         string   `json:"time_zone"`
	TrafficFiltering struct {
		Network        []string `json:"network"`
		ExcludeCountry []string `json:"exclude_country"`
	} `json:"traffic_filtering"`
	DeliveryMethod []string `json:"delivery_method"`
	DeliveryFormat string   `json:"delivery_format"`
	Notify         []string `json:"notify"`
	ColumnsOrder   []string `json:"columns_order"`
}

type RequestStatus struct {
	Status      string `json:"status"`
	ReportToken string `json:"report_token"`
}
