// structs are typed collections of fields

package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// instantiation
	fmt.Println(person{"Bob", 20})

	//you can name fields
	fmt.Println(person{name: "Alice", age: 30})

	//omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	//to yield a pointer to the struct:

	fmt.Println(&person{name: "Ann", age: 40})

	// access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	//you can also use dots with struct pointers
	// the pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

}
