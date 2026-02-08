CREATE TABLE IF NOT EXISTS refresh_tokens (
  id CHAR(36) NOT NULL,
  account_id CHAR(36) NOT NULL,
  token TEXT NOT NULL,
  expires_at BIGINT NOT NULL,
  created_at BIGINT NOT NULL,
  updated_at BIGINT,
  deleted_at BIGINT,
  PRIMARY KEY (id),
  FOREIGN KEY (account_id) REFERENCES accounts(id)
);
