package main

import (
	"github.com/gorilla/mux"
	"github.com/rishad004/bw01/apiGateway/pkg/di"
)

func main() {
	r := mux.NewRouter()

	di.Config(r)

}
