package main

import "fmt"

// func printSomething(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// use generic
func printSomething[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
func main() {
	nums := []int{1, 2, 3, 4}
	names := []string{"Mostafiz", "abir", "rayhan"}
	printSomething(nums)
	printSomething(names)
}
