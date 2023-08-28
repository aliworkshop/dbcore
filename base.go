package dbcore

type RDBMS interface {
	GetDB(query ...QueryModel) interface{}
	Repository
	Summable
	Transaction
}

type NoSql interface {
	Repository
	Summable
}
