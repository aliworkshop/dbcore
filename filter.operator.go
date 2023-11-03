package dbcore

type Operator byte

const (
	Equal Operator = iota + 1
	NotEqual
	Lower
	LowerEqual
	Greater
	GreaterEqual
	In
	Is
	IsNot
	NotIn
)

func (o Operator) String() string {
	switch o {
	case Equal:
		return "="
	case NotEqual:
		return "!="
	case Lower:
		return "<"
	case LowerEqual:
		return "<="
	case Greater:
		return ">"
	case GreaterEqual:
		return ">="
	case In:
		return "IN"
	case Is:
		return "IS"
	case IsNot:
		return "IS NOT"
	case NotIn:
		return "NOT IN"
	default:
		return ""
	}
}
