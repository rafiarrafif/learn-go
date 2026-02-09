package conditionals

import "fmt"

func IfElse() {
	grade := 90

	if grade >= 90 {
		fmt.Println("Excellent")
	} else if grade > 60 {
		fmt.Println("Good")
	} else if grade > 40 {
		fmt.Println("Not Bad")
	} else  {
		fmt.Println("Try Again")
	}
}

func Switch() {
	vehicle := "train"
	switch vehicle{
	case "plane":
		fmt.Println("can fly")
	case "train", "boat":
		fmt.Println("can't fly")
	default:
		fmt.Println("vehicle unknown")
	}
}