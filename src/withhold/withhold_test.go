package withhold_test

import (
	"withhold"
	"testing"
)

func fakeGetWitholdRate() float64 {
	return 0
}

func TestWitholding(t *testing.T) {
	w := withhold.WithHolding{
		fakeGetWitholdRate,
	}

	if w.GetWithHolding(0) != 0 {
		t.Error("Withholding should be 0 but not")
	}
}
