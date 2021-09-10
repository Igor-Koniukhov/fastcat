package handlers

import (
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/services"
	web "github.com/igor-koniukhov/webLogger/v3"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	logReq := new(model.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&logReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := services.TokenResponder(w, logReq)
	web.Log.Error(err, "message tokenResponder: ", err)

	err = json.NewEncoder(w).Encode(resp)
	http.SetCookie(w, &http.Cookie{
		Name:     "Authorisation",
		Value:    resp.AccessToken,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	},

	)

	w.WriteHeader(http.StatusOK)
	web.Log.Error(err, "message: ", err)

}

func Refresh(w http.ResponseWriter, r *http.Request) {
	logReq := new(model.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&logReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := services.TokenResponder(w, logReq)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
	}
}

