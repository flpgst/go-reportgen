package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReportInputDTO struct {
	ReportName string `bson:"reportName,omitempty"`
	Date       string `bson:"date,omitempty"`
}

type ReportOutputDTO struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ReportName string             `bson:"reportName,omitempty"`
	Date       string             `bson:"date,omitempty"`
}
