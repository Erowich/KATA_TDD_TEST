package main

import (
	calculate "KATA_TDD/internal"
	model "KATA_TDD/models"
	"fmt"
)

func main() {
	ch := make(chan *model.CalculatorResponseModel)
	defer close(ch)
	go calculate.Add("", ch)
	getSupplierResponse := <-ch
	fmt.Println(*getSupplierResponse)

}
