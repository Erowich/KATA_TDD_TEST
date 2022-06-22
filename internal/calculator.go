package internal

import (
	helper "KATA_TDD/helper"
	model "KATA_TDD/models"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Add(s string, ch chan *model.CalculatorResponseModel) {
	response := helper.IsStringEmpty(s)
	chGetCalculateResponse := make(chan *model.CalculatorResponseModel)
	defer close(chGetCalculateResponse)
	if response.Error != nil {
		ch <- &response
		return
	}

	detectedStringArray := detectNumbersInString(s)
	if detectedStringArray.Error != nil {
		ch <- &response
		return
	}
	findNegativeNumbers := helper.DetectNegativeNumbers(detectedStringArray.Result)

	if findNegativeNumbers.Error != nil {

		ch <- &model.CalculatorResponseModel{Error: findNegativeNumbers.Error, Result: findNegativeNumbers.Result}
		return

	}
	go calculate(detectedStringArray.Result, chGetCalculateResponse)
	a := <-chGetCalculateResponse
	ch <- &model.CalculatorResponseModel{Error: a.Error, Result: a.Result}

}

func calculate(sArray []string, ch chan *model.CalculatorResponseModel) {
	chGetCalculateResponse := make(chan *model.CalculatorResponseModel)
	defer close(chGetCalculateResponse)
	sum := 0
	for _, num := range sArray {
		number, err := strconv.Atoi(num)
		if err != nil {
			ch <- &model.CalculatorResponseModel{
				Error:  errors.New("can not parse string to int"),
				Result: 0,
			}
		} else if number < 1000 {
			sum += number
		}

	}

	ch <- &model.CalculatorResponseModel{
		Error:  nil,
		Result: sum,
	}

}

func detectNumbersInString(s string) model.DelimiterResponseModel {
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
