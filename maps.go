package main 

import "fmt"

func main() {
	// initializing a 'map', i.e. a hash
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// optional second return value indicates if key was present in map

	_, prs := m["k2"]
	fmt.Println("prs", prs)

	// initialization:

	n := map[string]int{"foo": 1, "bar": 2}

	fmt.Println("map:", n)
}