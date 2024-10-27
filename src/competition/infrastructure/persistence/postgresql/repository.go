package postgresql

import (
	"scibat/src/common/infrastructure/persistence/postgresql"
)

type Repository struct {
	postgresql *postgresql.Service
}

func (r *Repository) New(postgresql *postgresql.Service) *Repository {
	return &Repository{
		postgresql: postgresql,
	}
}

func (r *Repository) FindByUserId() {
	//
}
