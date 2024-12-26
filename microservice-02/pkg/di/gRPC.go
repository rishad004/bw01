package di

import (
	"log"
	"net"

	"github.com/rishad004/bw01/microservice-02/pkg/internal/service"
	m01_pb "github.com/rishad004/bw01_proto-files/microservice-01"
	pb "github.com/rishad004/bw01_proto-files/microservice-02"
	"google.golang.org/grpc"
)

func InitgRPC() {

	ConnM01, m01Svc := Micro01Conn()
	defer ConnM01.Close()

	src := service.NewService(m01Svc)

	if err := Micro02Start(src); err != nil {
		log.Fatal(err)
	}
}

func Micro02Start(src *service.Service) error {
	g := grpc.NewServer()
	pb.RegisterMicro02Server(g, src)

	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		return err
	}

	log.Println("microservice02-service server listening on port :50052")
	if err := g.Serve(listen); err != nil {
		return err
	}

	return nil
}

func Micro01Conn() (*grpc.ClientConn, m01_pb.Micro01Client) {
	connM01, err := grpc.Dial("bw01-microservice-01-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to microservice02 service:", err)
	}

	m01Svc := m01_pb.NewMicro01Client(connM01)

	return connM01, m01Svc
}
