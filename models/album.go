package models

import "time"

type Album struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
