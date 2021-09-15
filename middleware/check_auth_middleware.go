package middleware

import (
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("Authorized")
		if err !=nil {
			http.Redirect(w, r, "/show-login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
}



