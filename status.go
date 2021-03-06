package igdb

import "github.com/pkg/errors"

//go:generate gomodifytags -file $GOFILE -struct Status -add-tags json -w

// Status contains the usage report for the user's API key along with other
// metadata.
// For more information visit: https://api-docs.igdb.com/#api-status
type Status struct {
	Authorized   bool        `json:"authorized"`
	Plan         string      `json:"plan"`
	UsageReports UsageReport `json:"usage_reports"`
}

//go:generate gomodifytags -file $GOFILE -struct UsageReport -add-tags json -w

// UsageReport contains information and statistics for the the current user's
// API usage in the current period.
type UsageReport struct {
	Metric       string `json:"metric"`
	Period       string `json:"period"`
	PeriodStart  string `json:"period_start"`
	PeriodEnd    string `json:"period_end"`
	MaxValue     int    `json:"max_value"`
	CurrentValue int    `json:"current_value"`
}

// Status returns a usage report for the user's API key. It shows stats such as
// requests made in the current period and when that period ends.
// For more information visit: https://api-docs.igdb.com/#api-status
func (c *Client) Status() (*Status, error) {
	var stat []*Status

	err := c.get(EndpointStatus, &stat, []Option{}...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get API status")
	}

	return stat[0], nil
}
