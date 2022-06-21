package models

type CalculatorResponseModel struct {
	Error  error
	Result int
}

type DelimiterResponseModel struct {
	Error  error
	Result []string
}
