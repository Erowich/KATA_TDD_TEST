package main

import (
	helper "KATA_TDD/internal/helper"
	model "KATA_TDD/internal/models"
	"fmt"
)

func Add(s string) model.CalculatorResponseModel {
	response := helper.IsStringEmpty(s)

	if response.Error != nil {
		return response
	}

	return model.CalculatorResponseModel{Error: response.Error, Result: response.Result}
}

func main() {
	c := Add("123")
	fmt.Println(c)
}
