package internal

import (
	helper "KATA_TDD/helper"
	model "KATA_TDD/models"
)

func Add(s string) model.CalculatorResponseModel {
	response := helper.IsStringEmpty(s)

	if response.Error != nil {
		return response
	}

	return model.CalculatorResponseModel{Error: response.Error, Result: response.Result}
}
