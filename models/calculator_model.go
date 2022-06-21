package models

type CalculatorResponseModel struct {
	Result int
	Error  error
}

type DelimiterResponseModel struct {
	Result []string
	Error  error
}
