# Go Transaction

## Overview

This simple API is created to facilitate new transactions between customers and merchants. There are two types of access: customer and admin.

An admin can manage merchant and bank data, as well as view customer data and a list of all transactions. On the other hand, a customer can only add new transaction data and view transaction history based on the customer currently logged in.

A transaction consists of customer data, merchant data, bank data, and the transaction amount.

## How To Use

- Clone this repository to your local computer.
- Copy `.env.example` file to `.env` file.
- Fill in each key according to your configuration.
- Run `go get` to download and install the required packages or dependencies.
- Open `config/database/init.sql` file to create a database and tables that match those used in the project.
- Additionally, you can add the provided dummy data from the `config/database/dummy.sql` file.
- Create a `log.txt` file to store every request log from the methods being used.
- Run `go run main.go` or `go run .` to start the API.

> "For the API documentation, you can import `GO - Transaction.postman_collection.json` file into the Postman application that is already installed on your computer."
