CREATE TABLE IF NOT EXISTS spending(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id integer not NULL,
  spending_type integer not NULL,
  amount integer NOT NULL,
  date datetime,
  FOREIGN key(user_id) REFERENCES users(id),
  FOREIGN key(spending_type) REFERENCES spendingtype(id) 
);