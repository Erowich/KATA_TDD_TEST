package helper

import (
	model "KATA_TDD/models"
	"errors"
)

func IsStringEmpty(s string) model.CalculatorResponseModel {
	if len(s) == 0 {
		return model.CalculatorResponseModel{Error: errors.New("input can't be empty"), Result: 0}
	}
	return model.CalculatorResponseModel{Error: nil, Result: 1}

}
