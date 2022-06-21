package main

import (
	calculate "KATA_TDD/internal"
	"fmt"
)

func main() {

	res := calculate.Add("1\n2,3")
	fmt.Println(res)
}
