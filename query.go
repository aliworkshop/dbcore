package dbcore

import "github.com/aliworkshop/dfilter"

type QueryModel interface {
	// SetDB set custom db to query to use in further methods of db handler
	SetDB(db any)
	// GetDB get custom db, returns nil if db is not set
	GetDB() any

	AddFilter(Filterable)
	AddOrFilter(Filterable)
	GetFilters() []Filterable
	GetOrFilters() []Filterable
	AddSort(field string, order order)
	GetSort() (sort []SortItem)
	SetBody(body any)
	GetBody() (body any)
	AddExtraFilter(query string, params ...any)
	GetExtraFilters() []ExtraFilter
	WithJoin(query string, args ...any) QueryModel
	GetModel() (instance Modeler)
	SetModelFunc(func() Modeler)
	SetTransaction(transaction any)
	GetTransaction() (transaction any)
	SetPageSize(pageSize int)
	GetPageSize() int
	SetPage(page int)
	GetPage() (page int)
	GetJoin() []join

	WithModelFunc(func() Modeler) QueryModel
	WithBody(body any) QueryModel
	WithExtraFilter(query string, params ...any) QueryModel
	WithPage(page int) QueryModel
	WithPageSize(pageSize int) QueryModel
	WithFilter(Filterable) QueryModel
	WithOrFilter(Filterable) QueryModel
	WithDynamicFilters([]dfilter.Filter) QueryModel
	GetDynamicFilters() []dfilter.Filter
	WithSort(field string, order order) QueryModel
	WithSorts(sort ...SortItem) QueryModel
	WithTransaction(transaction any) QueryModel
	// Clone copy current query as new instance of query model
	Clone() QueryModel
	// Flush clears body and all filters existing in query and resets
	// other properties. it removes everything except
	// model/models and transaction
	Flush() QueryModel
	WithSelect(columns any, args ...any) QueryModel
	GetSelects() []Select
	GetHint() *Hint
	SetHint(hint Hint)
	SetTable(name string, args ...any) QueryModel
	GetTable() (string, []any)
	WithGroupBy(field string) QueryModel
	GetGroupBy() []string
	SetDynamicFilterTable(string) QueryModel
	GetDynamicFilterTable() string
}

var (
	DefaultPageSize = 30
)

type ModelFunc func() Modeler

type query struct {
	db           any
	filters      []Filterable
	orFilters    []Filterable
	dFilters     []dfilter.Filter
	dFilterTable string
	joins        []join
	modelFunc    ModelFunc
	transaction  any
	extraFilters []ExtraFilter
	pageSize     int
	page         int
	sortItem     []SortItem
	body         any
	extraActions map[string]any
	selects      []Select
	hint         *Hint
	table        struct {
		name string
		args []any
	}
	groupBy []string
}

type Select struct {
	Columns any
	Args    []any
}

type join struct {
	Query string
	Args  []any
}

func NewQuery(existing ...QueryModel) QueryModel {
	var q *query
	if existing != nil && len(existing) > 0 {
		q, _ = existing[0].(*query)
	}
	if q == nil {
		q = &query{}
	}
	q.joins = make([]join, 0)
	q.filters = nil
	return q
}

func IsQueryModel(model any) bool {
	if _, ok := model.(*query); ok {
		return ok
	}
	return false
}

func GetQueryModel(model any) QueryModel {
	if q, ok := model.(*query); ok {
		return q
	}
	return nil
}

func (q *query) SetDB(db any) {
	q.db = db
}

func (q *query) GetDB() any {
	return q.db
}

func (q *query) AddFilter(filter Filterable) {
	q.filters = append(q.filters, filter)
}

func (q *query) AddOrFilter(filter Filterable) {
	q.orFilters = append(q.orFilters, filter)
}

func (q *query) SetBody(body any) {
	q.body = body
}

func (q *query) GetBody() (body any) {
	return q.body
}

func (q *query) AddExtraFilter(filterQuery string, params ...any) {
	if q.extraFilters == nil {
		q.extraFilters = []ExtraFilter{}
	}
	q.extraFilters = append(q.extraFilters, ExtraFilter{
		Query:  filterQuery,
		Params: params,
	})
}

func (q *query) GetFilters() []Filterable {
	return q.filters
}

func (q *query) GetOrFilters() []Filterable {
	return q.orFilters
}

func (q *query) GetExtraFilters() (extraFilters []ExtraFilter) {
	return q.extraFilters
}

func (q *query) WithJoin(query string, args ...any) QueryModel {
	q.joins = append(q.joins, join{Query: query, Args: args})
	return q
}

func (q *query) GetJoin() []join {
	return q.joins
}

func (q *query) GetModel() (instance Modeler) {
	return q.modelFunc()
}

func (q *query) SetModelFunc(modelFunc func() Modeler) {
	q.modelFunc = modelFunc
}

func (q *query) SetTransaction(transaction any) {
	q.transaction = transaction
}

func (q *query) GetTransaction() (transaction any) {
	return q.transaction
}

func (q *query) AddSort(field string, order order) {
	q.sortItem = append(q.sortItem, SortItem{field, order})
}

func (q *query) GetSort() []SortItem {
	return q.sortItem
}

func (q *query) SetPageSize(pageSize int) {
	q.pageSize = pageSize
}

func (q *query) GetPageSize() (pageSize int) {
	if q.pageSize == 0 {
		// default value for page size
		q.pageSize = DefaultPageSize
	}
	return q.pageSize
}

func (q *query) SetPage(page int) {
	q.page = page
}

func (q *query) GetPage() (page int) {
	if q.page == 0 {
		// default value
		q.page = 1
	}
	return q.page
}

func (q *query) GetDynamicFilters() []dfilter.Filter {
	return q.dFilters
}

// with

func (q *query) WithModelFunc(f func() Modeler) QueryModel {
	q.SetModelFunc(f)
	return q
}

func (q *query) WithBody(body any) QueryModel {
	q.SetBody(body)
	return q
}

func (q *query) WithExtraFilter(query string, params ...any) QueryModel {
	q.AddExtraFilter(query, params...)
	return q
}

func (q *query) WithPage(page int) QueryModel {
	q.SetPage(page)
	return q
}

func (q *query) WithPageSize(pageSize int) QueryModel {
	q.SetPageSize(pageSize)
	return q
}

func (q *query) WithFilter(filter Filterable) QueryModel {
	q.AddFilter(filter)
	return q
}

func (q *query) WithOrFilter(filter Filterable) QueryModel {
	q.AddOrFilter(filter)
	return q
}

func (q *query) WithDynamicFilters(filter []dfilter.Filter) QueryModel {
	q.dFilters = filter
	return q
}

func (q *query) WithSort(field string, order order) QueryModel {
	q.AddSort(field, order)
	return q
}

func (q *query) WithSorts(sort ...SortItem) QueryModel {
	for _, sortItem := range sort {
		q.sortItem = append(q.sortItem, sortItem)
	}
	return q
}

func (q *query) WithTransaction(transaction any) QueryModel {
	q.transaction = transaction
	return q
}

func (q *query) Clone() QueryModel {
	var cloned = *q
	return &cloned
}

func (q *query) Flush() QueryModel {
	q.filters = make([]Filterable, 0)
	q.orFilters = make([]Filterable, 0)
	q.dFilters = make([]dfilter.Filter, 0)
	q.dFilterTable = ""
	q.joins = make([]join, 0)
	q.extraActions = make(map[string]any)
	q.body = nil
	q.page = 0
	q.pageSize = 0
	return q
}

func (q *query) WithSelect(columns any, args ...any) QueryModel {
	q.selects = append(q.selects, Select{Columns: columns, Args: args})
	return q
}

func (q *query) GetSelects() []Select {
	return q.selects
}

func (q *query) GetHint() *Hint {
	return q.hint
}

func (q *query) SetHint(hint Hint) {
	q.hint = &hint
}

func (q *query) SetTable(name string, args ...any) QueryModel {
	q.table.name = name
	q.table.args = args
	return q
}

func (q *query) GetTable() (string, []any) {
	return q.table.name, q.table.args
}

func (q *query) WithGroupBy(field string) QueryModel {
	q.groupBy = append(q.groupBy, field)
	return q
}

func (q *query) GetGroupBy() []string {
	return q.groupBy
}

func (q *query) SetDynamicFilterTable(table string) QueryModel {
	q.dFilterTable = table
	return q
}

func (q *query) GetDynamicFilterTable() string {
	return q.dFilterTable
}
