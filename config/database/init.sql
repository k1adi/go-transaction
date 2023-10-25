CREATE DATABASE go_transaction;

CREATE TABLE mst_customer (
    id VARCHAR(100) PRIMARY KEY,
    fullname VARCHAR(100) NOT NULL,
    username VARCHAR(20) UNIQUE NOT NULL,
    phone_number VARCHAR(15) UNIQUE NOT NULL
);

CREATE TABLE mst_merchant (
    id VARCHAR(100) PRIMARY KEY,
    merchant_name VARCHAR(100) UNIQUE NOT NULL,
    address TEXT NOT NULL
);

CREATE TABLE mst_bank (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE user_credential (
    id VARCHAR(100) PRIMARY KEY,
    customer_id VARCHAR(100) NOT NULL,
    password TEXT NOT NULL,
    role VARCHAR(15) NOT NULL,

    CONSTRAINT fk_customer_id FOREIGN KEY(customer_id) REFERENCES mst_customer(id)
);

CREATE TABLE tx_transaction (
    id VARCHAR(100) PRIMARY KEY,
    bank_id VARCHAR(100) NOT NULL,
    user_id VARCHAR(100) NOT NULL,
    merchant_id VARCHAR(100) NOT NULL,
    amount BIGINT NOT NULL,

    CONSTRAINT fk_bank_id FOREIGN KEY(bank_id) REFERENCES mst_bank(id),
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES user_credential(id),
    CONSTRAINT fk_merchant_id FOREIGN KEY(merchant_id) REFERENCES mst_merchant(id)
);