package dbcore

import (
	"context"
	"github.com/aliworkshop/error"
)

type Repository interface {
	Initialize() error.ErrorModel
	DB() any
	Ping(ctx context.Context) error.ErrorModel

	Count(query QueryModel) (count uint64, err error.ErrorModel)
	CountWithDFilter(query QueryModel) (count uint64, err error.ErrorModel)
	List(query QueryModel) (items interface{}, err error.ErrorModel)
	ListWithDFilter(query QueryModel) (items interface{}, err error.ErrorModel)
	Get(query QueryModel) (item interface{}, err error.ErrorModel)
	Exist(query QueryModel) (exists bool, err error.ErrorModel)

	Insert(query QueryModel) (result interface{}, err error.ErrorModel)
	Upsert(query QueryModel) (err error.ErrorModel)
	Update(query QueryModel) (err error.ErrorModel)
	Delete(query QueryModel) (err error.ErrorModel)
}
