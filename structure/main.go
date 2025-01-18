package main

import "fmt"

type Animal interface {
	walk() string
	bark(string) string
}

type Dog struct {
	w, b string
}

func (d *Dog) walk(str string) string {
	d.w = str // change the string through pointer
	return d.w
}

func (d Dog) bark() string {
	return d.b
}
func main() {
	dog := Dog{"Dog is walking", "Dog is barking "}

	fmt.Println(dog.walk("Ok i am chaning now"))
	fmt.Println(dog.bark())
}
