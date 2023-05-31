package mongodb

import (
	"context"
	"fmt"

	"github.com/flpgst/go-reportgen/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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

	result, err := reportsCollection.InsertOne(context.TODO(), &report)
	if err != nil {
		return err
	}
	fmt.Println(result.InsertedID)
	return nil
}

func (r *ReportRepository) GetReport(name, date string) (*entity.Report, error) {
	reportsCollection := r.Db.Collection("reports")
	filter := bson.M{"reportName": name, "date": date}
	var report entity.Report
	err := reportsCollection.FindOne(context.TODO(), filter).Decode(&report)
	if err != nil {
		return nil, err
	}
	return &report, nil
}
