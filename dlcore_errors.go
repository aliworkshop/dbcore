package dbcore

import "github.com/aliworkshop/error"

var NotFoundErr = error.NotFound(nil).WithCode(404)
var NotEnoughParams = error.Validation(nil).WithCode(422)
