package assert_test

import (
	"testing"
	"assert"
	"errors"
	"assert/fakes"
)

func TestIsNil(t *testing.T) {
	fakeT := fakes.NewFakeTest(t)

	assert.That(fakeT, nil).IsNil()
	fakeT.ShouldHavePassed()
	assert.That(fakeT, nil).IsNotNil()
	fakeT.ShouldHaveFatallyFailed()

	var err error

	assert.That(fakeT, err).IsNil()
	fakeT.ShouldHavePassed()
	assert.That(fakeT, err).IsNotNil()
	fakeT.ShouldHaveFatallyFailed()

	err = errors.New("Oops")

	assert.That(fakeT, err).IsNil()
	fakeT.ShouldHaveFatallyFailed()
	assert.That(fakeT, err).IsNotNil()
	fakeT.ShouldHavePassed()
}
