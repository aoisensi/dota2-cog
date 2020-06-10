CREATE TABLE users(
  discord_id BIGINT PRIMARY KEY,
  steam_id BIGINT NOT NULL
);
CREATE INDEX ON users(steam_id);