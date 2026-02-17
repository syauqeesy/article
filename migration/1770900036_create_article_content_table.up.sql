CREATE TABLE IF NOT EXISTS article_contents (
  id CHAR(36) NOT NULL,
  article_id CHAR(36) NOT NULL,
  language CHAR(2) NOT NULL,
  title VARCHAR(256) NOT NULL,
  slug VARCHAR(256) NOT NULL,
  summary varchar(4096) NOT NULL,
  content TEXT NOT NULL,
  created_at BIGINT NOT NULL,
  updated_at BIGINT,
  deleted_at BIGINT,
  PRIMARY KEY (id),
  UNIQUE (slug),
  FOREIGN KEY (article_id) REFERENCES articles(id)
);

