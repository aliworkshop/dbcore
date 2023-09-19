package dbcore

type Operation string

const (
	And Operation = "and"
	Or            = "or"
)

type Filter interface {
	GetMatches() []*Match
	GetOperation() Operation
	WithMatch(match *Match) Filter
	WithId(value any) Filter
}

type filter struct {
	operation Operation
	matches   []*Match
}

type Match struct {
	Key      string
	Value    any
	Operator Operator
}

func NewFilter(operation Operation) Filter {
	return &filter{
		operation: operation,
	}
}

func (f *filter) GetMatches() []*Match {
	return f.matches
}

func (f *filter) GetOperation() Operation {
	return f.operation
}

func (f *filter) WithMatch(match *Match) Filter {
	f.matches = append(f.matches, match)
	return f
}

func (f *filter) WithId(value any) Filter {
	f.matches = append(f.matches, &Match{
		Key:      "id",
		Value:    value,
		Operator: Equal,
	})
	return f
}

type ExtraFilter struct {
	Query  any
	Params []any
}
