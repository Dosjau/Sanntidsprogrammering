// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
)


func server(inc <-chan int, dec <-chan int, stop <-chan int, result chan<- int) {
	i := 0
	for {
		select {
		case <-inc:
			i++
		case <-dec:
			i--
		case <-stop:
			result <- i
			return
		}
	}
}

func incrementing(inc chan<- int, done chan<- int) {
	for j := 0; j < 1000005; j++ {
		inc <- 1 
	}
	done <- 1
}

func decrementing(dec chan<- int, done chan<- int) {
	for j := 0; j < 1000000; j++ {
		dec <- 1 
	}
	done <- 1
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	inc := make(chan int)
	dec := make(chan int)
	stop := make(chan int)
	result := make(chan int)

	incDone := make(chan int)
	decDone := make(chan int)

	go server(inc, dec, stop, result)
	go incrementing(inc, incDone)
	go decrementing(dec, decDone)

	
	incFinished := false
	decFinished := false
	for !(incFinished && decFinished) {
		select {
		case <-incDone:
			incFinished = true
		case <-decDone:
			decFinished = true
		}
	}

	stop <- 1
	i := <-result
	Println("The magic number is:", i)
}
