package main 

import "fmt"


// takes two ints and returns an int
// requires explicit returns
func plus(a int, b int) int{
	return a + b
}


// if all params are of same type, omit it until the last one
func plusPlus(a, b, c int)int {
	return a + b + c
}

func main() {
	res := plus(1,2)
	fmt.Println("1+2 =", res)


	//  := declares and initializes a new var
	res = plusPlus(1,2,3)
	fmt.Println("1+2+3 =", res)
}