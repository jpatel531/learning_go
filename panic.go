package main

import "os"

func main() {
	panic("a problem")

	// an example of panicking if we get an
	// unexpected error when creating a new file
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

// Note that unlike some languages
// which use exceptions for handling of many errors,
//  in Go it is idiomatic to use error-indicating return values wherever possible.
