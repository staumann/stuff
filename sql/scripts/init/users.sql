CREATE TABLE IF NOT EXISTS users
(
    id INT NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    image VARCHAR(100),
    password VARCHAR(100),

    CONSTRAINT user_pkey PRIMARY KEY (id)
);