package main

import "fmt"

type ErrNegSqrt float64

// Here we turn ErrNegSqrt into an error by adding this implementation:
func (e ErrNegSqrt) Error() string {
	return fmt.Sprintf("That's a negative number you're trying to root: %v", float64(e))
	// theory: if you put e, it's an error, so the stack will overflow as it continues to execute.
}

func Sqrt(x float64) (float64, error) {
	fmt.Printf("Taking the square root of %v:\n", x)
  if x < 0 {
    
		return x, ErrNegSqrt(x)
	}
	var z float64 = 1
	// This method is based on Newton's approximation for finding roots; z = 1 serves as the initial guess.
	for (z*z-x)*(z*z-x) > 0.01 {
		z -= ((z*z - x) / (2 * z))
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))  // This one gives you a
	fmt.Println(Sqrt(-2)) // This yields an error, because it's a negative number.
}

