package storage

import (
	"app/main/models"
	"app/main/utils"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	ReportsCollection = "user_reports"
)

func (s *ReportUsersService) ReadValue(collection string, key string, value uint64) (*models.GetReportMongoModel, error) {

	env := utils.Env()

	coll := s.cli.Database(env.Value("MONGO_INITDB_DATABASE")).Collection(collection)

	var getModel models.GetReportMongoModel
	err := coll.FindOne(context.TODO(), bson.D{{Key: key, Value: value}}).Decode(&getModel)
	return &getModel, err
}

func (s *ReportUsersService) WriteValue(collection string, model *models.Report) error {

	env := utils.Env()

	coll := s.cli.Database(env.Value("MONGO_INITDB_DATABASE")).Collection(collection)

	putModel := models.PutReportMongoModel{
		UserId:    model.UserId,
		Document:  model.Document,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := coll.InsertOne(context.TODO(), putModel)
	if err != nil {
		return err
	}

	log.Printf("inserted document ID %v\n", result.InsertedID)
	return nil
}

func (s *ReportUsersService) GetReportFromDatabase(userId uint64) (*models.Report, error) {

	model, err := s.ReadValue(ReportsCollection, "user_id", userId)
	if err != nil {
		return &models.Report{}, err
	}

	return &models.Report{
		UserId:     model.UserId,
		Document:   model.Document,
		Created_at: model.CreatedAt,
		Updated_at: model.UpdatedAt,
	}, nil
}

func (s *ReportUsersService) AddReportToDatabase(report *models.Report) (*models.Report, error) {

	err := s.WriteValue(ReportsCollection, report)
	return report, err
}
