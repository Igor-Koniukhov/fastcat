package main

import (
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"github.com/igor-koniukhov/fastcat/services"
	web "github.com/igor-koniukhov/webLogger/v3"
	"net/http"
	"os"
)

func AuthMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookAuth, err := r.Cookie("Bearer")
		bearerString := cookAuth.String()
		web.Log.Error(err, err)
		tokenString, err := services.GetTokenFromBearerString(bearerString)
		web.Log.Error(err, err)
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

		auth := &http.Cookie{
			Name:  "Authorized",
			Value: user.Email,
		}
		http.SetCookie(w, auth)
		http.Redirect(w, r, "/users", http.StatusSeeOther)
		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
	}
}

