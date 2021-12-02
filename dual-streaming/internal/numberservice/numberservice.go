package numberservice

import (
	"fmt"
	"io"

	numberservice "github.com/palash287gupta/learn_and_kount_devops_2021_12_01_grpc4/dual-streaming/numberservice/gen/go"
	"github.com/sirupsen/logrus"
)

type NumberService struct {
	Log *logrus.Logger
	numberservice.UnimplementedNumberServiceServer
}

func (ns *NumberService) GetSquares(numserver numberservice.NumberService_GetSquaresServer) error {
	for {
		reqNum, err := numserver.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			ns.Log.Errorf("failed to receive number", err)
			return err
		}
		fmt.Println("request : ", reqNum.GetNum())
		err = numserver.Send(&numberservice.GetSquaresResponse{
			Num: reqNum.GetNum() * reqNum.GetNum(),
		})
		if err != nil {
			ns.Log.Errorf("failed to send response")
			return err
		}
	}
}
