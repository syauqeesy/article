CREATE TABLE IF NOT EXISTS accounts (
  id         CHAR(36)      NOT NULL,
  email      VARCHAR(128)  NOT NULL,
  created_at BIGINT        NOT NULL,
  updated_at BIGINT,
  deleted_at BIGINT,
  PRIMARY KEY (id)
);
