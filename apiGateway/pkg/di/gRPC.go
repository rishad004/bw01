package di

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rishad004/bw01/apiGateway/pkg/handler"
	"github.com/rishad004/bw01/apiGateway/pkg/routers"
	m01_pb "github.com/rishad004/bw01_proto-files/microservice-01"
	m02_pb "github.com/rishad004/bw01_proto-files/microservice-02"
	"google.golang.org/grpc"
)

func Config(r *mux.Router) {
	ConnM01, m01Svc := Micro01Conn()
	defer ConnM01.Close()

	ConnM02, m02Svc := Micro02Conn()
	defer ConnM02.Close()

	hanlder := handler.NewHandler(m01Svc,m02Svc)

	routers.Routers(r, hanlder)

	log.Println("ApiGateway listening on port :8080")
	http.ListenAndServe("localhost:8080", r)
}

func Micro01Conn() (*grpc.ClientConn, m01_pb.Micro01Client) {
	connM01, err := grpc.Dial("bw01-microservice-01-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to microservice01 service:", err)
	}

	m01Svc := m01_pb.NewMicro01Client(connM01)

	return connM01, m01Svc
}

func Micro02Conn() (*grpc.ClientConn, m02_pb.Micro02Client) {
	connM02, err := grpc.Dial("bw01-microservice-02-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to microservice02 service:", err)
	}

	m02Svc := m02_pb.NewMicro02Client(connM02)

	return connM02, m02Svc
}
