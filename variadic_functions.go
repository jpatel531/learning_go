package main 

import "fmt"

// takes an arbitrary number of ints as arguments. like *nums
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1,2)
	sum(1,2,3)

	// initialize a slice
	nums := []int{1,2,3,4}
	
	// pass it in with the ... operator
	sum(nums...)
}