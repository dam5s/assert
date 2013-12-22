package assert

func (w Wrapper) IsNil() {
	if !w.isNil() {
		w.t.Fatalf("Expected %#v to be nil", w.compared)
	}
}

func (w Wrapper) IsNotNil() {
	if w.isNil() {
		w.t.Fatalf("Expected %#v not to be nil", w.compared)
	}
}

func (w Wrapper) isNil() bool {
	return w.compared == nil
}

