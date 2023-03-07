
CREATE DATABASE IF NOT EXISTS mersal;

-- sqlite
CREATE TABLE IF NOT EXISTS users (
     id INTEGER PRIMARY KEY AUTOINCREMENT,
     username varchar(64) NOT NULL DEFAULT "",
     email    varchar(128) UNIQUE NOT NULL,
     password varchar(64) NOT NULL ,
     photos   text NOT NULL DEFAULT ""
);

-- UPDATE EXAMPLE
-- UPDATE  users set gender = 'f' where userid = 1 ;
-- UPDATE table_name  SET  col3 = 'some_value', col4 = 'some_other_value WHERE  id = 3 ;
-- ALTER TABLE users ADD descript text default "" AFTER contry;
