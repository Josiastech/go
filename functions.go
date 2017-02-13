/**
Functions are centarl un GO.
We'll learn about functions with a few different examples
 */
package main


import "fmt"

// Here's a function that takes twi ints and returns their sum as a int
func plus(a int, b int) int {
	// GO requires explicit return, i.e it womt automatically return the value
	// of the last expression
	return a+b
}

// When you have multiple consecutive parameters of the same type
// you may omit the type name fot the like-typed parameters up to final parameter that declares the type
func plusPlus(a,b,c int) int{
	return a + b+ c
}

// Call a function just as you’d expect, with name(args).
func main(){
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1,2,3)
	fmt.Println("1+2+3 = ", res)
}