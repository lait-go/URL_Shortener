-- Миграция: откат для url_clicks (удаление индексов и таблицы)
DROP INDEX IF EXISTS idx_url_clicks_clicked_at;
DROP INDEX IF EXISTS idx_url_clicks_user_agent;
DROP TABLE IF EXISTS url_clicks;

-- Миграция: откат для shortened_urls (удаление таблицы)
DROP TABLE IF EXISTS shortened_urls;
