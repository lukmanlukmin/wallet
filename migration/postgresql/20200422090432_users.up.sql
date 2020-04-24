CREATE TYPE type_user AS ENUM('admin', 'guest');
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username varchar(100) DEFAULT NULL,
  email varchar(50) NOT NULL UNIQUE,
  password varchar(100) NOT NULL,
  user_type type_user NOT NULL,
  available BOOLEAN NOT NULL
);
ALTER TABLE users ALTER COLUMN available SET DEFAULT FALSE;