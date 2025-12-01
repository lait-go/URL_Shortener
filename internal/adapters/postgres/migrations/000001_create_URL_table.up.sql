-- Миграция: создание таблицы shortened_urls
CREATE TABLE shortened_urls (
    id SERIAL PRIMARY KEY,
    short_key VARCHAR(100) UNIQUE NOT NULL, -- уникальный ключ сокращенной ссылки
    original_url TEXT NOT NULL,              -- оригинальный длинный URL
    -- custom_name BOOLEAN DEFAULT FALSE,       -- признак пользовательского имени, если нужно
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Миграция: создание таблицы для аналитики переходов
CREATE TABLE url_clicks (
    id SERIAL PRIMARY KEY,
    short_url_id INTEGER NOT NULL REFERENCES shortened_urls(id) ON DELETE CASCADE, -- ссылка на короткую ссылку
    user_agent TEXT NOT NULL,           -- User-Agent пользователя
    ip_address INET,                   -- IP-адрес пользователя
    clicked_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() -- время перехода
);

-- Индексы для аналитики по дате и User-Agent
CREATE INDEX idx_url_clicks_clicked_at ON url_clicks (clicked_at);
CREATE INDEX idx_url_clicks_user_agent ON url_clicks (user_agent);
