CREATE TABLE users(
  	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
	email VARCHAR(255) NOT NULL unique,
	password CHAR(60) NOT NULL
);


CREATE TABLE IF NOT EXISTS spending(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id integer not NULL,
  spending_type integer not NULL,
  amount integer NOT NULL,
  description TEXT NOT NULL,
  date DATE,
  FOREIGN key(user_id) REFERENCES users(id),
  FOREIGN key(spending_type) REFERENCES spendingtype(id) 
);

CREATE TABLE IF NOT EXISTS spendingtype(
 	id INTEGER PRIMARY KEY AUTOINCREMENT,
 	spending_type VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS income(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id integer not NULL,
  income_type integer not NULL,
  amount integer NOT NULL,
  description TEXT NOT NULL,
  date DATE,
  FOREIGN key(user_id) REFERENCES users(id),
  FOREIGN key(income_type) REFERENCES incometype(id) 
);

CREATE TABLE incometype(
 	id INTEGER PRIMARY KEY AUTOINCREMENT,
 	income_type VARCHAR(100) NOT NULL
);

CREATE TABLE loandebttype(
 	id INTEGER PRIMARY KEY AUTOINCREMENT,
 	type VARCHAR(100) NOT NULL
);

CREATE TABLE images(
 	user_id INTEGER unique,
 	image_name TEXT NOT NULL,
  FOREIGN key(user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS loan(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id integer not NULL,
  type integer not NULL,
  amount integer NOT NULL,
  description TEXT NOT NULL,
  date DATE,
  FOREIGN key(user_id) REFERENCES users(id),
  FOREIGN key(type) REFERENCES loandebttype(id) 
);

CREATE TABLE IF NOT EXISTS debt(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id integer not NULL,
  type integer not NULL,
  amount integer NOT NULL,
  description TEXT NOT NULL,
  date DATE,
  FOREIGN key(user_id) REFERENCES users(id),
  FOREIGN key(type) REFERENCES loandebttype(id) 
);


INSERT INTO spendingtype (spending_type) VALUES ("Денсаулық"),("Ойын-сауық"),("Транспорт"),("Сыйлықтар"),("Тамақ"),("Спорт"),("Шоппинг"),("Несие"),("Байланыс"),("Ком.Қызметтер"),("Басқа");

INSERT INTO incometype (income_type) VALUES ("Айлық"),("Сыйақы"),("Қосымша табыс");

INSERT INTO loandebttype (type) VALUES ("Достар"),("Туысқандар"),("Коллегалар"),("Басқа адамдар");

INSERT INTO debt (user_id, type, description, amount, date)
VALUES (3, 1, 'Қарыз үшін', 500000, '2023-01-01'),
       (3, 2, 'Қарыз беру', 300000, '2023-02-05'),
       (3, 3, 'Кредитті төлеу', 700000, '2023-03-10'),
       (3, 4, 'Машина үшін', 1000000, '2023-03-15'),
       (3, 1, 'Қарыз алу', 200000, '2023-04-20');