package modules

import (
	"context"
	"fmt"
	"grpc-go/modules/pb"
	"grpc-go/modules/services"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedAvgAlarmServiceServer
	services services.Service
}

//ListenGRPC used for run grpc service
func ListenGRPC(s services.Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	serv := grpc.NewServer()
	pb.RegisterAvgAlarmServiceServer(serv, &grpcServer{services: s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) Push(ctx context.Context, input *pb.PushReq) (*pb.Empty, error) {
	return s.services.Push(ctx, input)
}

func (s *grpcServer) GetAvgAlarmStream(input *pb.GetAvgAlarmStreamReq, stream pb.AvgAlarmService_GetAvgAlarmStreamServer) error {
	return s.services.GetAvgAlarmStream(input, stream)
}
