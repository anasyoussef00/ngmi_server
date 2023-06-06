package authentication_resource

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ngmi_server/internal/models/user_model"
	authenticationService "ngmi_server/internal/services/authentication_service"
	"ngmi_server/pkg/log"
)

type Resource struct {
	Service authenticationService.Service
	Logger  log.Logger
}

func (res Resource) Register(w http.ResponseWriter, r *http.Request) {
	var userRegisterReq user_model.RegisterReq
	//u, err := res.service.Register()
	if err := json.NewDecoder(r.Body).Decode(&userRegisterReq); err != nil {
		http.Error(w, fmt.Sprintf("%s: %v", http.StatusText(http.StatusBadRequest), err), http.StatusBadRequest)
		return
	}

	if userRegisterRes, err := res.Service.Register(userRegisterReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(userRegisterRes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (res Resource) Login(w http.ResponseWriter, r *http.Request) {
	var userLoginReq user_model.LoginReq
	//u, err := res.service.Register()
	if err := json.NewDecoder(r.Body).Decode(&userLoginReq); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Write([]byte("WE OUT HERE"))
}
