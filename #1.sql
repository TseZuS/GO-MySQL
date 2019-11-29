CREATE DATABASE if not exists sbase collate utf8_general_ci;
USE sbase;
CREATE TABLE IF NOT EXISTS book (
	id INT (11) NOT NULL AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL UNIQUE,
    autor VARCHAR(20) NOT NULL,
    date varchar(11) NOT NULL,
    PRIMARY KEY (id)
);



SELECT * FROM sbase.book;
