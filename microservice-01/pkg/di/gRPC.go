package di

import (
	"log"
	"net"

	"github.com/rishad004/bw01/microservice-01/pkg/db"
	"github.com/rishad004/bw01/microservice-01/pkg/internal/repository"
	"github.com/rishad004/bw01/microservice-01/pkg/internal/service"
	pb "github.com/rishad004/bw01_proto-files/microservice-01"
	m02_pb "github.com/rishad004/bw01_proto-files/microservice-02"
	"google.golang.org/grpc"
)

func InitgRPC() {

	Db, err := db.PsqlConn()
	if err != nil {
		log.Fatal(err)
	}

	Rd := db.RedisConn()

	ConnM02, m02Svc := Micro02Conn()
	defer ConnM02.Close()

	repo := repository.NewRepository(Db, Rd)
	src := service.NewService(repo, m02Svc)

	if err := Micro01Start(src); err != nil {
		log.Fatal(err)
	}
}

func Micro01Start(src *service.Service) error {
	g := grpc.NewServer()
	pb.RegisterMicro01Server(g, src)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	log.Println("microservice01-service server listening on port :50051")
	if err := g.Serve(listen); err != nil {
		return err
	}

	return nil
}

func Micro02Conn() (*grpc.ClientConn, m02_pb.Micro02Client) {
	connM02, err := grpc.Dial("bw01-microservice-02-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to microservice02 service:", err)
	}

	m02Svc := m02_pb.NewMicro02Client(connM02)

	return connM02, m02Svc
}
