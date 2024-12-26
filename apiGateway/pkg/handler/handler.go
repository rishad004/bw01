package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/rishad004/bw01/apiGateway/pkg/domain"
	"github.com/rishad004/bw01/apiGateway/pkg/utils"
	m01_pb "github.com/rishad004/bw01_proto-files/microservice-01"
	m02_pb "github.com/rishad004/bw01_proto-files/microservice-02"
)

type Handler struct {
	M01 m01_pb.Micro01Client
	M02 m02_pb.Micro02Client
}

var mutex sync.Mutex

func NewHandler(m01 m01_pb.Micro01Client, m02 m02_pb.Micro02Client) *Handler {
	return &Handler{M01: m01, M02: m02}
}

func (h *Handler) Method(w http.ResponseWriter, r *http.Request) {
	log.Println("===========Method===========")

	var request domain.Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.SendJSONResponse(w, domain.H{
			"message": "Invalid payload!",
			"error":   err.Error(),
		}, http.StatusBadRequest, r)
		return
	}

	if request.Method == "1" {
		mutex.Lock()
		res, err := h.M02.Method01(context.Background(), &m02_pb.Data{})

		time.Sleep(time.Duration(request.WaitTime) * time.Second)
		mutex.Unlock()

		if err != nil {
			utils.SendJSONResponse(w, err.Error(), http.StatusInternalServerError, r)
			return
		}

		utils.SendJSONResponse(w, res, http.StatusOK, r)
		return
	} else if request.Method == "2" {
		res, err := h.M02.Method02(context.Background(), &m02_pb.Data{})

		time.Sleep(time.Duration(request.WaitTime) * time.Second)
		if err != nil {
			utils.SendJSONResponse(w, err.Error(), http.StatusInternalServerError, r)
			return
		}

		utils.SendJSONResponse(w, res, http.StatusOK, r)
		return
	} else {
		utils.SendJSONResponse(w, "No method found!", http.StatusNotFound, r)
		return
	}
}

func (h *Handler) UserCreate(w http.ResponseWriter, r *http.Request) {
	log.Println("===========UserCreate===========")

	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJSONResponse(w, domain.H{
			"message": "Invalid payload!",
			"error":   err.Error(),
		}, http.StatusBadRequest, r)
		return
	}

	res, err := h.M01.UserCreate(context.Background(), &m01_pb.Details{Name: user.Name, Email: user.Email})
	if err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusInternalServerError, r)
		return
	}

	utils.SendJSONResponse(w, domain.H{
		"message": "User created successfully!",
		"userId":  res.Id,
	}, http.StatusOK, r)

}

func (h *Handler) UserFetch(w http.ResponseWriter, r *http.Request) {
	log.Println("===========UserFetch===========")

	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJSONResponse(w, domain.H{
			"message": "Invalid payload!",
			"error":   err.Error(),
		}, http.StatusBadRequest, r)
		return
	}

	res, err := h.M01.UserFetch(context.Background(), &m01_pb.Get{Id: int32(user.Id)})
	if err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusInternalServerError, r)
		return
	}

	utils.SendJSONResponse(w, res, http.StatusOK, r)
}

func (h *Handler) UserUpdate(w http.ResponseWriter, r *http.Request) {
	log.Println("===========UserUpdate===========")

	var user domain.User

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		utils.SendJSONResponse(w, domain.H{
			"message": "Invalid payload!",
			"error":   err.Error(),
		}, http.StatusBadRequest, r)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJSONResponse(w, domain.H{
			"message": "Invalid payload!",
			"error":   err.Error(),
		}, http.StatusBadRequest, r)
		return
	}

	if _, err := h.M01.UserUpdate(context.Background(), &m01_pb.Details{Id: int32(id), Name: user.Name, Email: user.Email}); err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusInternalServerError, r)
		return
	}

	utils.SendJSONResponse(w, "Updated user!", http.StatusOK, r)
}

func (h *Handler) UserDelete(w http.ResponseWriter, r *http.Request) {
	log.Println("===========UserDelete===========")

	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJSONResponse(w, domain.H{
			"message": "Invalid payload!",
			"error":   err.Error(),
		}, http.StatusBadRequest, r)
		return
	}

	if _, err := h.M01.UserDelete(context.Background(), &m01_pb.Get{Id: int32(user.Id)}); err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusInternalServerError, r)
		return
	}

	utils.SendJSONResponse(w, "Deleted user successfully!", http.StatusOK, r)
}
