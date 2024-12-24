package main

import (
	"github.com/rishad004/bw01/microservice-01/pkg/config"
	"github.com/rishad004/bw01/microservice-01/pkg/di"
)

func main() {

	config.Config()

	di.InitgRPC()

}
