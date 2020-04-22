CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username varchar(100) DEFAULT NULL,
  email varchar(50) NOT NULL UNIQUE,
  password varchar(100) NOT NULL UNIQUE
);