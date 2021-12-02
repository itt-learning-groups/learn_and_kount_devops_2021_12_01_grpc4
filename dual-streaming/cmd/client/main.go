package main

import (
	"context"
	"fmt"
	"io"

	numberservice "github.com/palash287gupta/learn_and_kount_devops_2021_12_01_grpc4/dual-streaming/numberservice/gen/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	log := logrus.New()

	clDial, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to initiate grpc dialer: %v", err)
	}
	defer clDial.Close()

	numberClient := numberservice.NewNumberServiceClient(clDial)
	clientStream, err := numberClient.GetSquares(context.Background())
	if err != nil {
		log.Fatalf("failed to initiate client: %v", err)
	}

	waitc := make(chan struct{})

	// send go routine
	go func() {
		var i int64
		for i = 1; i <= 10; i++ {
			getSquareReq := numberservice.GetSquaresRequest{
				Num: i,
			}
			clientStream.Send(&getSquareReq)
			// time.Sleep(1000 * time.Millisecond)
		}
		clientStream.CloseSend()
	}()

	// receive
	for {
		respMsg, err := clientStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("failed to recieve response (client): %v", err)
			break
		}
		fmt.Println("response : ", respMsg.GetNum())
	}
	close(waitc)
	<-waitc
}
