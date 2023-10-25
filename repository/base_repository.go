package repository

type BaseRepository[Model any] interface {
	Create(bodyRequest Model) error
	List() ([]Model, error)
	Detail(id string) (Model, error)
	Update(bodyRequest Model) error
}
