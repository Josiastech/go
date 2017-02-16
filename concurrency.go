package main

import (
	"fmt"
	"sync/atomic"
)

type counter int32

func (c *counter) increment() int32{
	return atomic.AddInt32((*int32)(c), 1)
}

func (c *counter) decrease() int32{
	return atomic.AddInt32((*int32)(c), -1)
}

func (c *counter) get() int32{
	return atomic.LoadInt32((*int32)(c))
}


func main(){
	var count counter
	fmt.Println("Little padawan, the counter is: ", count)
	count.increment()
	fmt.Println("Little padawan, the counter is: ", count)
	count.decrease()
	fmt.Println("Little padawan, the counter is: ", count)
}