package reportsclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client struct
type Client struct {
	AppGuid        string
	APIKey         string
	DeliveryMethod []string
	TimeZone       string
	EmailList      []string

	machine *http.Client
}

// NewClient returns a Client struct
func NewClient(AppGuid, APIKey string, options ...func(*Client) error) (Client, error) {
	var c Client

	c.AppGuid = AppGuid
	c.APIKey = APIKey

	for _, option := range options {
		err := option(&c)
		if err != nil {
			return c, err
		}
	}

	return c, nil
}

type QueryAuthorizedResponse struct {
	Status      string   `json:"status"`
	ValidFields []string `json:"valid_fields"`
}

type QueryReportTemplateResponse struct {
	Status         string `json:"status"`
	TemplateValues []struct {
		Name     string `json:"name"`
		Category string `json:"category"`
		Traffic  struct {
			Click bool `json:"click"`
		} `json:"traffic"`
		Grouping  string `json:"grouping"`
		Filtering struct {
			Country bool `json:"country"`
			Network bool `json:"network"`
			Tracker bool `json:"tracker"`
		} `json:"filtering"`
		TrafficIncludes struct {
			Attribution         bool `json:"attribution"`
			CustomParameters    bool `json:"custom_parameters"`
			IdentityLink        bool `json:"identity_link"`
			TrafficVerification bool `json:"traffic_verification"`
		} `json:"traffic_includes"`
		TimeSeries     string `json:"time_series"`
		NetworkEnabled int    `json:"network_enabled"`
		AdminOnly      int    `json:"admin_only"`
	} `json:"template_values"`
}

type QueryColumnTemplateResponse struct {
	Status         string `json:"status"`
	TemplateValues []struct {
		APIKey           string      `json:"api_key"`
		AppID            int         `json:"app_id"`
		ReportType       string      `json:"report_type"`
		ReportSection    string      `json:"report_section"`
		SectionOrder     int         `json:"section_order"`
		ColumnsSelected  []string    `json:"columns_selected"`
		ColumnsAvailable interface{} `json:"columns_available"`
	} `json:"template_values"`
}

type UpdateColumnTemplateRequest struct {
	APIKey           string   `json:"api_key"`
	AppID            int      `json:"app_id"`
	ColumnsSelected  []string `json:"columns_selected,omitempty"`
	ReportSection    string   `json:"report_section"`
	ReportType       string   `json:"report_type"`
	SectionOrder     int      `json:"section_order"`
	ColumnsAvailable []string `json:"columns_available,omitempty"`
}

// QueryGrouping returns the possible report groupings allowed
func (c Client) QueryGrouping() (QueryAuthorizedResponse, error) {

	var r QueryAuthorizedResponse

	endpoint := "https://reporting.api.kochava.com/v1.4/grouping"

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return r, err
	}

	res, err := c.machine.Do(req)
	if err != nil {
		return r, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}

// QueryFiltering returns the possible filters available for reports
func (c Client) QueryFiltering() (QueryAuthorizedResponse, error) {

	var r QueryAuthorizedResponse

	endpoint := "https://reporting.api.kochava.com/v1.4/filtering"

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return r, err
	}

	res, err := c.machine.Do(req)
	if err != nil {
		return r, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}

	return r, nil

}

// QueryTimezones returns the different timezones that are allowed for reports
func (c Client) QueryTimezones() (QueryAuthorizedResponse, error) {

	var r QueryAuthorizedResponse

	endpoint := "https://reporting.api.kochava.com/v1.4/timezones"

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return r, err
	}

	res, err := c.machine.Do(req)
	if err != nil {
		return r, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}

	return r, nil

}

// QueryReportTemplates returns the different available report templates
func (c Client) QueryReportTemplates() (QueryReportTemplateResponse, error) {

	var r QueryReportTemplateResponse

	endpoint := "https://reporting.api.kochava.com/v1.4/templates"

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return r, err
	}

	res, err := c.machine.Do(req)
	if err != nil {
		return r, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}

	return r, nil

}

// QueryColumnTemplate returns the current settings for which columns will be
// returned for different report types
func (c Client) QueryColumnTemplate() (QueryColumnTemplateResponse, error) {

	var r QueryColumnTemplateResponse

	endpoint := "https://reporting.api.kochava.com/v1.4/reportcolumns"

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return r, err
	}

	res, err := c.machine.Do(req)
	if err != nil {
		return r, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}

	return r, nil

}

// SaveTemplateColumns
func (c Client) SaveTemplateColumns(body UpdateColumnTemplateRequest) error {

	reqBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	endpoint := "https://reporting.api.kochava.com/v1.4/reportcolumns/update"

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	res, err := c.machine.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode < 300 && res.StatusCode > 199 {
		return nil
	}

	return fmt.Errorf("%v", res.Status)

}

func (c Client) sendRequest() {

}
