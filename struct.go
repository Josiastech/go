/**
* Go’s structs are typed collections of fields.
* They’re useful for grouping data together to form records.
 */
package main

import "fmt"

// This person struct type has name and age fields.
type person struct {
	name string
	age  int
}

func main() {
	// Creates new struct
	fmt.Println(person{"Bob", 20})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
}
