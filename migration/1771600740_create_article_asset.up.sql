CREATE TABLE IF NOT EXISTS article_assets (
  id CHAR(36) NOT NULL,
  article_id CHAR(36) NOT NULL,
  object_key TEXT NOT NULL,
  content_type VARCHAR(128) NOT NULL,
  content_size  BIGINT NOT NULL,
  status VARCHAR(32) NOT NULL,
  created_at BIGINT NOT NULL,
  updated_at BIGINT,
  deleted_at BIGINT,
  PRIMARY KEY (id),
  FOREIGN KEY (article_id) REFERENCES articles(id)
);
