package main

import (

	"os"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)


type Zip struct {
	ID string `json:"_id"`
	City string `json:"city"`
	Loc []float64 `json:"loc"`
	Pop int `json:"pop"`
	State string `json:"state"`
}


func main() {

	// connection string
	uri := os.Getenv("MLAB_URI")
	if uri == "" {
		fmt.Println("mlab.com environment variable not working")
		os.Exit(1)
	}

	// start session
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// zips collection
	c := session.DB("zips_production").C("zips")


	// What is the population with San Francisco zipcodes?
	var result []Zip
	err = c.Find(bson.M{"city": "SAN FRANCISCO"}).Select(bson.M{"pop": 1}).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	var cityPopulation = 0
	for _,code := range result {
		cityPopulation += code.Pop
	}

	fmt.Printf("San Francisco zip code population is %v\n", cityPopulation)



}