package dbcore

import "github.com/aliworkshop/dfilterlib"

type QueryModel interface {
	// SetDB set custom db to query to use in further methods of db handler
	SetDB(db interface{})
	// GetDB get custom db, returns nil if db is not set
	GetDB() interface{}

	Populate(request QueryModel) (populated QueryModel)
	AddFilter(key string, value interface{})
	GetFilter(key string) (value interface{})
	GetFilters() Filters
	RemoveFilter(key string)
	AddSort(field string, order order)
	GetSort() (sort []SortItem)
	SetBody(body interface{})
	GetBody() (body interface{})
	AddExtraFilter(query string, params ...interface{})
	GetExtraFilters() []ExtraFilter
	WithJoin(query string, args ...interface{}) QueryModel
	AddExtraAction(key string, action interface{})
	// TODO: check if thread-safe or not
	GetModel() (instance interface{})
	GetModels() (instance interface{})
	SetModelFunc(func() interface{})
	SetModelsFunc(func() interface{})
	SetTransaction(transaction interface{})
	GetTransaction() (transaction interface{})
	SetPageSize(pageSize int)
	GetPageSize() int
	SetPage(page int)
	GetPage() (page int)
	GetJoin() []join
	GetExtraActions() map[string]interface{}
	// with
	WithModelFunc(func() interface{}) QueryModel
	WithModelsFunc(func() interface{}) QueryModel
	WithBody(body interface{}) QueryModel
	WithExtraFilter(query string, params ...interface{}) QueryModel
	WithPage(page int) QueryModel
	WithPageSize(pageSize int) QueryModel
	WithFilter(key string, value interface{}) QueryModel
	WithDynamicFilters([]dfilterlib.Filter) QueryModel
	GetDynamicFilters() []dfilterlib.Filter
	WithSort(field string, order order) QueryModel
	WithSorts(sort ...SortItem) QueryModel
	// Clone copy current query as new instance of query model
	Clone() QueryModel
	// Flush clears body and all filters existing in query and resets
	// other properties. it removes everything except
	// model/models and transaction
	Flush() QueryModel
	WithSelect(columns interface{}, args ...interface{}) QueryModel
	GetSelects() []Select
	GetHint() *Hint
	SetHint(hint Hint)
	SetTable(name string, args ...interface{}) QueryModel
	GetTable() (string, []interface{})
	WithGroupBy(field string) QueryModel
	GetGroupBy() []string
	SetDynamicFilterTable(string) QueryModel
	GetDynamicFilterTable() string
}

var (
	DefaultPageSize = 30
)

type ModelFunc func() interface{}

var defaultModelFunc ModelFunc = func() interface{} {
	return nil
}

type query struct {
	db           interface{}
	filters      *Filters
	dFilters     []dfilterlib.Filter
	dFilterTable string
	joins        []join
	modelFunc    ModelFunc
	modelsFunc   ModelFunc
	transaction  interface{}
	extraFilters []ExtraFilter
	pageSize     int
	page         int
	sortItem     []SortItem
	body         interface{}
	extraActions map[string]interface{}
	selects      []Select
	hint         *Hint
	table        struct {
		name string
		args []interface{}
	}
	groupBy []string
}

type Select struct {
	Columns interface{}
	Args    []interface{}
}

type join struct {
	Query string
	Args  []interface{}
}

func NewQuery(existing ...QueryModel) QueryModel {
	var q *query
	if existing != nil && len(existing) > 0 {
		q, _ = existing[0].(*query)
	}
	if q == nil {
		q = &query{
			modelFunc:  defaultModelFunc,
			modelsFunc: defaultModelFunc,
		}
	}
	q.joins = make([]join, 0)
	q.filters = nil
	return q
}

func IsQueryModel(model interface{}) bool {
	if _, ok := model.(*query); ok {
		return ok
	}
	return false
}

func GetQueryModel(model interface{}) QueryModel {
	if q, ok := model.(*query); ok {
		return q
	}
	return nil
}

func (q *query) SetDB(db interface{}) {
	q.db = db
}

func (q *query) GetDB() interface{} {
	return q.db
}

func (q *query) Populate(request QueryModel) (populated QueryModel) {
	return request
}

func (q *query) AddFilter(key string, value interface{}) {
	if q.filters == nil {
		q.filters = &Filters{}
	}
	q.filters.Add(key, value)
}

func (q *query) GetFilter(key string) (value interface{}) {
	if q.filters == nil {
		return
	}
	value, _ = (*q.filters)[key]
	return
}

func (q *query) RemoveFilter(key string) {
	if q.filters == nil {
		return
	}
	q.filters.Delete(key)
}

func (q *query) SetBody(body interface{}) {
	q.body = body
}

func (q *query) GetBody() (body interface{}) {
	return q.body
}

func (q *query) AddExtraFilter(filterQuery string, params ...interface{}) {
	if q.extraFilters == nil {
		q.extraFilters = []ExtraFilter{}
	}
	q.extraFilters = append(q.extraFilters, ExtraFilter{
		Query:  filterQuery,
		Params: params,
	})
}

func (q *query) GetFilters() Filters {
	if q.filters == nil {
		return Filters{}
	}
	return *q.filters
}

func (q *query) GetExtraFilters() (extraFilters []ExtraFilter) {
	return q.extraFilters
}

func (q *query) WithJoin(query string, args ...interface{}) QueryModel {
	q.joins = append(q.joins, join{Query: query, Args: args})
	return q
}

func (q *query) GetJoin() []join {
	return q.joins
}

func (q *query) GetModel() (instance interface{}) {
	return q.modelFunc()
}

func (q *query) GetModels() (instance interface{}) {
	return q.modelsFunc()
}

func (q *query) SetModelFunc(modelFunc func() interface{}) {
	q.modelFunc = modelFunc
}

func (q *query) SetModelsFunc(modelsFunc func() interface{}) {
	q.modelsFunc = modelsFunc
}

func (q *query) SetTransaction(transaction interface{}) {
	q.transaction = transaction
}

func (q *query) GetTransaction() (transaction interface{}) {
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

func (q *query) GetDynamicFilters() []dfilterlib.Filter {
	return q.dFilters
}

func (q *query) AddExtraAction(key string, action interface{}) {
	if q.extraActions == nil {
		q.extraActions = map[string]interface{}{}
	}
	q.extraActions[key] = action
}

func (q *query) GetExtraActions() map[string]interface{} {
	return q.extraActions
}

// with

func (q *query) WithModelFunc(f func() interface{}) QueryModel {
	q.SetModelFunc(f)
	return q
}

func (q *query) WithModelsFunc(f func() interface{}) QueryModel {
	q.SetModelsFunc(f)
	return q
}

func (q *query) WithBody(body interface{}) QueryModel {
	q.SetBody(body)
	return q
}

func (q *query) WithExtraFilter(query string, params ...interface{}) QueryModel {
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

func (q *query) WithFilter(key string, value interface{}) QueryModel {
	q.AddFilter(key, value)
	return q
}

func (q *query) WithDynamicFilters(filter []dfilterlib.Filter) QueryModel {
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

func (q *query) Clone() QueryModel {
	var cloned = *q
	return &cloned
}

func (q *query) Flush() QueryModel {
	q.filters = &Filters{}
	q.dFilters = make([]dfilterlib.Filter, 0)
	q.joins = make([]join, 0)
	q.extraActions = make(map[string]interface{})
	q.body = nil
	q.page = 0
	q.pageSize = 0
	return q
}

func (q *query) WithSelect(columns interface{}, args ...interface{}) QueryModel {
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

func (q *query) SetTable(name string, args ...interface{}) QueryModel {
	q.table.name = name
	q.table.args = args
	return q
}

func (q *query) GetTable() (string, []interface{}) {
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
