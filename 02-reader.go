package main

import (
  "fmt"
  "io"
  "os"
)

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
  Validate(MyReader{})
}

/*******************************************************/
// The following code was supplied by the Go Authors
// to validate that our MyReader struct is indeed a Reader,
// one that emits an infinite stream of 'A'

// https://cs.opensource.google/go/x/tour/+/refs/tags/v0.1.0:reader/validate.go
// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func Validate(r io.Reader) {
  b := make([]byte, 1024, 2048)
  i, o := 0, 0
  for ; i < 1<<20 && o < 1<<20; i++ { // test 1mb
    n, err := r.Read(b)
    for i, v := range b[:n] {
      if v != 'A' {
        fmt.Fprintf(os.Stderr, "got byte %x at offset %v, want 'A'\n", v, o+i)
        return
      }
    }
    o += n
    if err != nil {
      fmt.Fprintf(os.Stderr, "read error: %v\n", err)
      return
    }
  }
  if o == 0 {
    fmt.Fprintf(os.Stderr, "read zero bytes after %d Read calls\n", i)
    return
  }
  fmt.Println("OK!")
}

