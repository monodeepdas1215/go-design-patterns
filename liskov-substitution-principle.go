package main

import "fmt"

// Liskov Substituion Principle primary deals with inheritance and hence is not that applicable in Go.
// But still a slight modification of it can be done.


// creating an interface to inherit from
type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(width int)
}

// inheriting from the interface
type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetWidth(w int) {
	r.width = w
}

func (r *Rectangle) SetHeight(h int) {
	r.height = h
}

// This definition of square will break Liskov substitution principle as
// we could see when we would use this this func UseIt() it will give wrong results
type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = s.height
}

func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

// improved square object that will not violate liskov substitution principle.
type Square2 struct {
	size int // width and height
}

func (s *Square2) Rectangle() *Rectangle {
	return &Rectangle{s.size, s.size}
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)

	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Println("Expected: ", expectedArea, "\t but got: ", actualArea)
}

func main() {

	rc := &Rectangle{2, 3}
	UseIt(rc)
	fmt.Println("So far so good. everything works as expected")

	sq := NewSquare(5)
	UseIt(sq)
	fmt.Println("As we can see things became wrong here")
	fmt.Println("Liskov substitution principle states that when we generalize and make assumptions on top of that generalizations" +
		" then after we inherit the functionalities in the future the assumptions made earlier should not break.")
	fmt.Println("But as in this case it breaks")

	sq2 := Square2{9}
	UseIt(sq2.Rectangle())
	fmt.Println("again everything works as expected now")
}