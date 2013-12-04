package assert_test

import (
	"assert"
	"assert/fakes"
	"testing"
)

func TestEqualityWithSimpleValues(t *testing.T) {
	fakeT := fakes.NewFakeTest(t)

	assert.That(fakeT, 1).IsEqualTo(1)
	fakeT.ShouldHavePassed()

	assert.That(fakeT, 1).IsEqualTo(2)
	fakeT.ShouldHaveFailed()

	assert.That(fakeT, 1).IsNotEqualTo(2)
	fakeT.ShouldHavePassed()

	assert.That(fakeT, 1).IsNotEqualTo(1)
	fakeT.ShouldHaveFailed()
}
