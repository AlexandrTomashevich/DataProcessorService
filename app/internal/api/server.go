package api

import (
	"DataProcessorService/app/internal/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	DBConnect *sql.DB
}

func NewServer(DBConnect *sql.DB) *Server {
	return &Server{DBConnect: DBConnect}
}

func (s *Server) Event(w http.ResponseWriter, r *http.Request) {
	var event models.EventRequest

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(event)

	query := `INSERT INTO events (eventType, userID, eventTime, payload)
              VALUES ($1, $2, $3, $4)`

	_, err := s.DBConnect.Exec(query, event.EventType, event.UserID, event.EventTime, event.Payload)
	if err != nil {
		log.Printf("Failed to insert event: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{})
}
