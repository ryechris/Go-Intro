/*
ToDo/Instructions from go.dev/tour/concurrency/8:
1. Implement the `Walk` function

2a. Create a new channel ch and kick off the walker:
`go Walk (NewTree(1), ch)`
2b. Then read and print 10 values from the channel. It should be the numbers 1,2,3,...,10.

3. Implement the `Same` function using `Walk` to determine whether t1 and t2 store the same values.

4. Test the `Same` function:
Same(NewTree(1), NewTree(1)) should return true, and Same(NewTree(1), NewTree(2)) should return false.
*/

package main

import (
	"fmt"
	"math/rand"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	// ToDo #1 Implement the Walk function
	Walker(t, ch)
	close(ch) // notify no more values will be sent; lest errors.
}

func Walker(t *Tree, ch chan int) {
	if t == nil { // empty tree looks like this: ()
		return
	}
	Walker(t.Left, ch)
	ch <- t.Value
	Walker(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	// ToDo #3: implement the `Same` function using `Walk

	// like with maps and slices, make channels before use:
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Goroutine: lightweight thread managed by the Goruntime.
	go Walk(t1, ch1) // this will fill up ch1
	go Walk(t2, ch2) // this will fill up ch2

	for v1 := range ch1 { // chan int only accepts 1 iterator, which is v1 in this case
		v2 := <-ch2
		if v2 != v1 {
			return false
		}
	}

	// in the end this function should return true or false
	return true
}

func main() {
	// ToDo #2
	ch3 := make(chan int)
	go Walk(NewTree(1), ch3)
	for v3 := range ch3 {
		fmt.Print(v3)
	}
	fmt.Println("")

	// ToDo #4
	fmt.Println(Same(NewTree(1), NewTree(1)))
	fmt.Println(Same(NewTree(1), NewTree(2)))
}

// The following is code supplied by the Go authors;
// The code provides a Tree structure.
// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func NewTree(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

