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
)
