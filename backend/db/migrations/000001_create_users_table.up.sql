CREATE TABLE users(
  	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
	email VARCHAR(255) NOT NULL unique,
	password CHAR(60) NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions(
  token primary key,
  userid INTEGER NOT NULL,
  FOREIGN key(userid) REFERENCES users(id) 
);

CREATE TABLE IF NOT EXISTS spending(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id integer not NULL,
  spending_type integer not NULL,
  amount integer NOT NULL,
  date datetime,
  FOREIGN key(user_id) REFERENCES users(id),
  FOREIGN key(spending_type) REFERENCES spendingtype(id) 
);

CREATE TABLE IF NOT EXISTS spendingtype(
 	id INTEGER PRIMARY KEY AUTOINCREMENT,
 	spendingtype VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS income(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id integer not NULL,
  income_type integer not NULL,
  amount integer NOT NULL,
  date datetime,
  FOREIGN key(user_id) REFERENCES users(id),
  FOREIGN key(income_type) REFERENCES incometype(id) 
);

CREATE TABLE incometype(
 	id INTEGER PRIMARY KEY AUTOINCREMENT,
 	incometype VARCHAR(100) NOT NULL
);