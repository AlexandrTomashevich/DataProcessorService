package models

type EventRequest struct {
	EventType string `json:"eventType"`
	UserID    int64  `json:"userID"`
	EventTime string `json:"eventTime"`
	Payload   string `json:"payload"`
}
