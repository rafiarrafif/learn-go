package maps

import "fmt"

func BasicTypeInsertion() {
	var example any = 123
	str, ok := example.(string)
	if (ok) {
		fmt.Println("Tipe valid, value: ", str)
	} else {
		fmt.Println("Tipe tidak valid")
	}
}

func CommaOK() {
	n := 8
  	if ok, err := isEven(n); ok {
		fmt.Println("It's even number: ", n)
  	} else {
		fmt.Println(err)
  	}
}

func isEven(n int)(bool, error) {
	if n & 1 == 1 {
		return false, fmt.Errorf("it's odd number")
	} else {
		return true, nil
	}
}