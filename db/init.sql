CREATE TABLE IF NOT EXISTS bots (
  id SERIAL PRIMARY KEY,
  token TEXT NOT NULL UNIQUE,
  confession TEXT,
  chat_id TEXT,
  channel_id TEXT,
  type TEXT,
  code TEXT
);
