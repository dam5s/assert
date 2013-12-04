package assert

type TestInterface interface {
	Errorf(format string, args ...interface{})
}
