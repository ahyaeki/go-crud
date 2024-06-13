package main

import (
	"fmt"
	"log"
	"net/http"

	"go-crud/controller"
	"go-crud/database"

	"github.com/gorilla/sessions"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := sessions.NewCookieStore([]byte("something-very-secret"))

	http.HandleFunc("/login", controller.LoginHandler(db, store))
	http.HandleFunc("/logout", controller.LogoutHandler(db, store))
	http.Handle("/items", controller.AuthMiddleware(db, http.HandlerFunc(controller.ItemsHandler(db))))
	http.Handle("/items/", controller.AuthMiddleware(db, http.HandlerFunc(controller.ItemHandler(db))))

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
