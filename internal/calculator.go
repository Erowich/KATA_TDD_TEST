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
	findNegativeNumbers := helper.DetectNegativeNumbers(detectedStringArray.Result)

	if findNegativeNumbers.Error != nil {

		return model.CalculatorResponseModel{
			Error: findNegativeNumbers.Error,
		}

	}
	return calculate(detectedStringArray.Result)
}

func calculate(sArray []string) model.CalculatorResponseModel {
	sum := 0
	for _, num := range sArray {
		number, err := strconv.Atoi(num)
		if err != nil {
			return model.CalculatorResponseModel{
				Error:  errors.New("can not parse string to int"),
				Result: 0,
			}
		}
		sum += number

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
			Error: errors.New("regex error occured"),
		}
	}

	regexOfString := reg.ReplaceAllString(s, " ")
	regexOfString = strings.Trim(regexOfString, " ")
	regexOfString = strings.Replace(regexOfString, " ", ",", -1)

	splitString := strings.Split(regexOfString, ",")

	finalOfString := delete_empty(splitString)

	return model.DelimiterResponseModel{
		Error:  nil,
		Result: finalOfString,
	}
}

func delete_empty(sArray []string) []string {
	var res []string
	for _, str := range sArray {
		if str != "" {
			res = append(res, str)
		}
	}
	return res
}
