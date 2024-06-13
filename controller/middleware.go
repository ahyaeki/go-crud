package controller

import (
	"database/sql"
	"go-crud/repository"
	"net/http"
)

func AuthMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionToken := r.Header.Get("Authorization")
		if sessionToken == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		user, err := repository.GetUserBySessionToken(db, sessionToken)
		if err == sql.ErrNoRows {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		r.Header.Set("username", user.Username)
		next.ServeHTTP(w, r)
	})
}
