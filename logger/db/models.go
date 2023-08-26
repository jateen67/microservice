package db

import "time"

type LogEntry struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}
