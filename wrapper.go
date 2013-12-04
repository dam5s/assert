package assert

type Wrapper struct {
	t        TestInterface
	compared interface{}
}

func That(t TestInterface, compared interface{}) (w Wrapper) {
	w.t = t
	w.compared = compared
	return
}
