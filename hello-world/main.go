package main

import (
	"fmt"
	"hello-world/arrays"
	"hello-world/conditionals"
	"hello-world/variables"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("========== Variable ==========")
	variables.FirstWay()
	variables.SecondWay()
	variables.MultipleDeclaration()
	fmt.Println("========== Conditional ==========")
	conditionals.IfElse()
	conditionals.Switch()
	fmt.Println("========== Array ==========")
	arrays.SimpleArray()
	arrays.SimpleAutodefinedArray()
}