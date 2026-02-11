CREATE TABLE IF NOT EXISTS account_permissions (
  id CHAR(36) NOT NULL,
  account_id CHAR(36) NOT NULL,
  permission_id CHAR(36) NOT NULL,
  created_at BIGINT NOT NULL,
  updated_at BIGINT,
  deleted_at BIGINT,
  PRIMARY KEY (id),
  FOREIGN KEY (account_id) REFERENCES accounts(id),
  FOREIGN KEY (permission_id) REFERENCES permissions(id)
);

