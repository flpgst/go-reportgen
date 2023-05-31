package entity

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Report struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ReportName string             `bson:"reportName,omitempty"`
}

func NewReport(reportName string) (*Report, error) {
	report := &Report{
		ReportName: reportName,
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
	return nil
}
