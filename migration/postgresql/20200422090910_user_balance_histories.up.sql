CREATE TYPE type_transaction AS ENUM('credit', 'debit');
CREATE TABLE IF NOT EXISTS user_balance_histories (
  id SERIAL PRIMARY KEY,
  user_balance_id integer NOT NULL,
  balance_before integer NOT NULL,
  balance_after integer NOT NULL,
  activity varchar(50) NOT NULL,
  type_activity type_transaction NOT NULL,
  ip varchar(20) NOT NULL,
  user_agent varchar(100) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);