DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

INSERT INTO users (id, email, password) 
VALUES 
('1', 'admin@secretaria.com', 'senha123'),
('2', 'professor@escola.com', 'abcd9876');