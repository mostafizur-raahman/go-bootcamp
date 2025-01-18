package main

import "fmt"

func main() {
	// একটি অ্যারে তৈরি করুন এবং এর দৈর্ঘ্য প্রিন্ট করুন

	var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(len(arr))

	// find x  of a list of array
	x := 10
	for i := 0; i < len(arr); i++ {
		if x == arr[i] {
			fmt.Println("found at index", i)
		}
	}

	// Implemant binary search

	searchValue := 100

	lo := 0
	high := len(arr) - 1

	for lo <= high {
		mid := (lo + high) >> 1

		if arr[mid] == searchValue {
			fmt.Println("Find at index ", mid)
		} else if arr[mid] < searchValue {
			lo = mid + 1
		} else {
			high = mid - 1
		}
	}

}
