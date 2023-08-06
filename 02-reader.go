package main

import "golang.org/x/tour/reader"

// https://cs.opensource.google/go/x/tour/+/refs/tags/v0.1.0:reader/validate.go

// We create the struct here.
type MyReader struct{}

// We turn the struct into a Reader with the following implementation:
func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	// this function from Go tests our code's validity:
	reader.Validate(MyReader{})
}

