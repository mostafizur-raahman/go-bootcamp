package main

import (
	"fmt"
)

// func PrintChanNumber(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("Printing channel number ", <-num)
// 		time.Sleep(time.Second)
// 	}

// }

func sum(sumChan chan int, a int, b int) {
	res := a + b
	sumChan <- res
}
func main() {
	// messageChan := make(chan string)

	// // sending data
	// messageChan <- "Hello world" // blocking

	// msg := <-messageChan

	// fmt.Println(msg)

	// numChan := make(chan int)

	// go PrintChanNumber(numChan)

	// // numChan <- 10

	// for {
	// 	num := rand.Intn(100)
	// 	numChan <- num
	// 	time.Sleep(500 * time.Millisecond)
	// }

	// print sum of two result

	sumOfTwoNumberChan := make(chan int)

	go sum(sumOfTwoNumberChan, 10, 20)

	resOfSum := <-sumOfTwoNumberChan

	fmt.Println(resOfSum)

}
