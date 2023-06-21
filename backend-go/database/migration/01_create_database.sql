DROP TABLE IF EXISTS notebook;

CREATE TABLE IF NOT EXISTS notebook (
  id SERIAL PRIMARY KEY,
  username VARCHAR NOT NULL,
  source VARCHAR NOT NULL,
  source_id VARCHAR NOT NULL,
  tag text[],
  note text
);