package main

import (
	"fmt"
	"time"
)

func doingSameTask(i int) {
	fmt.Println(i)
}
func main() {
	// var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		// non blocking : running concurrently
		// go doingSameTask(i)

		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	// we don't use that later we use WaitGroup
	time.Sleep(time.Second * 1)

}
