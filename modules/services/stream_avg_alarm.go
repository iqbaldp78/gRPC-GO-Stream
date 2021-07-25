package services

import (
	"grpc-go/modules/pb"
	"log"
	"time"
)

func (s *service) GetAvgAlarmStream(input *pb.GetAvgAlarmStreamReq, stream pb.AvgAlarmService_GetAvgAlarmStreamServer) error {
	log.Printf("running GetAvgAlarmStream  with treshold %v ...", input.Threshold)

	s.StreamCountConn++
	if s.StreamCountConn >= 100 {
		log.Fatal("Max 100 connections")
	}

	for {
		time.Sleep(time.Second)
		if checkingTreshold(s.CurrentAvg, float64(input.Threshold)) {
			recv := &pb.GetAvgAlarmStreamRes{Avg: float32(s.CurrentAvg)}
			// Send msg
			err := stream.Send(recv)
			if err != nil {
				log.Println(err.Error())
			}
		}

		select {
		case <-stream.Context().Done():
			log.Println("1 client disconnected")
			s.StreamCountConn--
			if s.StreamCountConn == 0 {
				log.Println("No active Clients")
			}
			return nil
		default:

		}

	}
}

func checkingTreshold(curr, treshold float64) bool {
	if curr > treshold {
		return true
	}
	return false
}
