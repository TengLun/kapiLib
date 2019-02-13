package reportsClient

func (c Client) RequestSummaryReport() {
	endpoint := "https://reporting.api.kochava.com/v1.4/summary"
}

func (c Client) RequestDetailReport() {
	endpoint := "https://reporting.api.kochava.com/v1.4/detail"
}

// UtilityRequest is a request for utility functions, such as checking on the
// status of a report, deleting a report, and more.
type UtilityRequest struct {
	APIKey  string `json:"api_key"`
	AppGUID string `json:"app_guid"`
	Token   string `json:"token"`
}

type ReportStatusResponse struct {
	Status        string `json:"status"`
	StatusDetail  string `json:"status_detail"`
	StatusDate    string `json:"status_date"`
	Progress      string `json:"progress"`
	FileSize      int    `json:"file_size"`
	Report        string `json:"report"`
	ReportType    string `json:"report_type"`
	ReportRequest struct {
		APIKey           string   `json:"api_key"`
		AppGUID          string   `json:"app_guid"`
		TimeStart        string   `json:"time_start"`
		TimeEnd          string   `json:"time_end"`
		Traffic          []string `json:"traffic"`
		TrafficIncluding []string `json:"traffic_including"`
		TimeZone         string   `json:"time_zone"`
		DeliveryFormat   string   `json:"delivery_format"`
		DeliveryMethod   []string `json:"delivery_method"`
		Notify           []string `json:"notify"`
	} `json:"report_request"`
}

func (c Client) CheckRequestedReportStatus() {
	endpoint := "https://reporting.api.kochava.com/v1.4/progress"
}

func (c Client) DeleteQueuedReport() {
	endpoint := "https://reporting.api.kochava.com/v1.4/delete"

}

type RequestAppListResponse []struct {
	Status      string `json:"status"`
	GUID        string `json:"guid"`
	AppName     string `json:"app_name"`
	AppID       string `json:"app_id"`
	AccountName string `json:"account_name"`
	Platform    string `json:"platform"`
}

func (c Client) RequestAppList() {
	endpoint := "https://reporting.api.kochava.com/v1.4/getapps"
}

type RequestedReportsByAPIResponse []struct {
	Status      string `json:"status"`
	ReportToken string `json:"report_token"`
	StatusDate  string `json:"status_date"`
	Progress    string `json:"progress"`
	Report      string `json:"report"`
	ReportType  string `json:"report_type"`
}

func (c Client) RequestReportsByAPIKey() {
	endpoint := "https://reporting.api.kochava.com/v1.4/tokens"
}

type ScheduleReportRequest struct {
	APIKey           string   `json:"api_key"`
	AppGUID          string   `json:"app_guid"`
	Traffic          []string `json:"traffic"`
	TrafficGrouping  []string `json:"traffic_grouping"`
	TimeSeries       string   `json:"time_series"`
	TrafficFiltering struct {
		Country []string `json:"country"`
	} `json:"traffic_filtering"`
	DeliveryFormat string   `json:"delivery_format"`
	DeliveryMethod []string `json:"delivery_method"`
	Notify         []string `json:"notify"`
	Frequency      string   `json:"frequency"`
	RunOn          string   `json:"run_on"`
	Delay          string   `json:"delay"`
	PreviousTime   string   `json:"previous_time"`
	ColumnsOrder   []string `json:"columns_order"`
}

func (c Client) ScheduleSummaryReport() {
	endpoint := "https://reporting.api.kochava.com/v1.4/schedule/summary"
}

func (c Client) ScheduleDetailReport() {
	endpoint := "https://reporting.api.kochava.com/v1.4/schedule/detail"
}

func (c Client) CheckStatusScheduledReport() {
	endpoint := "https://reporting.api.kochava.com/v1.4/schedule/progress"
}

type UpdateScheduledReportDistributionListRequest struct {
	APIKey  string   `json:"api_key"`
	AppGUID string   `json:"app_guid"`
	Token   string   `json:"token"`
	Notify  []string `json:"notify"`
}

func (c Client) UpdateScheduledReportDistributionList() {
	// endpoint := "https://reporting.api.kochava.com/v1.4/schedule/detail"
	// endpoint := "https://reporting.api.kochava.com/v1.4/schedule/summary"
	// endpoint := "https://reporting.api.kochava.com/v1.4/schedule/custom"
}

func (c Client) DeleteScheduledReport() {
	endpoint := "https://reporting.api.kochava.com/v1.4/schedule/delete"
}

func (c Client) ScheduledReportsByAPIKey() {
	endpoint := "https://reporting.api.kochava.com/v1.4/schedule/tokens"
}
