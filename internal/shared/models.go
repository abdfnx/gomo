package shared

const (
	Ready State = iota
	Loading
)

type SetMsg     string
type ErrorMsg   struct{}
type OtherMsg   struct{}
type SuccessMsg struct{}
type Message    struct{ Err error }
type State      int
type Index      int
