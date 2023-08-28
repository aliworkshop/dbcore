package dbcore

import (
	"github.com/aliworkshop/error"
	"github.com/shopspring/decimal"
)

type Summable interface {
	Sum(query QueryModel, key string) (decimal.Decimal, error.ErrorModel)
	SumWithDFilters(query QueryModel, key string) (decimal.Decimal, error.ErrorModel)
}
