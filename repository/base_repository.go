package repository

import "database/sql"

type BaseRepository[Model any] interface {
	Create(bodyRequest Model) error
	List() ([]Model, error)
	Scan(rows *sql.Rows) ([]Model, error)
	Detail(id string) (Model, error)
	Update(bodyRequest Model) error
}
