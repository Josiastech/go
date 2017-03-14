package main

import "math/rand"

func RandArray(n int) []int{
	arr := make([]int, n)
	for i := 0; i <= n -1; i++{
		arr[i] = rand.Int(n)
	}
	return arr
}