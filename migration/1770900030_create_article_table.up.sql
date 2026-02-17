CREATE TABLE IF NOT EXISTS articles (
  id CHAR(36) NOT NULL,
  account_id CHAR(36) NOT NULL,
  status VARCHAR(32) NOT NULL,
  views INT NOT NULL,
  created_at BIGINT NOT NULL,
  updated_at BIGINT,
  deleted_at BIGINT,
  PRIMARY KEY (id),
  FOREIGN KEY (account_id) REFERENCES accounts(id)
);

