package withhold

import (
	"gopkg.in/mgo.v2"
	"log"
)

type Rate struct {
	Rate float64 `bson:"rate"`
}

type WithHolding struct {
	FindWithHold func () Rate
}

func (w *WithHolding)CalculateWithHolding(amount float64) float64 {
	withHold := w.FindWithHold()
	return amount * withHold.Rate
}

func NewWithHolding() *WithHolding {
	return &WithHolding{ FindWithHold }
}

func FindWithHold() Rate {
	withHold := Rate{}

	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
		return withHold
	}
	defer session.Close()

	collection := session.DB("taxes").C("withold")
	collection.Find(nil).One(&withHold)

	return withHold
}