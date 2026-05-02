package dbcore

import (
	"context"
	"github.com/aliworkshop/errors"
)

type Repository interface {
	Initialize() errors.ErrorModel
	DB() any
	Ping(ctx context.Context) errors.ErrorModel

	Count(query QueryModel) (count uint64, err errors.ErrorModel)
	CountWithDFilter(query QueryModel) (count uint64, err errors.ErrorModel)
	List(query QueryModel) (items interface{}, err errors.ErrorModel)
	ListWithDFilter(query QueryModel) (items interface{}, err errors.ErrorModel)
	Get(query QueryModel) (item interface{}, err errors.ErrorModel)
	Exist(query QueryModel) (exists bool, err errors.ErrorModel)

	Insert(query QueryModel) (result interface{}, err errors.ErrorModel)
	Upsert(query QueryModel) (err errors.ErrorModel)
	Update(query QueryModel) (err errors.ErrorModel)
	Delete(query QueryModel) (err errors.ErrorModel)
}
