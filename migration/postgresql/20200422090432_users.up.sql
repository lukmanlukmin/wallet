CREATE TABLE IF NOT EXISTS users (
  id integer NOT NULL,
  username varchar(100) DEFAULT NULL,
  email varchar(50) NOT NULL UNIQUE,
  password varchar(100) NOT NULL UNIQUE,
  PRIMARY KEY (id)
);