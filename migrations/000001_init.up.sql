CREATE SCHEMA service_product

CREATE TABLE service_product.product (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description varchar(1024) NOT NULL,
    price BIGINT NOT NULL,
    status INTEGER NOT NULL
);