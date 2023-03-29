CREATE TABLE IF NOT EXISTS income(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id integer not NULL,
  income_type integer not NULL,
  amount integer NOT NULL,
  date datetime,
  FOREIGN key(user_id) REFERENCES users(id),
  FOREIGN key(income_type) REFERENCES incometype(id) 
);