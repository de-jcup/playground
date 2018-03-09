package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	sumByGoRoutinesAndCombinedByChannel()
	bufferedChannel()
	closingChannelExample()
}

// see https://tour.golang.org/concurrency/2
func sumByGoRoutinesAndCombinedByChannel(){
	fmt.Println("--sum by 2 goroutines combined by channel")
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func bufferedChannel(){
	fmt.Println("--buffered channel")
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func closingChannelExample() {
	fmt.Println("--closing channel + use range of channel")
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
