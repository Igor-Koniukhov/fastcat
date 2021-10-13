package middleware

import (
	"context"
	"github.com/igor-koniukhov/fastcat/services"
	web "github.com/igor-koniukhov/webLogger/v2"
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth, err := r.Cookie("Authorization")
		if err!=nil {
			http.Redirect(w, r, "/show-login", http.StatusSeeOther)
			return
		}
		if auth.Value != "" {
			tokenString, err := services.GetTokenFromBearerString(auth.Value)
			if err != nil {
				web.Log.Error(err)
				return
			}
			_, err = services.ValidateToken(tokenString, services.AccessSecret)
			if err != nil {
				web.Log.Error(err, "expired")
				http.Redirect(w, r, "/refresh", http.StatusSeeOther)
				return
			}
			w.Header().Set("Authorization", auth.Value)
			ctx := context.WithValue(r.Context(), "Authorization", auth.Value)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	}
}




