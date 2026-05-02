package dbcore

import (
	"github.com/aliworkshop/errors"
	"github.com/shopspring/decimal"
)

type Summable interface {
	Sum(query QueryModel, key string) (decimal.Decimal, errors.ErrorModel)
	SumWithDFilters(query QueryModel, key string) (decimal.Decimal, errors.ErrorModel)
}
