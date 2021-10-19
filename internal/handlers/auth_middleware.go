package handlers

import (
	"context"
	"github.com/igor-koniukhov/fastcat/services"
	web "github.com/igor-koniukhov/webLogger/v2"
	"net/http"
)

func (us *UserHandler) AuthCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := us.App.TemplateInfo["Authorization"]
		if auth == "" {
			http.Redirect(w, r, "/show-login", http.StatusSeeOther)
			return
		}
		if auth != "" {
			tokenString, err := services.GetTokenFromBearerString(auth)
			if err != nil {
				web.Log.Error(err)
				return
			}
			creds, err := services.ValidateToken(tokenString, services.AccessSecret)
			if err != nil {
				web.Log.Error(err, "expired")
				us.App.TemplateInfo["Expired"]="token expired"
				http.Redirect(w, r, "/show-login", http.StatusSeeOther)
				return
			}
			us.App.TemplateInfo["Expired"]=""

			ctx := context.WithValue(r.Context(), "user_id", creds.ID)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	}
}

func (us *UserHandler) AuthSet(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := us.App.TemplateInfo["Authorization"]
		w.Header().Set("Authorization", auth)
		next.ServeHTTP(w, r)

	}
}
