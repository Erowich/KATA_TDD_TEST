package internal

import (
	helper "KATA_TDD/helper"
	model "KATA_TDD/models"
	"errors"
	"strconv"
	"strings"
)

func Add(s string) model.CalculatorResponseModel {
	response := helper.IsStringEmpty(s)

	if response.Error != nil {
		return response
	}
	splitedStringArray := strings.Split(s, ",")

	return calculate(splitedStringArray)
}

func calculate(sArray []string) model.CalculatorResponseModel {
	sum := 0
	for _, num := range sArray {
		num, err := strconv.Atoi(num)
		if err != nil {
			return model.CalculatorResponseModel{
				Error:  errors.New("can not parse string to int"),
				Result: 0,
			}
		}
		sum += num

	}
	return model.CalculatorResponseModel{
		Error:  nil,
		Result: sum,
	}
}
