package storage

import (
	"app/main/utils"
	"context"
	"fmt"
	"log"

	proto "github.com/dimon5360/SportTechProtos/gen/go/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReportUsersService struct {
	proto.UnimplementedReportUsersServiceServer

	cli *mongo.Client
}

func CreateService() *ReportUsersService {

	return &ReportUsersService{}
}

func (s *ReportUsersService) Init() {

	env := utils.Env()
	uri := fmt.Sprintf("mongodb://%s:%s@%s/?maxPoolSize=%s",
		env.Value("MONGO_INITDB_ROOT_USERNAME"),
		env.Value("MONGO_INITDB_ROOT_PASSWORD"),
		env.Value("MONGO_DB_HOST"),
		env.Value("MONGO_INITDB_MAX_POOL_SIZE"),
	)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	s.cli = client
}
