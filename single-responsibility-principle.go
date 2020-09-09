package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0
type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {

	if len(j.entries) - 1 >= index {
		// do the remove logic here
	}
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\t")
}

// WRONG !!! GOD OBJECT... ANTIPATTERN... SHOULD NEVER HAVE THESE FUNCTIONALITIES IN THIS CLASS
/* Here I am trying to add functionalities of persistence along with this class which is an antipattern
func (j *Journal) SaveToDataStore(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

// trying to handle persistence with the same class
func (j *Journal) LoadFromWeb(url *url.URL) {
	// ...
}
*/


// Using package approach to save to file.. this is still better as we are still moving the persistence function out of Journal
func SaveToFile(j *Journal, seperator, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, seperator)), 0644)
}

func main() {

	j := Journal{entries: make([]string, 0)}
	j.AddEntry("I coded today")
	j.AddEntry("My pubg gameplay is improving")

	fmt.Print(j.String())

	// persistence
	//SaveToFile(&j, "\t", "my_journal.txt")
}