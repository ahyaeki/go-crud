package controller

import (
	"database/sql"
	"encoding/json"
	"go-crud/model"
	"go-crud/repository"
	"net/http"

	"github.com/gorilla/sessions"

	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func LoginHandler(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		login(w, r, db, store)
	}
}

func login(w http.ResponseWriter, r *http.Request, db *sql.DB, store *sessions.CookieStore) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbUser, err := repository.GetUserByUsername(db, user.Username)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.Password != dbUser.Password {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	sessionToken, err := generateULID()
	if err != nil {
		http.Error(w, "Failed to generate session token", http.StatusInternalServerError)
		return
	}

	err = repository.UpdateUserSessionToken(db, dbUser.ID, sessionToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, _ := store.Get(r, "session-name")
	session.Values["session_token"] = sessionToken
	session.Save(r, w)

	dbUser.SessionToken = sessionToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbUser)
}

func LogoutHandler(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken := r.Header.Get("Authorization")
		if sessionToken == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		err := repository.ClearUserSessionToken(db, sessionToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session, _ := store.Get(r, "session-name")
		session.Values["session_token"] = nil
		session.Save(r, w)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logged out successfully"))
	}
}

func generateULID() (string, error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String(), nil
}
