CREATE TABLE IF NOT EXISTS account_identities (
  id CHAR(36) NOT NULL,
  account_id CHAR(36) NOT NULL,
  provider VARCHAR(32) NOT NULL,
  provider_user_id VARCHAR(128) NOT NULL,
  created_at BIGINT NOT NULL,
  updated_at BIGINT,
  deleted_at BIGINT,
  PRIMARY KEY (id),
  UNIQUE (provider, provider_user_id),
  FOREIGN KEY (account_id) REFERENCES accounts(id)
);
