package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func dieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}

func rollTwo(size1, size2 int) (int, int) {

	return dieRoll(size1), dieRoll(size2)

}
func returnsNamed(input1 string, input2 int) (theResult string, err error) {
	theResult = "modified" + input1 + "," + strconv.Itoa(input2)
	return theResult, err
}

func main() {
	fmt.Printf("Rolled a die of zise %d, result: %d\n", 6, dieRoll(6))

	res1, res2 := rollTwo(6, 10)
	fmt.Printf("Rolled pair of dice(%d, %d), results: %d,%d\n",
		6, 10, res1, res2)

	named, err := returnsNamed("globule", 42)
	fmt.Printf("named params returned: '%s', %v\n", named, err)
}