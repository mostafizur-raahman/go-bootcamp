package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	fmt.Println("Error to open file")
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()

	// if err != nil {
	// 	fmt.Println("Error to open file")
	// 	panic(err)
	// }

	// fmt.Println(fileInfo)

	// read file

	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	fmt.Println("Error to open file")
	// 	panic(err)
	// }

	// defer f.Close()

	// buffer := make([]byte, 10)

	// d, err := f.Read(buffer)

	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buffer); i++ {
	// 	fmt.Print(d, string(buffer[i]))
	// }

	// short cut
	// if file is large then we will not use that
	// data, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// read folders

	dir, err := os.Open("../") // current folder if we use (".")

	if err != nil {
		panic(err)
	}

	defer dir.Close()

	fileInfo, err := dir.ReadDir(-1)

	for _, fi := range fileInfo {
		fmt.Println(fi.Name())
	}

}
