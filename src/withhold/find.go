package withhold

import (
	"gopkg.in/mgo.v2"
	"log"
)

type Rate struct {
	Rate float64 `bson:"rate"`
}

func FindWithHold() Rate {
	rate := Rate{}

	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
		return rate
	}
	defer session.Close()

	collection := session.DB("taxes").C("withold")
	collection.Find(nil).One(&rate)

	return rate
}
