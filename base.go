package dbcore

import (
	"context"
	"github.com/aliworkshop/errorslib"
	"github.com/shopspring/decimal"
)

type DBModel interface {
	Initialize() errorslib.ErrorModel

	DB() interface{}
	// GetDB try to get transaction from query, otherwise returns gorm db
	GetDB(query ...QueryModel) interface{}
	// Ping checks connectivity and response time of db connection
	Ping(ctx context.Context) errorslib.ErrorModel

	DFilter(query QueryModel) (dbQuery interface{}, filtered bool)

	Insert(query QueryModel) (result interface{}, err errorslib.ErrorModel)

	GetItemsCount(query QueryModel) (count uint64, err errorslib.ErrorModel)
	GetItemsCountWithDFilters(query QueryModel) (count uint64, err errorslib.ErrorModel)
	GetItems(query QueryModel) (items interface{}, err errorslib.ErrorModel)
	GetItemsWithDFilters(query QueryModel) (items interface{}, err errorslib.ErrorModel)
	GetItem(query QueryModel) (item interface{}, err errorslib.ErrorModel)

	Sum(query QueryModel, key string) (decimal.Decimal, errorslib.ErrorModel)
	SumWithDFilters(query QueryModel, key string) (decimal.Decimal, errorslib.ErrorModel)

	Upsert(query QueryModel) (err errorslib.ErrorModel)
	Update(query QueryModel) (err errorslib.ErrorModel)

	Delete(query QueryModel) (err errorslib.ErrorModel)

	BeginTx(ctx context.Context, query QueryModel, args ...interface{}) (err errorslib.ErrorModel)
	StartTransaction(query QueryModel) (err errorslib.ErrorModel)
	CommitTransaction(query QueryModel) (err errorslib.ErrorModel)
	RollbackTransaction(query QueryModel) (err errorslib.ErrorModel)
	GetTransaction(query QueryModel) (transaction interface{})
	// FinalizeTransaction Checks param err
	// If err is not nil, tries to rollback transaction and returns error
	// Otherwise tries to commit transaction and return err if there is any
	FinalizeTransaction(ctx context.Context, query QueryModel, err errorslib.ErrorModel) errorslib.ErrorModel
}

func GetItems(repo DBModel, query QueryModel) (interface{}, errorslib.ErrorModel) {
	if query.GetModels() == nil {
		return nil, errorslib.New().WithType(errorslib.TypeValidation).WithDetail("model is not set in query model")
	}
	err := validatePagination(query)
	if err != nil {
		return nil, err
	}
	// check filters
	return repo.GetItems(query)
}

func validatePagination(q QueryModel) errorslib.ErrorModel {
	if q.GetPage() == 0 {
		return errorslib.New().WithType(errorslib.TypeValidation).WithDetail("page number is 0")
	}
	if q.GetPageSize() == 0 {
		return errorslib.New().WithType(errorslib.TypeValidation).WithDetail("page size is 0")
	}
	return nil
}
