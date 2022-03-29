package main

import (
	"fmt"

	"github.com/inputvalidation/student"
)

func main() {

	std1 := student.Student{}
	errs := std1.IsValid()
	fmt.Println(errs)

}
