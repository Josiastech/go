/**
Switch statements express confitionals
accros many branches
 */
package main

import (
	"fmt"
	"time"
)


func main(){
	i := 2
	fmt.Print("Write ", i, " as ")

	// basic Switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// You can use commas to separate multiple expressions
	// in the same case statement. We use optional default
	// case in this example as well.
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("Weekday")
	}

	// Switch without an expression is an alternate way to
	// express if/else logic.

	t:= time.Now()

	//
	switch {
	case t.Hour() < 12:
		fmt.Println("its before noon")
	default:
		fmt.Println("its afet noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}