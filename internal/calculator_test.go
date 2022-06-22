package internal_test

import (
	calculate "KATA_TDD/internal"
	model "KATA_TDD/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add_Numbers_With_Empty_Return_Fail(t *testing.T) {
	ch := make(chan *model.CalculatorResponseModel)
	go calculate.Add("", ch)
	getSupplierResponse := <-ch
	assert.Equal(t, model.CalculatorResponseModel{Error: errors.New("input can't be empty"), Result: 0}, *getSupplierResponse)
}
func Test_Add_Numbers_With_Not_Empty_Return_Success(t *testing.T) {
	ch := make(chan *model.CalculatorResponseModel)
	go calculate.Add("1,2,3", ch)
	getSupplierResponse := <-ch
	assert.Equal(t, model.CalculatorResponseModel{Error: nil, Result: 6}, *getSupplierResponse)
}
func Test_Add_Numbers_With_NewLine_Return_Success(t *testing.T) {
	ch := make(chan *model.CalculatorResponseModel)
	go calculate.Add("1\n2,3", ch)
	getSupplierResponse := <-ch
	assert.Equal(t, model.CalculatorResponseModel{Error: nil, Result: 6}, *getSupplierResponse)
}
func Test_Add_Numbers_With_DifferentFormat_Delimiter_Return_Success(t *testing.T) {
	ch := make(chan *model.CalculatorResponseModel)
	go calculate.Add("//[***]\n1***2***3", ch)
	getSupplierResponse := <-ch
	assert.Equal(t, model.CalculatorResponseModel{Error: nil, Result: 6}, *getSupplierResponse)
}
func Test_Add_Numbers_With_Support_Different_Delimiters_Return_Success(t *testing.T) {
	ch := make(chan *model.CalculatorResponseModel)
	go calculate.Add("//[*][%]\n1*2%3*4", ch)
	getSupplierResponse := <-ch
	assert.Equal(t, model.CalculatorResponseModel{Error: nil, Result: 10}, *getSupplierResponse)
}

func Test_Add_Numbers_With_NegativeNumber_Return_Success(t *testing.T) {
	ch := make(chan *model.CalculatorResponseModel)
	go calculate.Add("//;\n1;-2", ch)
	getSupplierResponse := <-ch
	assert.Equal(t, model.CalculatorResponseModel{Error: errors.New("negatives not allowed"), Result: 0}, *getSupplierResponse)

}

func Test_Add_Numbers_With_NegativeNumbers_Return_Success(t *testing.T) {
	ch := make(chan *model.CalculatorResponseModel)
	go calculate.Add("//[*][%]\n1*2%-3*-4", ch)
	getSupplierResponse := <-ch
	assert.Equal(t, model.CalculatorResponseModel{Error: errors.New("negatives not allowed-3,-4"), Result: 0}, *getSupplierResponse)

}

func Test_Add_Numbers_With_Bigger_Than_1000_Return_Success(t *testing.T) {
	ch := make(chan *model.CalculatorResponseModel)
	go calculate.Add("//[*][%]\n1*2%3*4000*7*1009*7", ch)
	getSupplierResponse := <-ch
	assert.Equal(t, model.CalculatorResponseModel{Error: nil, Result: 20}, *getSupplierResponse)

}
