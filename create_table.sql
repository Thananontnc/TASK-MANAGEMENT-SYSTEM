CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    phone VARCHAR(20));

CREATE TABLE tasks(
    ID int AUTO_INCREMENT PRIMARY KEY,
    tasks TEXT NOT NULL
);