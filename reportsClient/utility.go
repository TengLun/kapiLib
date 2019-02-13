package reportsClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Error is a standard string value
const Error = "Error"

// CheckStatusScheduledReport returns the status of a scheduled report
func (c Client) CheckStatusScheduledReport(reqBody UtilityRequest) (ReportStatusResponse, error) {
	var r ReportStatusResponse

	body, err := json.Marshal(reqBody)
	if err != nil {
		return r, err
	}

	endpoint := "https://reporting.api.kochava.com/v1.4/schedule/progress"

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		return r, err
	}

	res, err := c.machine.Do(req)
	if err != nil {
		return r, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, err
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return r, fmt.Errorf("%v", res.Status)
	}

	err = json.Unmarshal(resBody, &r)
	if err != nil {
		return r, err
	}

	if r.Status == Error {
		return r, fmt.Errorf("%v", r.StatusDetail)
	}

	return r, nil
}

// UpdateScheduledReportDistributionListRequest is exactly what it says
type UpdateScheduledReportDistributionListRequest struct {
	APIKey  string   `json:"api_key"`
	AppGUID string   `json:"app_guid"`
	Token   string   `json:"token"`
	Notify  []string `json:"notify"`
}

func (c Client) UpdateScheduledReportDistributionList(reqBody UtilityRequest) (ReportRequestResponse, error) {
	// endpoint := "https://reporting.api.kochava.com/v1.4/schedule/detail"
	// endpoint := "https://reporting.api.kochava.com/v1.4/schedule/summary"
	// endpoint := "https://reporting.api.kochava.com/v1.4/schedule/custom"
}

func (c Client) DeleteScheduledReport(reqBody UtilityRequest) error {

	var r ReportStatusResponse

	endpoint := "https://reporting.api.kochava.com/v1.4/schedule/delete"

	body, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	res, err := c.machine.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("%v", res.Status)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, &r)
	if err != nil {
		return err
	}

	if r.Status == Error {
		return fmt.Errorf("%v", r.StatusDetail)
	}

	return nil
}

func (c Client) ScheduledReportsByAPIKey(reqBody UtilityRequest) (ReportsByAPIResponse, error) {

	endpoint := "https://reporting.api.kochava.com/v1.4/schedule/tokens"

	var r ReportsByAPIResponse

	body, err := json.Marshal(reqBody)
	if err != nil {
		return r, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		return r, err
	}

	res, err := c.machine.Do(req)
	if err != nil {
		return r, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return r, fmt.Errorf("%v", res.Status)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(resBody, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (c Client) CheckRequestedReportStatus(reqBody UtilityRequest) (ReportRequestResponse, error) {
	endpoint := "https://reporting.api.kochava.com/v1.4/progress"

	return c.requestReport(reqBody, endpoint)
}

func (c Client) DeleteQueuedReport(reqBody UtilityRequest) (ReportRequestResponse, error) {
	endpoint := "https://reporting.api.kochava.com/v1.4/delete"
	return c.requestReport(reqBody, endpoint)
}

type RequestAppListResponse []struct {
	Status      string `json:"status"`
	GUID        string `json:"guid"`
	AppName     string `json:"app_name"`
	AppID       string `json:"app_id"`
	AccountName string `json:"account_name"`
	Platform    string `json:"platform"`
}

func (c Client) RequestAppList(reqBody UtilityRequest) (ReportRequestResponse, error) {
	endpoint := "https://reporting.api.kochava.com/v1.4/getapps"

	return c.requestReport(reqBody, endpoint)
}

type ReportsByAPIResponse []struct {
	Status          string `json:"status"`
	ReportToken     string `json:"report_token"`
	StatusDate      string `json:"status_date"`
	NextScheduledOn string `json:"next_scheduled_on"`
	ReportType      string `json:"report_type"`
	Frequency       string `json:"frequency"`
	RunOn           string `json:"run_on"`
	Progress        string `json:"progress"`
	Report          string `json:"report"`
	StatusDetail    string `json:"status_detail"`
}

func (c Client) RequestReportsByAPIKey(reqBody UtilityRequest) (ReportRequestResponse, error) {
	endpoint := "https://reporting.api.kochava.com/v1.4/tokens"

	return c.requestReport(reqBody, endpoint)
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