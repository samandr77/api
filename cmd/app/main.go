package main

import (
	"github.com/gorilla/mux"
	"log"
	"nefstr/internal/database"
	"nefstr/internal/handlers"
	"nefstr/internal/messagesService"
	"net/http"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/messages/{id}", handler.DeleteMessageHandler).Methods("DELETE")
	router.HandleFunc("/api/messages/{id}", handler.PatchMessageHandler).Methods("PATCH")

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
