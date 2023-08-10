package main

import (
	"fmt"
)

type ErrNegSqrt float64

// Here we turn ErrNegSqrt into an error by adding this implementation:
func (e ErrNegSqrt) Error() string {
	return fmt.Sprintf("That's a negative number you're trying to root: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	fmt.Printf("Taking the square root of %v:\n", x)
	if x < 0 {

		return x, ErrNegSqrt(x)
	}
	var z float64 = 1
	// This method is based on Newton's approximation for finding roots;
	// z = 1 serves as the initial guess.
	for (z*z-x)*(z*z-x) > 0.01 {
		z -= ((z*z - x) / (2 * z))
	}
	return z, nil
}

func main() {
	fmt.Printf("...Example...\n\n")
	fmt.Println(Sqrt(2))  // This one gives you a
	fmt.Println(Sqrt(-2)) // This yields an error, because it's a negative number.
	fmt.Println("----------")

  var i int // declare the var here, so it doesn't get redeclared in the for loop each time.
	// User Input / Interactive Code
	for true { // We place this in a loop, because we need an integer from the user, not string.

		// This one handles the input
		fmt.Print("\nEnter the number you want to take the square root of: ")
		_, e := fmt.Scan(&i)

		// If the user did not input an integer, then the error != nil.
		if e != nil {
			fmt.Printf("%v is not an integer.\nPlease enter an integer (...,0,1,2,3,...).\n", i)
		} else { // If the user inputted an integer, then we need to deliver the square root.
			fmt.Println(Sqrt(float64(i)))
			return
		}
	}
}
