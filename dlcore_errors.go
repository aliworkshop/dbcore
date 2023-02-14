package dbcore

import "github.com/aliworkshop/errorslib"

var NotFoundErr = errorslib.NotFound(nil).WithCode(100)
var NotEnoughParams = errorslib.Validation(nil).WithCode(101)
