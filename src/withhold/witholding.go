package withhold

type Controller struct {
	FindWithHold func () Rate
}

func (w *Controller) ForAmount(amount float64) float64 {
	return w.multiplyWithHoldingRate( amount, w.FindWithHold().Rate )
}

func (w *Controller) multiplyWithHoldingRate(amount float64, withHoldingRate float64) float64 {
	return amount * withHoldingRate
}

func NewController() *Controller {
	return &Controller{FindWithHold }
}