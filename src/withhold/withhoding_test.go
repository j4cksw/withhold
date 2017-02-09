package withhold_test

import (
	"testing"
	"withhold"
)

type testData struct {
	WithHoldingRate float64
	Given float64
	Expected float64
}

func prepareTestData() []testData {
	return []testData {
		{
			0,0,0,
		},
		{
			0.05,0,0,
		},
		{
			0.05,100,5,
		},
	}
}

func TestWithHolding_0(t *testing.T) {

	for _, testCase := range prepareTestData() {
		t.Logf("Given witholding rate %.2f amount %.2f should return %.2f", testCase.WithHoldingRate, testCase.Given, testCase.Expected)

		withHolding := &withhold.Controller{
			FindWithHold: prePareFakeFindWithHoldingRate(0.05),
		}

		actual := withHolding.ForAmount(testCase.Given)

		if actual != testCase.Expected {
			t.Errorf("Given %f should return %f but got %d", testCase.Given, testCase.Expected, actual)
		}
	}
}

func prePareFakeFindWithHoldingRate(rate float64) func() withhold.Rate {
	return func () withhold.Rate {
		return withhold.Rate{
			Rate: rate,
		}
	}
}
