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

func TestEqualityWithComplexValues(t *testing.T) {
	type person struct {
		name     string
		children []person
	}

	children := []person{{name: "Damien"}, {name: "Davy"}, {name: "Sam"}}
	dad := person{
		name:     "Jean-Claude",
		children: children,
	}
	mum := person{
		name:     "Véronique",
		children: children,
	}

	fakeT := fakes.NewFakeTest(t)

	assert.That(fakeT, dad).IsEqualTo(dad)
	fakeT.ShouldHavePassed()

	assert.That(fakeT, dad).IsEqualTo(mum)
	fakeT.ShouldHaveFailed()

	assert.That(fakeT, dad).IsNotEqualTo(mum)
	fakeT.ShouldHavePassed()
}
