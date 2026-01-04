package variable

import "fmt"

func FirstWay() {
	var age int
	var salary float64
	var isMarried bool
	var name string

	fmt.Printf("First way: %v %f %v %q\n", age, salary, isMarried, name)
}

func SecondWay() {
	age := 18 // int
	salary := 2997.18 // float64
	isMarried := false // boolean
	name := "Hidup Jokowi" // string
	imajiner := 0.88 + 5i // complex128

	fmt.Printf("Second way: %v %.2f %v %q %T\n", age, salary, isMarried, name, imajiner)
}

func MultipleDeclaration() {
	tickerProportion, tickerName := .89, "MSFT"
	fmt.Println("The price of ticker", tickerName, "is", tickerProportion) 
}

/*
Unless you have a good reason to, stock to the following types:
- bool
- string
- int
- uint32
- byte
- rune
- float64
- complex128
*/

/*
Formatting string:
- %s (print as string)
- %d (print as int)
- %f (print as float, it can also contain format %.2f)
- %t (print as boolean)
- %v (print following variable type, most flexible)
- %+v (print construct and field name)
- %T (print data type)
*/