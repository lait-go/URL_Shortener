package db

import "time"

type Analytics struct {
	ID        int       `json:"id"`
	ShortURL  string    `json:"short_url" validate:"required"`
	Ip        string    `json:"ip" validate:"required"`
	UserAgent string    `json:"user_agent" validate:"required"`
	Time      time.Time `json:"time" validate:"required"`
}
