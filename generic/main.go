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

// struct
type Stack struct {
	items []int
}

// use generic
type StackG[T any] struct {
	items []T
}

func main() {
	nums := []int{1, 2, 3, 4}
	names := []string{"Mostafiz", "abir", "rayhan"}
	vals := []bool{true, false, false, true}
	printSomething(nums)
	printSomething(names)
	printSomething(vals)

	// struct
	stack := Stack{
		items: []int{1, 2, 4, 5},
	}
	fmt.Println(stack)

	// with generic
	stackG := StackG[int]{
		items: []int{1, 2, 2, 4},
	}

	stackG2 := StackG[string]{
		items: []string{"a", "b", "c"},
	}

	fmt.Println(stackG)
	fmt.Println(stackG2)

}
