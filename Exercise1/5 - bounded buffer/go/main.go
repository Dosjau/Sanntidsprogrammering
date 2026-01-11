package main

import "fmt"
import "time"


func producer(buffer chan int) {

    for i := 0; i < 10; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("[producer]: pushing %d\n", i)
        buffer <- i // Push value to buffer
    }

}

func consumer(buffer chan int) {

    time.Sleep(1 * time.Second)
    for {
        i := <-buffer // Get value from buffer
        fmt.Printf("[consumer]: %d\n", i)
        time.Sleep(50 * time.Millisecond)
    }
    
}


func main() {
    // Create a bounded buffer with a capacity of 5
    buffer := make(chan int, 5)

    go consumer(buffer)
    go producer(buffer)
    
    select {}
}