package query

import "scibat/src/competition/domain"

type GetAll struct {
	Repository domain.Repository
}

func New() *GetAll {
	return &GetAll{}
}

func (o *GetAll) Execute() map[string]interface{} {
	return map[string]interface{}{
		"field": "value",
	}
}
