package main

import (
	"DataProcessorService/app/database/postgres"
	"DataProcessorService/app/internal/config"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"time"
)

type Event struct {
	EventType string `json:"eventType"`
	UserID    int64  `json:"userID"`
	EventTime string `json:"eventTime"`
	Payload   string `json:"payload"`
}

func main() {
	cfg := config.InitConfig()

	// Инициализация подключения к базе данных
	dbConnect, err := postgres.NewConnection(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialized database: %s", err)
	}
	defer dbConnect.Close()

	// Чтение данных из файла
	file, err := ioutil.ReadFile("/Users/tomik/GolandProjects/DataProcessorService/app/testdata/testdata.json")
	if err != nil {
		log.Fatal(err)
	}

	var events []Event
	if err := json.Unmarshal(file, &events); err != nil {
		log.Fatal(err)
	}

	// Вставка данных в таблицу events
	for _, event := range events {
		_, err := dbConnect.Exec("INSERT INTO events (eventType, userID, eventTime, payload) VALUES ($1, $2, $3, $4)",
			event.EventType, event.UserID, event.EventTime, event.Payload)
		if err != nil {
			log.Println("Failed to insert event:", err)
		}
	}

	fmt.Println("Data inserted successfully.")

	eventType := "music"
	startTime := time.Date(2018, 06, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 10, 31, 23, 59, 59, 999, time.UTC)
	sortEvents(dbConnect, eventType, startTime, endTime)
}

func sortEvents(db *sql.DB, eventType string, startTime, endTime time.Time) {
	query := `SELECT eventType, userID, eventTime, payload FROM events WHERE eventType = $1 AND eventTime BETWEEN $2 AND $3`
	rows, err := db.Query(query, eventType, startTime, endTime)
	if err != nil {
		log.Fatalf("Failed to query events: %v", err)
	}
	defer rows.Close()

	var selectedEvents []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.EventType, &event.UserID, &event.EventTime, &event.Payload); err != nil {
			log.Printf("Failed to scan event: %v", err)
		}
		selectedEvents = append(selectedEvents, event)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error occurred while iterating over rows: %v", err)
	}

	// Вывод выбранных событий
	for _, event := range selectedEvents {
		fmt.Printf("Event: %+v\n", event)
	}
}
