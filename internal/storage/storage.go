package storage

import (
	"app/main/internal/models"
	"app/main/internal/utils"
	"context"
	"fmt"
	"log"

	proto "github.com/dimon5360/SportTechProtos/gen/go/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *ReportUsersService) emptyReportResponse(userId uint64) (*proto.ReportResponse, error) {
	return &proto.ReportResponse{
		UserId:    userId,
		Report:    "",
		CreatedAt: &timestamppb.Timestamp{},
		UpdatedAt: &timestamppb.Timestamp{},
	}, nil // at this time no error code just empty report document
}

func (s *ReportUsersService) GetReport(ctx context.Context, req *proto.GetReportRequest) (*proto.ReportResponse, error) {

	report, err := s.GetReportFromDatabase(req.UserId)
	if err != nil {
		return s.emptyReportResponse(report.UserId)
	}

	return &proto.ReportResponse{
		UserId:    report.UserId,
		Report:    report.Document,
		CreatedAt: timestamppb.New(report.Created_at),
		UpdatedAt: timestamppb.New(report.Updated_at),
	}, nil
}

func (s *ReportUsersService) AddReport(ctx context.Context, req *proto.AddReportRequst) (*proto.ReportResponse, error) {

	report := &models.Report{
		UserId:   req.UserId,
		Document: req.Report,
	}

	report, err := s.AddReportToDatabase(report)
	if err != nil {
		return s.emptyReportResponse(report.UserId)
	}

	return &proto.ReportResponse{
		UserId:    report.UserId,
		Report:    report.Document,
		CreatedAt: timestamppb.New(report.Created_at),
		UpdatedAt: timestamppb.New(report.Updated_at),
	}, nil
}
