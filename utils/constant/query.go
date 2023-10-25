package constant

const (
	BANK_INSERT = "INSERT INTO mst_bank(id,name) VALUES($1,$2)"
	BANK_LIST   = "SELECT id,name FROM mst_bank"
	BANK_DETAIL = "SELECT id,name FROM mst_bank WHERE id=$1"
	BANK_UPDATE = "UPDATE mst_bank SET name=$1 WHERE id=$2"
)
