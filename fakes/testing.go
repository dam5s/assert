package fakes

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

type FakeTest struct {
	t      *testing.T
	failed bool
	fatal bool
}

func NewFakeTest(t *testing.T) (fakeT *FakeTest) {
	fakeT = new(FakeTest)
	fakeT.t = t
	return
}

func (test *FakeTest) Errorf(format string, args ...interface{}) {
	test.failed = true
}

func (test *FakeTest) Fatalf(format string, args ...interface{}) {
	test.failed = true
	test.fatal = true
}

func (test *FakeTest) ShouldHaveFailed() {
	if !test.failed {
		test.t.Errorf("Expected test to have failed: %s", test.caller())
	}
	test.reset()
}

func (test *FakeTest) ShouldHaveFatallyFailed() {
	if !test.failed || !test.fatal {
		test.t.Errorf("Expected test to have fatally failed: %s", test.caller())
	}
	test.reset()
}

func (test *FakeTest) ShouldHavePassed() {
	if test.failed {
		test.t.Errorf("Expected test to have passed: %s", test.caller())
	}
	test.reset()
}

func (test *FakeTest) reset() {
	test.failed = false
	test.fatal = false
}

func (test *FakeTest) caller() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}
