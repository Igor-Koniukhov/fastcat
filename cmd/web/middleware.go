package main

import (
	"context"
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"github.com/igor-koniukhov/fastcat/services"
	web "github.com/igor-koniukhov/webLogger/v3"
	"net/http"
	"os"
)


func  AuthMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		bearerString := r.Header.Get("Authorization")
		tokenString, err := services.GetTokenFromBearerString(bearerString)
		web.Log.Error(err)
		claims, err := services.ValidateToken(tokenString, os.Getenv("AccessSecret"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		user, err := repository.Repo.GetUserByID(claims.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := &models.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		web.Log.Error(err, "message : ", err)
		r = r.WithContext(context.WithValue(r.Context(), "user_id", resp.ID))

		next.ServeHTTP(w, r)
	}
}








