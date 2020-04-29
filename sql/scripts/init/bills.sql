CREATE TABLE IF NOT EXISTS bills
(
    id INT NOT NULL AUTO_INCREMENT,
    userId INT NOT NULL,
    shopId INT NOT NULL,
    discount FLOAT,
    timeStamp TIMESTAMP,
    CONSTRAINT bill_pkey PRIMARY KEY (id)
);
