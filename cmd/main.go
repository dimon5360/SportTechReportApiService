package main

import (
	"app/main/internal/storage"
	"app/main/internal/utils"
	"fmt"
	"net"

	"github.com/dimon5360/SportTechProtos/gen/go/proto"
	"google.golang.org/grpc"
)

const (
	configPath = "../SportTechDockerConfig/"
	serviceEnv = ".env"
	mongoEnv   = configPath + "mongo.env"
)

func main() {

	utils.Env().Load(serviceEnv, mongoEnv)

	fmt.Println("SportTech report API service v." + utils.Env().Value("SERVICE_VERSION"))

	service := storage.CreateService()
	service.Init()

	lis, err := net.Listen("tcp", utils.Env().Value("REPORT_GRPC_HOST"))
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterReportUsersServiceServer(grpcServer, service)
	grpcServer.Serve(lis)
}
