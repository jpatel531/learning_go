package main 

import "fmt"

// (int, int) shows it returns two ints

func vals()(int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

}