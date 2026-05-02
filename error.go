package dbcore

import "github.com/aliworkshop/errors"

var NotFoundErr = errors.NotFound(nil).WithCode(404)
var NotEnoughParams = errors.Validation(nil).WithCode(422)
