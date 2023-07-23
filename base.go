package dbcore

import (
	"context"
	"github.com/aliworkshop/error"
	"github.com/shopspring/decimal"
)

type DBModel interface {
	Initialize() error.ErrorModel

	DB() interface{}
	// GetDB try to get transaction from query, otherwise returns gorm db
	GetDB(query ...QueryModel) interface{}
	// Ping checks connectivity and response time of db connection
	Ping(ctx context.Context) error.ErrorModel

	DFilter(query QueryModel) (dbQuery interface{}, filtered bool)

	Insert(query QueryModel) (result interface{}, err error.ErrorModel)

	GetItemsCount(query QueryModel) (count uint64, err error.ErrorModel)
	GetItemsCountWithDFilters(query QueryModel) (count uint64, err error.ErrorModel)
	GetItems(query QueryModel) (items interface{}, err error.ErrorModel)
	GetItemsWithDFilters(query QueryModel) (items interface{}, err error.ErrorModel)
	GetItem(query QueryModel) (item interface{}, err error.ErrorModel)

	Sum(query QueryModel, key string) (decimal.Decimal, error.ErrorModel)
	SumWithDFilters(query QueryModel, key string) (decimal.Decimal, error.ErrorModel)

	Upsert(query QueryModel) (err error.ErrorModel)
	Update(query QueryModel) (err error.ErrorModel)

	Delete(query QueryModel) (err error.ErrorModel)

	BeginTx(ctx context.Context, query QueryModel, args ...interface{}) (err error.ErrorModel)
	StartTransaction(query QueryModel) (err error.ErrorModel)
	CommitTransaction(query QueryModel) (err error.ErrorModel)
	RollbackTransaction(query QueryModel) (err error.ErrorModel)
	GetTransaction(query QueryModel) (transaction interface{})
	// FinalizeTransaction Checks param err
	// If err is not nil, tries to rollback transaction and returns error
	// Otherwise tries to commit transaction and return err if there is any
	FinalizeTransaction(ctx context.Context, query QueryModel, err error.ErrorModel) error.ErrorModel
}

func GetItems(repo DBModel, query QueryModel) (interface{}, error.ErrorModel) {
	if query.GetModel() == nil {
		return nil, error.New().WithType(error.TypeValidation).WithDetail("model is not set in query model")
	}
	err := validatePagination(query)
	if err != nil {
		return nil, err
	}
	// check filters
	return repo.GetItems(query)
}

func validatePagination(q QueryModel) error.ErrorModel {
	if q.GetPage() == 0 {
		return error.New().WithType(error.TypeValidation).WithDetail("page number is 0")
	}
	if q.GetPageSize() == 0 {
		return error.New().WithType(error.TypeValidation).WithDetail("page size is 0")
	}
	return nil
}
