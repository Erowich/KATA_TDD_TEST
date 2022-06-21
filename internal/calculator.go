package internal

import (
	helper "KATA_TDD/helper"
	model "KATA_TDD/models"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Add(s string) model.CalculatorResponseModel {
	response := helper.IsStringEmpty(s)

	if response.Error != nil {
		return response
	}

	detectedStringArray := DetectNumbersInString(s)
	if detectedStringArray.Error != nil {
		return model.CalculatorResponseModel{
			Error: detectedStringArray.Error,
		}
	}
	return calculate(detectedStringArray.Result)
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

func DetectNumbersInString(s string) model.DelimiterResponseModel {
	reg, err := regexp.Compile("[^a-zA-Z0-9 -]")
	if err != nil {
		return model.DelimiterResponseModel{
			Error: errors.New("reger error occured"),
		}
	}

	regexOfString := reg.ReplaceAllString(s, " ")
	regexOfString = strings.Trim(regexOfString, " ")
	regexOfString = strings.Replace(regexOfString, " ", ",", -1)

	finalOfString := strings.Split(regexOfString, ",")

	return model.DelimiterResponseModel{
		Error:  nil,
		Result: finalOfString,
	}
}
