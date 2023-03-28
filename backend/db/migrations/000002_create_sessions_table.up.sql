CREATE TABLE sessions(
  token primary key,
  userid INTEGER NOT NULL,
  FOREIGN key(userid) REFERENCES users(id) 
);