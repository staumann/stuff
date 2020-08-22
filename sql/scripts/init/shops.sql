CREATE TABLE IF NOT EXISTS shops
(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100),
    city VARCHAR(100),
    postal_code VARCHAR(50),
    street VARCHAR(100),
    house_number VARCHAR(50),
    infos TEXT,
    CONSTRAINT shop_pkey PRIMARY KEY (id),
    INDEX idx_shop_name (name)
);