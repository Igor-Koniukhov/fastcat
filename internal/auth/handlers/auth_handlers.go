package handlers

import (
	"encoding/json"
	"gopath/auth/model"
	"gopath/auth/repository"
	"gopath/auth/services"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		logReq := new(model.LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&logReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp , err := services.TokenResponder(w, logReq)
		if err !=nil {
			log.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Println(err)
		}

	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		logReq := new(model.LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&logReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp , err := services.TokenResponder(w, logReq)
		if err !=nil {
			log.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Println(err)
		}

	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func GetProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		bearerString := r.Header.Get("Authorization")
		tokenString := services.GetTokenFromBearerString(bearerString)
		claims, err := services.ValidateToken(tokenString, services.AccessSecret)
		//claims, err := ValidateToken(GetTokenFromBearerString(r.Header.Get("Authorization")), RefreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}


		user, err := repository.NewUserRepository().GetUserByID(claims.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := &services.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Println(err)
		}

	default:
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)

	}
}

