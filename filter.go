package dbcore

type Operation string

const (
	And Operation = "AND"
	OR            = "OR"
)

type Filter interface {
	GetMatches() []*Match
	GetOperation() Operation
	WithAndMatch(match ...*Match) Filter
	WithOrMatch(match ...*Match) Filter
	WithId(value any) Match
	WithUuid(value any) Match
}

type filter struct {
	operation Operation
	matches   []*Match
}

type Match struct {
	Key      string
	Value    any
	Operator Operator
	Op       Operation
}

func NewFilter(operation ...Operation) Filter {
	op := And
	if len(operation) > 0 {
		op = operation[0]
	}
	return &filter{
		operation: op,
	}
}

func (f *filter) GetMatches() []*Match {
	return f.matches
}

func (f *filter) GetOperation() Operation {
	return f.operation
}

func (f *filter) WithAndMatch(match ...*Match) Filter {
	for _, m := range match {
		m.Op = And
		f.matches = append(f.matches, m)
	}
	return f
}

func (f *filter) WithOrMatch(match ...*Match) Filter {
	for _, m := range match {
		m.Op = OR
		f.matches = append(f.matches, m)
	}
	return f
}

func (f *filter) WithId(value any) Match {
	return Match{
		Key:      "id",
		Value:    value,
		Operator: Equal,
	}
}

func (f *filter) WithUuid(value any) Match {
	return Match{
		Key:      "uuid",
		Value:    value,
		Operator: Equal,
	}
}

type ExtraFilter struct {
	Query  any
	Params []any
}
