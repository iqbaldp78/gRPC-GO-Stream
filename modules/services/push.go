package services

import (
	"context"
	"grpc-go/modules/pb"
	"log"
)

func (s *service) Push(ctx context.Context, input *pb.PushReq) (*pb.Empty, error) {

	s.PushCountConn++
	if s.PushCountConn >= 100 {
		log.Fatal("Max 100 connections")
	}
	log.Println("count ", s.PushCountConn)
	lengCollection := len(s.PushCollection)

	if lengCollection == 0 {
		lengCollection = 1
	}

	s.TotalSumCollection = s.TotalSumCollection + float64(input.Num)
	s.PushCollection = append(s.PushCollection, float64(input.Num))
	s.CurrentAvg = s.TotalSumCollection / float64(len(s.PushCollection))
	// log.Println(" TotalSumCollection ", s.TotalSumCollection)
	// log.Println(" PushCollection ", s.PushCollection)
	return &pb.Empty{}, nil
}
