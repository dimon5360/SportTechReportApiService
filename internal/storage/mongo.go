package storage

import (
	"app/main/utils"
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ReportUsersService) ReadValue(connection string, key string, value string) (string, error) {

	env := utils.Env()

	coll := s.cli.Database(env.Value("MONGO_INITDB_DATABASE")).Collection(connection)

	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{Key: key, Value: value}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return fmt.Sprintf("No document was found with the title %s\n", key), err
	}

	if err != nil {
		return "", err
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\n", jsonData), nil
}

func (s *ReportUsersService) WriteValue(connection string, key string, value string) (string, error) {

	return "", nil
}
