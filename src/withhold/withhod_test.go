package withhold_test

import (
	"testing"
	"withhold"
)

func fakeFindWithHolding() withhold.Rate {
	return withhold.Rate{
		Rate: 0.05,
	}
}

func TestWithHolding_0(t *testing.T) {
	withHolding := &withhold.Controller{
		FindWithHold: fakeFindWithHolding,
	}

	actual := withHolding.ForAmount(0)

	if actual != 0 {
		t.Errorf("Given 0 should return 0 but got %d", actual)
	}
}

func TestWithHolding_100(t *testing.T) {
	withHolding := &withhold.Controller{
		FindWithHold: fakeFindWithHolding,
	}

	actual := withHolding.ForAmount(100)

	if actual != 5 {
		t.Errorf("Given 100 should return 5 but got %d", actual)
	}
}
