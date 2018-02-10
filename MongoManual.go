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

type UserStats struct {
	Name string `bson:"_id"`
	Amount int `bson:"value"`
}


func UserCount()int{
	session, err := mgo.Dial(_url)
	if err !=nil{ log.Fatal(err)}
	defer session.Close()
	col := session.DB(_database).C(_collection)

	var result []string
	col.Find(nil).Distinct("user", &result)
	return len(result)
}

func MostActive()[]UserStats{

	session, err := mgo.Dial(_url)
	if err !=nil{ log.Fatal(err)}
	defer session.Close()
	col := session.DB(_database).C(_collection)
	FinalResults := []UserStats{}
	o2 := bson.M{"$group": bson.M{"_id":"$user", "value":bson.M{"$sum":1}}}
	o3 := bson.M{"$sort": bson.M{"value":-1}}
	o4 := bson.M{"$limit":10}
	pipe := []bson.M{o2, o3, o4}
	aggr := col.Pipe(pipe)
	err = aggr.All(&FinalResults)
	if err != nil{
		log.Fatal(err)
	}

	return FinalResults
}


func MostLinks()[]UserStats{

	session, err := mgo.Dial(_url)
	if err !=nil{ log.Fatal(err)}
	defer session.Close()
	col := session.DB(_database).C(_collection)
	FinalResults := []UserStats{}


	o1 := bson.M{"$match":bson.M{"text": bson.M{"$regex":`@\w+`}}}
	o2 := bson.M{"$group": bson.M{"_id":"$user", "value":bson.M{"$sum":1}}}
	o3 := bson.M{"$sort": bson.M{"value":-1}}
	o4 := bson.M{"$limit":10}

	pipe := []bson.M{ o1, o2, o3, o4}

	/// Need work...

	aggr := col.Pipe(pipe)
	err = aggr.Iter().All(&FinalResults)
	if err != nil{
		log.Fatal(err)
	}

	return FinalResults
}

func MostMentions()[]UserStats{
	session, err := mgo.Dial(_url)
	if err !=nil{ log.Fatal(err)}
	defer session.Close()
	col := session.DB(_database).C(_collection)
	FinalResults := []UserStats{}

	deepInter := []interface{}{"$text", " "}
	LessDeepInter := []interface{}{ bson.M{"$split":deepInter}, 0}

	o1 := bson.M{"$match":bson.M{"text": bson.M{"$regex":`@\w+`}}}

	o2 := bson.M{"$project": bson.M{"mentions" : bson.M{ "$arrayElemAt": LessDeepInter}}}

	o3 := bson.M{"$group": bson.M{"_id":"$mentions", "value":bson.M{"$sum":1}}}
	o4 := bson.M{"$sort": bson.M{"value":-1}}
	o5 := bson.M{"$limit":10}

	pipe := []bson.M{ o1, o2, o3, o4, o5}

	/// Need work...

	aggr := col.Pipe(pipe)
	err = aggr.Iter().All(&FinalResults)
	if err != nil{
		log.Fatal(err)
	}

	return FinalResults
}

func MostGrumpy()[]UserStats{
	session, err := mgo.Dial(_url)
	if err !=nil{ log.Fatal(err)}
	defer session.Close()
	col := session.DB(_database).C(_collection)
	FinalResults := []UserStats{}


	o2 := bson.M{"$group": bson.M{"_id":"$user", "emotion":bson.M{"$avg":"$polarity"}}}
	o3 := bson.M{"$sort": bson.M{"value":1}}
	o4 := bson.M{"$limit":10}

	pipe := []bson.M{o2, o3, o4}

	/// Need work...

	aggr := col.Pipe(pipe)
	err = aggr.Iter().All(&FinalResults)
	if err != nil{
		log.Fatal(err)
	}

	return FinalResults
}

func MostHappy()[]UserStats{
	session, err := mgo.Dial(_url)
	if err !=nil{ log.Fatal(err)}
	defer session.Close()
	col := session.DB(_database).C(_collection)
	FinalResults := []UserStats{}


	o2 := bson.M{"$group": bson.M{"_id":"$user", "value":bson.M{"$avg":"$polarity"}}}
	o3 := bson.M{"$sort": bson.M{"emotion":-1}}
	o4 := bson.M{"$limit":10}

	pipe := []bson.M{o2, o3, o4}

	/// Need work...

	aggr := col.Pipe(pipe)
	err = aggr.Iter().All(&FinalResults)
	if err != nil{
		log.Fatal(err)
	}

	return FinalResults
}