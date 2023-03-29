CREATE TABLE IF NOT EXISTS sessions(
  token primary key,
  userid INTEGER NOT NULL,
  FOREIGN key(userid) REFERENCES users(id) 
);