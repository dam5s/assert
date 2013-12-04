package assert

import "reflect"

func (w Wrapper) IsEqualTo(expected interface{}) {
	if !w.equals(expected) {
		w.t.Errorf("Expected %#v to equal %#v", w.compared, expected)
	}
}

func (w Wrapper) IsNotEqualTo(unexpected interface{}) {
	if w.equals(unexpected) {
		w.t.Errorf("Expected %#v not to equal %#v", w.compared, unexpected)
	}
}

func (w Wrapper) equals(expected interface{}) bool {
	return reflect.DeepEqual(w.compared, expected)
}
