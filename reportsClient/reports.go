package reportsClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type ReportRequest struct {
	APIKey           string           `json:"api_key"`
	AppGUID          string           `json:"app_guid"`
	TimeStart        string           `json:"time_start"`
	TimeEnd          string           `json:"time_end"`
	Traffic          []string         `json:"traffic"`
	TrafficGrouping  []string         `json:"traffic_grouping"`
	TimeSeries       string           `json:"time_series"`
	TimeZone         string           `json:"time_zone"`
	TrafficFiltering TrafficFiltering `json:"traffic_filtering"`
	DeliveryMethod   []string         `json:"delivery_method"`
	DeliveryFormat   string           `json:"delivery_format"`
	Notify           []string         `json:"notify"`
	ColumnsOrder     []string         `json:"columns_order"`
	Frequency        string           `json:"frequency"`
	RunOn            string           `json:"run_on"`
	Delay            string           `json:"delay"`
	PreviousTime     string           `json:"previous_time"`
}

type RequestStatus struct {
	Status      string `json:"status"`
	ReportToken string `json:"report_token"`
}

// NewReport returns a report request; this method simplifies creating a report
// request object.
//
// If the scheduledReport bool is set to true, a SetDeliverySchedule
// function is required, otherwise an err will be returned
//
// If the scheduledReport is set to false, a SetReportTime
// will be required, otherwise an error will be returned
func (c Client) NewReport(scheduledReport bool, settings ...func(*ReportRequest) error) (ReportRequest, error) {
	var r ReportRequest

	r.DeliveryMethod = []string{"S3link"}

	for _, setting := range settings {
		err := setting(&r)
		if err != nil {
			return r, err
		}
	}

	if scheduledReport == true && (r.TimeStart == "" || r.TimeEnd == "") {
		return r, fmt.Errorf("one-time report requested and timestart or timeend are not defined")
	}

	if r.Frequency == "" || r.RunOn == "" || r.Delay == "" || r.PreviousTime == "" {
		return r, fmt.Errorf("scheduled report requested and frequency, runon, delay, or previoustime are not defined")
	}

	return r, nil
}

// SetReportTime sets the reports time start and time end.
func SetReportTime(timeStart, timeEnd time.Time) func(*ReportRequest) error {
	return func(r *ReportRequest) error {
		r.TimeStart = strconv.FormatInt(timeStart.Unix(), 10)
		r.TimeEnd = strconv.FormatInt(timeEnd.Unix(), 10)
		return nil
	}
}

// SetReportTraffic sets what traffic data will be returned by the report
func SetReportTraffic(traffic []string) func(*ReportRequest) error {
	return func(r *ReportRequest) error {
		r.Traffic = traffic
		return nil
	}
}

// SetReportTrafficGrouping determines how to group a summary report
func SetReportTrafficGrouping(groupings []string) func(*ReportRequest) error {
	return func(r *ReportRequest) error {
		r.TrafficGrouping = groupings
		return nil
	}
}

type TrafficFiltering struct {
	Network        []string `json:"network"`
	ExcludeCountry []string `json:"exclude_country"`
}

func SetTrafficFilters(networkFilter, countryFilter []string) func(*ReportRequest) error {
	return func(r *ReportRequest) error {
		var tf TrafficFiltering
		tf.Network = networkFilter
		tf.ExcludeCountry = countryFilter
		r.TrafficFiltering = tf
		return nil
	}
}

// SetColumnsOrder determines which order the columns will be returned in
func SetColumnsOrder(columnsOrder []string) func(*ReportRequest) error {
	return func(r *ReportRequest) error {
		r.ColumnsOrder = columnsOrder
		return nil
	}
}

func (c Client) RequestSummaryReport(reqBody ReportRequest) (ReportRequestResponse, error) {
	endpoint := "https://reporting.api.kochava.com/v1.4/summary"

	return c.requestReport(reqBody, endpoint)
}

func (c Client) requestReport(reqBody ReportRequest, endpoint string) (ReportRequestResponse, error) {
	var r ReportRequestResponse

	body, err := json.Marshal(reqBody)
	if err != nil {
		return r, nil
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		return r, nil
	}

	res, err := c.machine.Do(req)
	if err != nil {
		return r, nil
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, nil
	}

	err = json.Unmarshal(resBody, &r)
	if err != nil {
		return r, nil
	}

	if res.StatusCode < 200 || res.StatusCode > 300 {
		return r, fmt.Errorf("%v", res.Status)
	}

	if r.Status == "Error" {
		return r, fmt.Errorf("%v", r.Error)
	}

	return r, nil
}

func (c Client) RequestDetailReport(reqBody ReportRequest) (ReportRequestResponse, error) {
	endpoint := "https://reporting.api.kochava.com/v1.4/detail"

	return c.requestReport(reqBody, endpoint)
}

type ReportRequestResponse struct {
	Status      string `json:"status"`
	ReportToken string `json:"report_token"`
	Error       string `json:"error"`
}

func (c Client) ScheduleSummaryReport(reqBody ReportRequest) (ReportRequestResponse, error) {
	endpoint := "https://reporting.api.kochava.com/v1.4/schedule/summary"
	return c.requestReport(reqBody, endpoint)
}

func (c Client) ScheduleDetailReport(reqBody ReportRequest) (ReportRequestResponse, error) {
	endpoint := "https://reporting.api.kochava.com/v1.4/schedule/detail"
	return c.requestReport(reqBody, endpoint)
}
