package dbcore

type Operator byte

const (
	Equal Operator = iota + 1
	NotEqual
	Lower
	LowerEqual
	Greater
	GreaterEqual
	BitwiseIs
	BitwiseIsNot
	BitwiseAndEqual
	BitwiseAndNotEqual
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
	default:
		return ""
	}
}
