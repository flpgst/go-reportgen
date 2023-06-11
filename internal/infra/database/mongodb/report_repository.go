package mongodb

import (
	"context"

	"github.com/flpgst/go-reportgen/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type reportFilter struct {
	ReportName string `bson:"reportName,omitempty"`
	Date       string `bson:"date,omitempty"`
}

type ReportRepository struct {
	Db *mongo.Database
}

func NewReportRepository(db *mongo.Database) *ReportRepository {
	return &ReportRepository{
		Db: db,
	}
}

func (r *ReportRepository) Save(report *entity.Report) error {
	reportsCollection := r.Db.Collection("reports")

	doc := bson.M{
		"reportName": report.ReportName,
		"date":       report.Date,
	}

	_, err := reportsCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReportRepository) GetReport(name, date string) (*entity.Report, error) {
	reportsCollection := r.Db.Collection("reports")

	filter := reportFilter{
		ReportName: name,
		Date:       date,
	}

	var report entity.Report
	err := reportsCollection.FindOne(context.TODO(), filter).Decode(&report)
	if err != nil {
		return nil, err
	}
	return &report, nil
}
