DROP TABLE IF EXISTS notebook;
DROP TABLE IF EXISTS notebooks;

CREATE TABLE IF NOT EXISTS notebooks (
  id SERIAL PRIMARY KEY,
  username VARCHAR NOT NULL,
  source VARCHAR NOT NULL,
  source_id VARCHAR NOT NULL,
  tag text[],
  note text
);