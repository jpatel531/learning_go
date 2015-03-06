// interfaces are named collections of method signatures

package main

import "fmt"

import "math"

// basic interface for geometry shapes
type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

// to implement an interface, we just need to implement
// all the methods in the interface
// Here we implement geometry on squares
func (s square) area() float64 {
	return s.width * s.height
}

func (s square) perim() float64 {
	return 2*s.width + 2*s.height
}

//now on circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// if a variable has an interface type, then we can call methods
// that are in the named interface.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	s := square{width: 3, height: 4}
	c := circle{radius: 5}

	//circle and square struct types both implement geometry
	// interface so we can use instances of these structs
	// as arguments to `measure()`
	measure(s)
	measure(c)
}
