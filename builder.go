package main

import (
	"fmt"
	"strings"
)

/*
Creation of some structs are very simple, while some are not.
In some situations object creation is not very simple.

eg - creation of an object requires 10 arguments. In this case we can build the object step by step
in opposed to creating the object all at once.

When piecewise object construction is complicated, provide an API for doing it succinctly.
*/


func main() {
	hello := "hello"

	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())


	words := []string{"hello", "world"}
	sb.Reset()

	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())
}