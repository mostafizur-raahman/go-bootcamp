// Go Language: Basic Data Types and Examples

package main

import (
	"fmt"
)

func main() {
	// 1. Numeric Types

	// 1.1 Integers
	var i int = 42         // Default integer type (32 or 64 bits depending on architecture)
	var i8 int8 = -128     // 8-bit signed integer
	var u16 uint16 = 65535 // 16-bit unsigned integer
	fmt.Println("Integers:", i, i8, u16)

	// 1.2 Floating Point Numbers
	var f32 float32 = 3.14         // 32-bit float
	var f64 float64 = 2.7182818284 // 64-bit float (default for decimals)
	fmt.Println("Floats:", f32, f64)

	// 1.3 Complex Numbers
	var c64 complex64 = 1 + 2i   // Complex number with float32 real and imaginary parts
	var c128 complex128 = 2 + 3i // Complex number with float64 parts
	fmt.Println("Complex Numbers:", c64, c128)

	// 2. String Type
	var s string = "Hello, Go!" // String type
	fmt.Println("String:", s)

	// 3. Boolean Type
	var b bool = true // Boolean type
	fmt.Println("Boolean:", b)

	// 4. Derived Types

	// 4.1 Array
	var arr [3]int = [3]int{1, 2, 3} // Fixed-size array
	fmt.Println("Array:", arr)

	// 4.2 Slice
	sl := []int{1, 2, 3, 4} // Dynamic-size slice
	fmt.Println("Slice:", sl)

	// 4.3 Map
	m := map[string]int{"one": 1, "two": 2} // Key-value pairs
	fmt.Println("Map:", m)

	// 4.4 Struct
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "John", Age: 30} // Custom struct
	fmt.Println("Struct:", p)

	// 4.5 Pointer
	x := 42
	pPtr := &x                           // Pointer to an integer
	fmt.Println("Pointer:", pPtr, *pPtr) // Print address and value

	// 5. Constants
	const Pi = 3.14
	fmt.Println("Constant Pi:", Pi)

	// Default Values
	var defaultInt int
	var defaultBool bool
	var defaultString string
	fmt.Println("Default Values:", defaultInt, defaultBool, defaultString)
}
