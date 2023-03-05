
CREATE DATABASE IF NOT EXISTS social;


CREATE TABLE IF NOT EXISTS social.users (
    userid int unsigned NOT NULL AUTO_INCREMENT,
    username varchar(255) NOT NULL DEFAULT "",
    password varchar(255) NOT NULL ,
    age      int
    email    varchar(255) UNIQUE NOT NULL,
    gender   varchar(255) NOT NULL,
    marital  VARCHAR(255)
    photos   text NOT NULL DEFAULT "",
    number_photos int NOT NULL DEFAULT 0,
    PRIMARY KEY (userid)
);


-- UPDATE EXAMPLE
-- UPDATE  users set gender = 'f' where userid = 1 ;
-- UPDATE table_name  SET  col3 = 'some_value', col4 = 'some_other_value WHERE  id = 3 ;
-- ALTER TABLE users ADD descript text default "" AFTER contry;
