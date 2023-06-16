package entity

import (
	"errors"
)

type Report struct {
	ReportName string
	Date       string
	Header     []string
	Body       []string
	Footer     []string
}

func NewReport(reportName, date string, header, body, footer []string) (*Report, error) {
	report := &Report{
		ReportName: reportName,
		Date:       date,
		Header:     header,
		Body:       body,
		Footer:     footer,
	}
	err := report.isValid()
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (r *Report) isValid() error {
	if r.ReportName == "" {
		return errors.New("report name cannot be empty")
	}
	if r.Date == "" {
		return errors.New("date cannot be empty")
	}
	return nil
}
