package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

var(
	_url string = "192.168.99.100"
	_database string = "social_net"
	_collection string = "tweets"
)

type m bson.M


func UserCount()int{
	session, err := mgo.Dial(_url)
	if err !=nil{ log.Fatal(err)}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	col := session.DB(_database).C(_collection)

	var result []string
	col.Find(nil).Distinct("user", &result)
	return len(result)
}

func MostLinks(){
	session, err := mgo.Dial(_url)
	if err !=nil{ log.Fatal(err)}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	col := session.DB(_database).C(_collection)

	/*

	////Works with mongodb API. Translation needed

	db.tweets.aggregate(
    {$match:{text: /@\w+\// }},
    {$group:{_id:null,text:{$push:"$text"}}
	})

	 */


	o1 := bson.M{"$match":bson.M{"text": bson.RegEx{`/@\w+\//`, ""}},}
	o2 := bson.M{"$group":[]bson.M{{"_id":1},{"text":bson.M{"$push":"$text"}}}}
	pipe := []bson.M{o1, o2}

	/// Need work...

	aggr := col.Pipe(pipe)
	resp := []bson.M{}
	err = aggr.All(&resp)
	if err != nil{
		log.Fatal(err)
	}
	println(len(resp))
}