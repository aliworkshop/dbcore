package dbcore

type Filterable interface {
	Add(key string, value any) Filterable
	Get() Filters
}

type Filters map[string]any

func (f *Filters) Add(key string, value interface{}) {
	(*f)[key] = value
}

func (f *Filters) Delete(key string) {
	delete(*f, key)
}

func (f *Filters) Extend(filters *Filters) {
	if filters == nil {
		return
	}
	for key, value := range *filters {
		f.Add(key, value)
	}
}

type ExtraFilter struct {
	Query  any
	Params []any
}

type filterable struct {
	filters Filters
}

func NewFilterable() Filterable {
	return &filterable{
		filters: map[string]any{},
	}
}

func (f *filterable) Add(key string, value any) Filterable {
	f.filters.Add(key, value)
	return f
}

func (f *filterable) Get() Filters {
	return f.filters
}
