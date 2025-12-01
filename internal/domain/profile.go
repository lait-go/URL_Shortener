package domain

import (
	"time"
)

// Сущность короткой ссылки
type ShortURL struct {
	ID          string    `json:"id" db:"id"`                     // уникальный идентификатор
	OriginalURL string    `json:"original_url" db:"original_url"` // полный URL
	ShortCode   string    `json:"short_code" db:"short_code"`      // короткий код (abc123)
	CreatedAt   time.Time `json:"created_at" db:"created_at"`     // время создания
}

// Запрос на создание короткой ссылки
type ShortURLRequest struct {
	OriginalURL string `json:"original_url" binding:"required,url"`
	ShortCode   string `json:"short_code,omitempty"`
}

// Информация о переходе по короткой ссылке (для аналитики)
type ClickAnalytics struct {
	ShortCode string    `json:"short_code" db:"short_code"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	UserAgent string    `json:"user_agent" db:"user_agent"`
	IP        string    `json:"ip,omitempty" db:"ip"`
	// Можно добавить поля для устройства, браузера, страны и т.д.
}

// Ответ с аналитикой по короткой ссылке
type AnalyticsResponse struct {
	ShortCode     string         `json:"short_code"`
	TotalClicks   int            `json:"total_clicks"`
	ClicksByDay   map[string]int `json:"clicks_by_day"`
	ClicksByAgent map[string]int `json:"clicks_by_user_agent"`
	// Дополнительная агрегация при необходимости
}
