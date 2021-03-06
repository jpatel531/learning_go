// defer is used to ensure that a function
// call is performed later in a program's execution
// usually for the purposes of a cleanup
// defer is often used where e.g. `ensure` or `finally`
// is used in other languages

package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("/tmp/.defer.txt")

	defer closeFile(f) // this will be executed at the end of the closing function `main`

	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
