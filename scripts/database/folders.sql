CREATE TABLE folders (
  id SERIAL,
  parent_id INT,
  name VARCHAR(60) NOT NULL,
  created_at TIMESTAMP DEFAULT current_timestamp,
  modified_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
  PRIMARY KEY(id)
  CONSTRAINT fk_parent
    FOREIGN KEY(parent_id)
    REFERENCES folders(id)
)