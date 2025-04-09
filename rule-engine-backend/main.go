package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"rule-engine-backend/rules"
	"rule-engine-backend/db"
)

func main() {
	db.ConnectDB()
	r := mux.NewRouter()
	r.HandleFunc("/rules", rules.CreateRule).Methods("POST")
	r.HandleFunc("/rules", rules.GetRules).Methods("GET")
	r.HandleFunc("/event", rules.HandleEvent).Methods("POST")

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Create handler with CORS middleware
	handler := c.Handler(r)

	log.Println("🚀 Server started on http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
