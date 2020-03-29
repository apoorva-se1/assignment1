package assignment

import "fmt"

func Constants() {
	var x string = "Hello World"
	fmt.Printf("x is of type %T\n", x)

	const y string = "Hello Go"
	fmt.Printf("y is of type %T\n", y)
	var z int = 20.0
	fmt.Printf("z is of type %T\n", z)
	t := 42
	fmt.Printf("t is of type %T\n", t)
	var a, b, c = 3, 4, "foo"
	fmt.Printf("a,b,c is of type %T, %T, %T\n", a, b, c)
}
