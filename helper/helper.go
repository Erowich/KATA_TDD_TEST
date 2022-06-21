package helper

import (
	model "KATA_TDD/models"
	"errors"
	"strconv"
	"strings"
)

func IsStringEmpty(s string) model.CalculatorResponseModel {
	if len(s) == 0 {
		return model.CalculatorResponseModel{Error: errors.New("input can't be empty"), Result: 0}
	}
	return model.CalculatorResponseModel{Error: nil, Result: 1}

}

func DetectNegativeNumbers(sArray []string) model.CalculatorResponseModel {
	negativeNumbers := []string{}
	for _, num := range sArray {
		number, err := strconv.Atoi(num)
		if err == nil && number < 0 {
			negativeNumbers = append(negativeNumbers, num)
		}
	}
	if len(negativeNumbers) == 0 {
		return model.CalculatorResponseModel{
			Error:  nil,
			Result: 0,
		}
	} else if len(negativeNumbers) == 1 {
		return model.CalculatorResponseModel{
			Error:  errors.New("negatives not allowed"),
			Result: 0,
		}
	} else {
		return model.CalculatorResponseModel{
			Error:  errors.New("negatives not allowed" + strings.Join(negativeNumbers, ",")),
			Result: 0,
		}
	}
}
