package main

import (
	"fmt"
	"sort"
)

// in order to sort by a custom function
// we need a corresponding type.
// ByLength type is just an alias for the builtin
// []string type
type ByLength []string

// we implement the sort.Interface - Len, Less and Swap
// on our type so we can use the sort package's generic
// sort function
// Len and Swap will usually be similar across types
// and Less will hold the actual custom sorting logic
// in our case we want to sort in order of increasing
// string length, so we use
// len(s[i]) and len(s[j])
// here

// these functions are required by the sort.Interface
// they are used by the sort.Sort function
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// with all of this in place, we can now implement our custom
// sort by casting the originalfruits slice to ByLength
// then use sort.Sort on that typed slice
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
