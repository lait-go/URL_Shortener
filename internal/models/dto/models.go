package dto

import "time"

type FullURL struct {
	URL string `json:"url" validate:"required"`
}

type Analytics struct {
	ShortURL  string    `json:"short_url" validate:"required"`
	Ip        string    `json:"ip" validate:"required"`
	UserAgent string    `json:"user_agent" validate:"required"`
	Time      time.Time `json:"time" validate:"required"`
}
