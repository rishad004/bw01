package service

import (
	"github.com/rishad004/bw01/microservice-01/pkg/domain"
	pb "github.com/rishad004/bw01_proto-files/microservice-01"
)

type Repo interface {
	UserCreate(name, email string) (int32, error)
	UserFetch(Id int32) (domain.User, error)
	UserUpdate(id int32, name, email string) error
	UserDelete(Id int32) error
	FromMethod() (*pb.Data, error)
}
