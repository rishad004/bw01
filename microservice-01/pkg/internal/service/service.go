package service

import (
	"context"

	pb "github.com/rishad004/bw01_proto-files/microservice-01"
	m02_pb "github.com/rishad004/bw01_proto-files/microservice-02"
)

type Service struct {
	repo Repo
	M02  m02_pb.Micro02Client
	pb.UnimplementedMicro01Server
}

func NewService(repo Repo, m02 m02_pb.Micro02Client) *Service {
	return &Service{repo: repo, M02: m02}
}

func (s *Service) UserCreate(c context.Context, req *pb.Details) (*pb.Get, error) {
	id, err := s.repo.UserCreate(req.Name, req.Email)
	if err != nil {
		return nil, err
	}

	return &pb.Get{Id: id}, nil
}

func (s *Service) UserFetch(c context.Context, req *pb.Get) (*pb.Details, error) {
	user, err := s.repo.UserFetch(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Details{Id: req.Id, Name: user.Name, Email: user.Email}, nil
}

func (s *Service) UserUpdate(c context.Context, req *pb.Details) (*pb.Empty, error) {
	if err := s.repo.UserUpdate(req.Id, req.Name, req.Email); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *Service) UserDelete(c context.Context, req *pb.Get) (*pb.Empty, error) {
	if err := s.repo.UserDelete(req.Id); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *Service) FromMethod(c context.Context, req *pb.Empty) (*pb.Data, error) {
	data, err := s.repo.FromMethod()
	if err != nil {
		return nil, err
	}

	return data, nil
}
