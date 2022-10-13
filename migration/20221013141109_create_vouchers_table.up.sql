CREATE TABLE vouchers (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    brand_id VARCHAR(36) NULL,
    name VARCHAR(150) NOT NULL,
    cost_in_point INT NOT NULL,
    stock INT NULL NOT NULL,
    expiration_date DATE NOT NULL,
    CONSTRAINT vouchers_brand_id_on_brands FOREIGN KEY (brand_id) REFERENCES brands (id)
    ON UPDATE CASCADE ON DELETE NO ACTION
);