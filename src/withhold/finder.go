package withhold

import (
	"gopkg.in/mgo.v2"
	"log"
)

type WithHoldRate struct {
	Rate float64 `bson:"rate"`
}

func getWithHoldRate() float64 {
	w := WithHoldRate{}

	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
		return w.Rate
	}
	defer session.Close()

	collection := session.DB("taxes").C("withold")
	collection.Find(nil).One(&w)

	return w.Rate
}
