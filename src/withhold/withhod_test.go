package withhold_test

import (
	"testing"
	"withhold"
)

func fakeFindWithHolding() withhold.Rate {
	return withhold.Rate{
		0.05,
	}
}

func TestWithHolding_0(t *testing.T) {
	withHolding := &withhold.WithHolding{
		fakeFindWithHolding,
	}

	actual := withHolding.CalculateWithHolding(0)

	if actual != 0 {
		t.Errorf("Given 0 should return 0")
	}
}

func TestWithHolding_100(t *testing.T) {
	withHolding := &withhold.WithHolding{
		fakeFindWithHolding,
	}

	actual := withHolding.CalculateWithHolding(100)

	if actual != 5 {
		t.Errorf("Given 100 should return 5 but got %d", actual)
	}
}
