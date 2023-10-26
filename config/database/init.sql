CREATE DATABASE go_transaction;

CREATE TABLE mst_merchant (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    address TEXT NOT NULL
);

CREATE TABLE mst_bank (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE mst_customer (
    id VARCHAR(100) PRIMARY KEY,
    fullname VARCHAR(150) NOT NULL,
    username VARCHAR(20) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    phone_number VARCHAR(15) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,

);

CREATE TABLE tx_transaction (
    id VARCHAR(100) PRIMARY KEY,
    bank_id VARCHAR(100) NOT NULL,
    customer_id VARCHAR(100) NOT NULL,
    merchant_id VARCHAR(100) NOT NULL,
    amount BIGINT NOT NULL,
    transaction_at TIMESTAMP NOT NULL,

    CONSTRAINT fk_bank_id FOREIGN KEY(bank_id) REFERENCES mst_bank(id),
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES user_credential(id),
    CONSTRAINT fk_merchant_id FOREIGN KEY(merchant_id) REFERENCES mst_merchant(id)
);

CREATE TABLE mst_admin (
    id VARCHAR(100) PRIMARY KEY,
    username VARCHAR(30) NOT NULL,
    password VARCHAR(200) NOT NULL
);