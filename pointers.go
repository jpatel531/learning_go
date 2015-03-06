package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// takes a pointer, shown by the * operator
func zeroptr(iptr *int) {
	// this deferences the pointer from its memory address
	// to the current value at that address
	// this changes the value at the previously referenced address
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	// this returns 1, as the value of i will not change.
	// #zeroval() has gotten a value for i distinct from the one in #main()
	fmt.Println("zeroval:", i)

	// &i gives the memory address of i, i.e. a pointer to i
	zeroptr(&i)
	fmt.Println("zeroptr", i)

	//printing the pointer
	fmt.Println("pointer", &i)
}
