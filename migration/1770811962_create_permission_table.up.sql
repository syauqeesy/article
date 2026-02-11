CREATE TABLE IF NOT EXISTS permissions (
  id CHAR(36) NOT NULL,
  code VARCHAR(128) NOT NULL,
  name VARCHAR(128) NOT NULL,
  created_at BIGINT NOT NULL,
  updated_at BIGINT,
  deleted_at BIGINT,
  PRIMARY KEY (id)
);

