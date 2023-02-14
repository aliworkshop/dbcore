package dbcore

type HintKind string

const (
	HintForce  = HintKind("force")
	HintUse    = HintKind("use")
	HintIgnore = HintKind("ignore")
)

type Hint struct {
	Kind HintKind
	Name string
}
