package repository

type Parse interface {
}

type ParseSave interface {
}

type Repository struct {
	Parse
	ParseSave
}

func NewRepository() *Repository {
	return &Repository{}
}
