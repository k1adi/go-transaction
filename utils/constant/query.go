package constant

const (
	BANK_INSERT = "INSERT INTO mst_bank(id,name) VALUES($1,$2)"
	BANK_LIST   = "SELECT id,name FROM mst_bank"
	BANK_DETAIL = "SELECT id,name FROM mst_bank WHERE id=$1"
	BANK_UPDATE = "UPDATE mst_bank SET name=$1 WHERE id=$2"

	MERCHANT_INSERT = "INSERT INTO mst_merchant(id,name,address) VALUES($1,$2,$3)"
	MERCHANT_LIST   = "SELECT id,name,address FROM mst_merchant"
	MERCHANT_DETAIL = "SELECT id,name,address FROM mst_merchant WHERE id=$1"
	MERCHANT_UPDATE = "UPDATE mst_merchant SET name=$1, address=$2 WHERE id=$3"

	CUSTOMER_INSERT     = "INSERT INTO mst_customer(id,fullname,username,password,phone_number,created_at) VALUES($1,$2,$3,$4,$5,$6)"
	CUSTOMER_LIST       = "SELECT id,fullname,username,phone_number,created_at FROM mst_customer"
	CUSTOMER_DETAIL     = "SELECT id,fullname,username,phone_number,created_at FROM mst_customer WHERE id=$1"
	CUSTOMER_UPDATE     = "UPDATE mst_merchant SET fullname=$1, username=$2, phone_number=$3 WHERE id=$4"
	CUSTOMER_VALIDATION = "SELECT id,username,password FROM mst_customer WHERE username=$1"

	ADMIN_INSERT     = "INSERT INTO mst_admin(id,username,password) VALUES($1,$2,$3)"
	ADMIN_LIST       = "SELECT id,username FROM mst_admin"
	ADMIN_VALIDATION = "SELECT id,username,password FROM mst_admin WHERE username=$1"

	TRANSACTION_INSERT      = "INSERT INTO tx_transaction(id,customer_id,merchant_id,bank_id,amount,transaction_at) VALUES($1,$2,$3,$4,$5,$6)"
	TRANSACTION_LIST        = "SELECT id,customer_id,merchant_id,bank_id,amount,transaction_at FROM tx_transaction"
	TRANSACTION_BY_CUSTOMER = "SELECT id,customer_id,merchant_id,bank_id,amount,transaction_at FROM tx_transaction WHERE customer_id=$1"
)
