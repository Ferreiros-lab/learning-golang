// Documentation at https://golang.org/pkg/fmt/
package main

import "fmt"

func main() {
	fmt.Println(`Print text with line break at the end`)
	fmt.Print(`Print text without line break at the end`)
	fmt.Println(` Look. Same line!`)

	a := 97
	fmt.Printf("The value of a is: %v\n", a)
	fmt.Printf("The type of a  is: %T\n", a)
	fmt.Printf("The value of a in base 2 is: %b\n", a)
	fmt.Printf("The char corresponding to the value of a is: %c\n", a)
	fmt.Printf("The value of a in base 10 is: %d\n", a)
	fmt.Printf("The value of a in base 8 is: %o\n", a)
	fmt.Printf("The value of a with single quotes is: %q\n", a)
	fmt.Printf("a in base 16 with lower case letters is: :%x\n", a)
	fmt.Printf("a in base 16 with upper case letters is: %X\n", a)
	fmt.Printf("a in unicode format is: %U\n", a)

	b := 123432427.23434345
	fmt.Printf("%b\n", b)
	fmt.Printf("The value of b in scientific notation is: %e\n", b)
	fmt.Printf("The value of b in scientific notation is: %E\n", b)
	fmt.Printf("The value of b in decimal point bun no exponent is: %f\n", b)

	s := "Hello"
	fmt.Printf("Bytes of the string: %s\n", s)
	fmt.Printf("Double quotation of the string: %q\n", s)
	fmt.Printf("Base 16: %x\n", s)

	var i int
	fmt.Print("Introduce an integer: ")
	fmt.Scanf("%d", &i)
	fmt.Println("The integer is", i)
}
