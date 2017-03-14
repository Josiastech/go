/**

Interfaces are named collections of method signatures.

 */
package main

import(
	"fmt"
	"math"
)


type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

