package main

import "fmt"

func sum(s []int, c chan int) { // note there is no return type here
	sum := 0

	for _, v := range s {
		sum += v
	}

	// send sum to c
	c <- sum
}

func main() {

	// SLICE must be created before use.
	// sums the number 0,1,2,3,4,5,6,7,8,9
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// CHANNEL, like MAPS and SLICES, must be created before use,
	c := make(chan int)

	// distribute the work into two goroutines that run together:
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	a, b := <-c, <-c // receive from the channel
  fmt.Println("Summing 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:")
	fmt.Printf("a: %v, b: %v, a+b: %v \n", a, b, a+b)
}
