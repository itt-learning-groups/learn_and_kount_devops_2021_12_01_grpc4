package main

import (
	"net"

	ns "github.com/palash287gupta/learn_and_kount_devops_2021_12_01_grpc4/dual-streaming/internal/numberservice"
	numberservice "github.com/palash287gupta/learn_and_kount_devops_2021_12_01_grpc4/dual-streaming/numberservice/gen/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := logrus.New()

	svrListen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to generate server listener: %v", err)
	}

	nss := ns.NumberService{
		Log: log,
	}

	svr := grpc.NewServer()

	numberservice.RegisterNumberServiceServer(svr, &nss)
	reflection.Register(svr)

	if err := svr.Serve(svrListen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
