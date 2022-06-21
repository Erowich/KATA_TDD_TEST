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
