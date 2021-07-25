package services

import (
	"context"
	"grpc-go/modules/pb"
)

type service struct {
	PushCollection                 []float64
	TotalSumCollection, CurrentAvg float64
	StreamCountConn, PushCountConn int
}

//Service _
type Service interface {
	Push(ctx context.Context, input *pb.PushReq) (*pb.Empty, error)
	GetAvgAlarmStream(input *pb.GetAvgAlarmStreamReq, stream pb.AvgAlarmService_GetAvgAlarmStreamServer) error
}

//New _
func New() Service {
	return &service{}
}
