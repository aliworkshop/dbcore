package dbcore

type QueryParseResult interface {
	GetQuery() (query interface{})
	GetParams() (params []interface{})
	GetSort() (sort interface{})
}

type QueryParser interface {
	Parse(query QueryModel, extra ...interface{}) (result QueryParseResult)
}
