CREATE TABLE IF NOT EXISTS user_balances (
  id integer NOT NULL,
  user_id integer NOT NULL,
  balance integer NOT NULL,
  balance_achieve integer NOT NULL,
  PRIMARY KEY (id)
);