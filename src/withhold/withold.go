package withhold

type WithHolding struct {
	FindWithHoldRate func () float64
}

func (w *WithHolding)GetWithHolding(amount float64) float64 {
	return amount* w.FindWithHoldRate()
}

func NewWithHolding() *WithHolding {
	return &WithHolding{ getWithHoldRate }
}

