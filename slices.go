package main 

import "fmt"

func main() {
	
	// initialize a slice of 3 strings
	s := make([]string, 3)

	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// append to s and assign s to the new slice
	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("apd:", s)

	// copying slices
	c := make([]string, len(s))
	copy(c,s)

	fmt.Println("cpy:", c)

	// 'slice' operator with syntax `slice[low:high]`. gets elements s[2], s[3], s[4]. excludes index 5
	l := s[2:5]
	fmt.Println("sl1:", l)

	// slices up to but excluding 5
	l = s[:5]
	fmt.Println("sl2:", l)

	// slices up from and including s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	// initialize a slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// multidimensional slices

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}