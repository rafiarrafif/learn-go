package structtags

import (
	"fmt"
	"reflect"
)

type Student struct {
	Fullname string
	Email string
	Age int 
	Grade float32 `about:"Cummulative grade from first exam until now (including homework task)"`
}

func TagExample() {
	student := Student{
		Fullname: "Kaizuka Yuki",
		Email: "kaizuka@astofo.com",
		Age: 18,
		Grade: 99.99,
	}

	if data, ok := reflect.TypeOf(student).FieldByName("Grade"); ok {
		fmt.Printf("Grade is %v. Good bye\n", data.Tag.Get("about"))
	} else {
		fmt.Println("Unknown grade")
	}
}