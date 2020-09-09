package main

import "fmt"

type Person struct {
	StreetAddress, Postcode, City string

	CompanyName, Position string
	AnnualIncome float64
}

// This is a simple builder that helps build a person object by default most or all of the values
type PersonBuilder struct {
	person *Person
}

func (pb *PersonBuilder) Build() *Person {
	return pb.person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{
		StreetAddress: "",
		Postcode:      "",
		City:          "",
		CompanyName:   "",
		Position:      "",
		AnnualIncome:  0,
	}}
}

// This is a builder facet which returns a person-address-builder
func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

// This is a builder facet which returns a person-job-builder
func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}


type PersonAddressBuilder struct {
	PersonBuilder
}

func (it *PersonAddressBuilder) City(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (it *PersonAddressBuilder) Postcode(postCode string) *PersonAddressBuilder {
	it.person.Postcode = postCode
	return it
}

func (it *PersonAddressBuilder) Street(address string) *PersonAddressBuilder {
	it.person.StreetAddress = address
	return it
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (jb *PersonJobBuilder) Company(name string) *PersonJobBuilder {
	jb.person.CompanyName  = name
	return jb
}

func (jb *PersonJobBuilder) Earning(income float64) *PersonJobBuilder {
	jb.person.AnnualIncome = income
	return jb
}

func (jb *PersonJobBuilder) Designation(position string) *PersonJobBuilder {
	jb.person.Position = position
	return jb
}

func main() {

	pb := NewPersonBuilder()

	pb.
		Lives().
			City("London").
			Street("22b Baker Street").
			Postcode("123123").
		Works().
			Company("Detectives Inc.").
			Designation("Chief Detective").
			Earning(100000000)

	person := pb.Build()
	fmt.Printf("%+v\n", person)
}