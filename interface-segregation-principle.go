package main

import "fmt"

// interface segregation principle states that we should always try to avoid very big interfaces
// and instead break them into smaller interfaces

type Document struct {
}


// this machine interface is perfect we want all the machines to have all these mentioned functionalities
// as shown before a modern age printer can easily have all the capabilities such as printing, faxing and scanning
// but the problem here is that if we are trying to implement an old fashioned printer into the system then this would be a problem
type Machine interface {
	PrintDoc(d Document)
	FaxDoc(d Document)
	ScanDoc(d Document)
}

// new modern age printer having all the functionalities
type MultifunctionPrinter struct {
}

func (mfp *MultifunctionPrinter) PrintDoc(d Document) {
	fmt.Println("I can print")
}

func (mfp *MultifunctionPrinter) FaxDoc(d Document) {
	fmt.Println("I can fax")
}

func (mfp *MultifunctionPrinter) ScanDoc(d Document) {
	fmt.Println("I can scan")
}


// the problem here is that the old fashioned printer cannot fax and scan
type OldPrinter struct {
}

func (op *OldPrinter) PrintDoc(d Document) {
	fmt.Println("I can print")
}

func (op *OldPrinter) FaxDoc(d Document) {
	panic("sorry no scanning facility present")
}

func (op *OldPrinter) ScanDoc(d Document) {
	panic("sorry no scanning facility present")
}

/*
Now as we can see after segregating the interfaces we have solved the problem as a modern printer can as implement all the required interfaces
*/
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Faxxer interface {
	Fax(d Document)
}

/*
Here photocopier being a modern device have implemented 2 out of 3 functionality without having to implement all the 3.
So we can see here how segregating the large interface helps us as now we can select any combination of functions.
*/
type Photocopier struct {
}

func (pc *Photocopier) Print(d Document) {
	fmt.Println("I can print very well")
}

func (pc *Photocopier) Scan(d Document) {
	fmt.Println("I can scan pretty well")
}

/*
We can as well create multifunction interfaces by composition of more than 1 interface into a new interface which is pretty cool
*/
type SuperDevice interface {
	Printer
	Scanner
	Faxxer
}

// Now the best part is if we want we can easily implement this composite interface
type CannonSuper struct {
}

func (cp *CannonSuper) Print(d Document) {
	fmt.Println("Cannon Super is capable of printing a doc")
}

func (cp *CannonSuper) Scan(d Document) {
	fmt.Println("Cannon Super is capable of scanning a doc")
}

func (cp *CannonSuper) Fax(d Document) {
	fmt.Println("Cannon Super is faxing of printing a doc")
}