package entity

import (
	"errors"
)

type Report struct {
	ReportName string
	Date       string
}

func NewReport(reportName, date string) (*Report, error) {
	report := &Report{
		ReportName: reportName,
		Date:       date,
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
