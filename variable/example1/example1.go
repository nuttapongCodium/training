package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// Declare variables that are set to their zero value.
	var a int32
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v], size: %d byte\n", a, a, unsafe.Sizeof(a))
	fmt.Printf("var b string \t %T [%v], size: %d byte\n", b, b, unsafe.Sizeof(b))
	fmt.Printf("var c float64 \t %T [%v], size: %d byte\n", c, c, unsafe.Sizeof(c))
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// Specify type and perform a conversion.
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}

// อ้างอิง https://blog.learngoprogramming.com/learn-go-lang-variables-visual-tutorial-and-ebook-9a061d29babe
