package main

import (
	"fmt"
	"sync"
)

func doingSameTask(i int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println(i)
}
func main() {
	// // var wg sync.WaitGroup
	// for i := 0; i < 5; i++ {
	// 	// non blocking : running concurrently
	// 	// go doingSameTask(i)

	// 	go func(i int) {
	// 		fmt.Println(i)
	// 	}(i)
	// }

	// // we don't use that later we use WaitGroup
	// time.Sleep(time.Second * 1)

	// implement wait group

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doingSameTask(i, &wg)
	}

	wg.Wait()

}
