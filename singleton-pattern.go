package main

import (
	"fmt"
	"sync"
)

/*
Motivation
- For some components it only makes sense to have one in the system like Database repository, Object factory
- Constructor call is expensive
- We want everyone to have the same instance
- Want to prevent anyone from creating additional copies
- Need to take care of lazy instantiation
*/

// Singleton is a component which is instantiated only once


// private database that holds city names with their population
type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// thread safety and laziness are needed to be taken care of.
// now the package level init() does not take care of laziness hence we will use sync.Once

// ensures that this is called only once
var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		// bla bla bla.... construct the
		db := singletonDatabase{}
		capitals := make(map[string]int)
		capitals["Seoul"] = 5000
		capitals["Mumbai"] = 1000000
		capitals["New Delhi"] = 798922
		capitals["New York"] = 58912211
		db.capitals = capitals
		instance = &db
	})

	return instance
}


/* Problems with Singleton Design Pattern
 - breaks dependency inversion principle
*/

// problem description
// trying to get the populations of few cities
func GetTotalPopulation(cities []string) int {
	tot := 0
	for _, city := range cities {
		tot += GetSingletonDatabase().GetPopulation(city)
	}
	return tot
}

// Introducing an interface so as to be able to abstract it
type Database interface {
	GetPopulation(name string) int
}

// making a dummy database to be injected into and perform unit tests
type DummyDatabase struct {
	dummyData map[string]int
}

func (dd *DummyDatabase) GetPopulation(name string) int {

	if len(dd.dummyData) == 0 {
		dd.dummyData = map[string]int{
			"aplha": 1,
			"beta": 2,
			"gamma": 3,
		}
	}
	return dd.dummyData[name]
}

func GetTotalPopulationAbs(db Database, cities []string) int {
	tot := 0
	for _, city := range cities {
		tot += db.GetPopulation(city)
	}
	return tot
}


func main() {

	db := GetSingletonDatabase()
	population := db.GetPopulation("New York")
	fmt.Println("Population of New York: ", population)

	cities := []string{"Seoul", "New York"}
	tp := GetTotalPopulation(cities)
	fmt.Println("Total Population: ", tp)
	fmt.Println("We can verify the data to be correct.")

	fmt.Println("Using this pattern testing is difficult, because we are depending on the concrete implementation and this VIOLATES DEPENDENCY INVERSION")
	fmt.Println("Cannot abstract here in the above example")

	fmt.Println("Now we will try to implement abstraction so as to be able to have dependency inversion")


	fmt.Println("performing dummy database")
	cities = []string{"alpha", "gamma"}
	tp = GetTotalPopulationAbs(&DummyDatabase{}, cities)
	fmt.Println("Total Population: ", tp)
	fmt.Println("We can verify the data to be correct.")
}

/*
Summary
- We should not depend directly upon a singleton as in itself it will violate Dependency Inversion
- We should use abstract interface to depend to, so that we could use any other implementation of it !
*/