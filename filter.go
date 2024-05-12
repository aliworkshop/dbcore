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
	WithId(value any, op ...Operation) Filter
	WithUuid(value any, op ...Operation) Filter
}

type filter struct {
	operation Operation
	matches   []*Match
}

type Match struct {
	ElBoost  float32
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
		if len(m.Operator.String()) == 0 {
			m.Operator = Equal
		}
		if m.ElBoost == 0 {
			m.ElBoost = 1
		}
		m.Op = And
		f.matches = append(f.matches, m)
	}
	return f
}

func (f *filter) WithOrMatch(match ...*Match) Filter {
	for _, m := range match {
		if len(m.Operator.String()) == 0 {
			m.Operator = Equal
		}
		if m.ElBoost == 0 {
			m.ElBoost = 1
		}
		m.Op = OR
		f.matches = append(f.matches, m)
	}
	return f
}

func (f *filter) WithId(value any, operation ...Operation) Filter {
	op := And
	if len(operation) > 0 {
		op = operation[0]
	}
	f.matches = append(f.matches, &Match{
		Key:      "id",
		Value:    value,
		Operator: Equal,
		Op:       op,
	})
	return f
}

func (f *filter) WithUuid(value any, operation ...Operation) Filter {
	op := And
	if len(operation) > 0 {
		op = operation[0]
	}
	f.matches = append(f.matches, &Match{
		Key:      "uuid",
		Value:    value,
		Operator: Equal,
		Op:       op,
	})
	return f
}

type ExtraFilter struct {
	Query  any
	Params []any
}
