package arrays

import "fmt"

func SimpleArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
}

func SimpleAutodefinedArray() {
	odds := []int{1, 3, 5, 7}
	odds = append(odds, 9)
	fmt.Println(odds)
}