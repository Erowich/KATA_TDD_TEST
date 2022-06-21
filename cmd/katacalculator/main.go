package main

import (
	calculate "KATA_TDD/internal"
	"fmt"
)

func main() {

	res := calculate.Add("//[***]\n1***2000***-3")
	fmt.Println(res)
}
