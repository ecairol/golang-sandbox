// Channels and concurrency: http://www.golangbootcamp.com/book/concurrency

package main

import "fmt"

func main() {
	// Channel c of type int, buffered 2
	c := make(chan int, 2)
	c <- 1
	c <- 2

	go func() { c <- 3 }()
	go func() { c <- 4 }()
	go func() { c <- 5 }()
	go func() { c <- 6 }()
	go func() { c <- 7 }()
	go func() { c <- 8 }()

	fmt.Println(<-c) // 1
	fmt.Println(<-c) // ..
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c) // 8
	
	//fmt.Println(<-c)  << 9 would be a deadlock!
}
