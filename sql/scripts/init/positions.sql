CREATE TABLE IF NOT EXISTS positions
(
    id INT NOT NULL AUTO_INCREMENT,
    amount FLOAT,
    description text,
    price FLOAT NOT NULL,
    discount FLOAT,
    bill INT NOT NULL,
    type INT NOT NULL DEFAULT 0,
    CONSTRAINT position_pkey PRIMARY KEY (id)
);