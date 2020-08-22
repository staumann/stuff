CREATE TABLE IF NOT EXISTS shops
(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100),
    city VARCHAR(100),
    postalCode VARCHAR(50),
    street VARCHAR(100),
    houseNumber VARCHAR(50),
    CONSTRAINT shop_pkey PRIMARY KEY (id),
    INDEX idx_shop_name (name)
);