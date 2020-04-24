CREATE TABLE IF NOT EXISTS user_balances (
  id SERIAL PRIMARY KEY,
  user_id integer NOT NULL,
  balance integer NOT NULL,
  balance_achieve integer NOT NULL
);