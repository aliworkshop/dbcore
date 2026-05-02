package dbcore

import (
	"context"
	"github.com/aliworkshop/errors"
)

type Transaction interface {
	StartTransaction(query QueryModel) (err errors.ErrorModel)
	CommitTransaction(query QueryModel) (err errors.ErrorModel)
	RollbackTransaction(query QueryModel) (err errors.ErrorModel)
	GetTransaction(query QueryModel) (transaction interface{})
	// FinalizeTransaction Checks param err
	// If err is not nil, tries to rollback transaction and returns error
	// Otherwise tries to commit transaction and return err if there is any
	FinalizeTransaction(ctx context.Context, query QueryModel, err errors.ErrorModel) errors.ErrorModel
}
