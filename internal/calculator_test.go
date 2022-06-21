package internal_test

import (
	calculate "KATA_TDD/internal"
	model "KATA_TDD/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add_Numbers_With_Empty_Return_Fail(t *testing.T) {
	res := calculate.Add("")
	assert.Equal(t, model.CalculatorResponseModel{Error: errors.New("input can't be empty"), Result: 0}, res)
}
func Test_Add_Numbers_With_Not_Empty_Return_Success(t *testing.T) {
	res := calculate.Add("1,2,3")
	assert.Equal(t, model.CalculatorResponseModel{Error: nil, Result: 6}, res)
}
func Test_Add_Numbers_With_NewLine_Return_Success(t *testing.T) {
	res := calculate.Add("1\n2,3")
	assert.Equal(t, model.CalculatorResponseModel{Error: nil, Result: 6}, res)
}

func Test_Add_Numbers_With_Support_Different_Delimiters_Return_Success(t *testing.T) {
	res := calculate.Add("//;\n1;2")
	assert.Equal(t, model.CalculatorResponseModel{Error: nil, Result: 3}, res)
}
