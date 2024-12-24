package service

import (
	"context"

	m01_pb "github.com/rishad004/bw01_proto-files/microservice-01"
	pb "github.com/rishad004/bw01_proto-files/microservice-02"
)

type Service struct {
	M01 m01_pb.Micro01Client
	pb.UnimplementedMicro02Server
}

func NewService(m02 m01_pb.Micro01Client) *Service {
	return &Service{M01: m02}
}

func (s *Service) Method01(c context.Context, req *pb.Data) (*pb.Details, error) {
	res, err := s.M01.FromMethod(context.Background(), &m01_pb.Empty{})
	if err != nil {
		return nil, err
	}

	f := &pb.Details{}
	for _, v := range res.Users {
		k := &pb.User{
			Name: v.Name,
		}
		f.Users = append(f.Users, k)
	}

	return f, nil
}

func (s *Service) Method02(c context.Context, req *pb.Data) (*pb.Details, error) {
	res, err := s.M01.FromMethod(context.Background(), &m01_pb.Empty{})
	if err != nil {
		return nil, err
	}

	f := &pb.Details{}
	for _, v := range res.Users {
		k := &pb.User{
			Name: v.Name,
		}
		f.Users = append(f.Users, k)
	}

	return f, nil
}
