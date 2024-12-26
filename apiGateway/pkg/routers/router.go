package routers

import (
	"github.com/gorilla/mux"
	"github.com/rishad004/bw01/apiGateway/pkg/handler"
)

func Routers(r *mux.Router, h *handler.Handler) {

	// microservice-01
	r.HandleFunc("/user/create", h.UserCreate).Methods("POST")
	r.HandleFunc("/user/fetch", h.UserFetch).Methods("GET")
	r.HandleFunc("/user/update", h.UserUpdate).Methods("PUT")
	r.HandleFunc("/user/delete", h.UserDelete).Methods("DELETE")

	// microservice-02
	r.HandleFunc("/method", h.Method).Methods("POST")

}
