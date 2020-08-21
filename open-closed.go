package main

import (
	"fmt"
	"strconv"
)

// open for extension but closed for modification
// also called Specification which is an enterprise pattern

// selling widgets

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

// filter is used to filter the products upon some specific parameters such as colors and sizes
// THE PROBLEM WITH THIS IS THAT WE WOULD NEED TO MODIFY THIS CLASS FOR DIFFERENT FUNCTIONALITY
type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {

	res := make([]*Product, 0)
	for i, _ := range products {
		if products[i].color == color {
			res = append(res, &products[i])
		}
	}

	return res
}

type Product struct {
	name string
	color Color
	size Size
}

func (p *Product) String() string {
	return "Product: name=" + p.name + "\tcolor=" + strconv.Itoa(int(p.color)) + "\tSize=" + strconv.Itoa(int(p.size))
}

// this is the specification interface
type Specification interface {
	IsSatisfied(p *Product) bool
}

// implementation of the interface to filter by color
type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return c.color == p.color
}

// implementation of the interface to filter by size
type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return s.size == p.size
}

// COMPOSITE Specification
type AndSpecification struct {
	first, second Specification
}

func (as AndSpecification) IsSatisfied(p *Product) bool {
	return as.first.IsSatisfied(p) && as.second.IsSatisfied(p)
}

// Better Filter is used to implement specification pattern which takes in a specification to filter
// Thus we simply need to plug in any specification of our choice
type BetterFilter struct {
}

func (bf *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, _ := range products {
		if spec.IsSatisfied(&products[i]) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}
	soap := Product{"Glycerine", blue, small}
	brush := Product{"Colgate", blue, medium}

	products := []Product{apple, tree, house, soap, brush}

	// creating a standard filter whose filtering capabilities are just fixed and to modify it we would need to extend the class
	f := Filter{}
	fmt.Printf("Green Products (old): \n")
	for _, v := range f.FilterByColor(products, blue) {
		fmt.Println(v.String())
	}

	// better filter simply takes in a specification and modifies the filter accordingly thus we can plug in any specification
	// without having to MODIFY the filter functionality
	fmt.Printf("Green Products (new): \n")
	bf := BetterFilter{}
	grenSpec := ColorSpecification{green}
	for _, v := range bf.Filter(products, grenSpec) {
		fmt.Println(v.String())
	}

	largeGreenSpec := AndSpecification{ColorSpecification{green}, SizeSpecification{large}}
	fmt.Printf("Green Large Products (new): \n")
	for _, v := range bf.Filter(products, largeGreenSpec) {
		fmt.Println(v.String())
	}
}