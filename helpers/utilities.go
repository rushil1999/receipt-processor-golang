package helpers

import (
	"unicode"
	"strconv"
	"strings"
	"time"
	"math"
	"receipt-processor-module/models"
)

func IsTimeBetween2And4PM(inputTimeStr string) (bool, error) {
	parsedTime, err:= time.Parse("15:04", inputTimeStr)
	if err != nil {
		customError := models.CustomError {
			Message: "Invalid Input",
			DebugMessage: "Cannot parse float",
			HttpCode: 404,
		}
		return false, customError
	}
	twoPM := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 14, 0, 0, 0, parsedTime.Location())
	fourPM := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 16, 0, 0, 0, parsedTime.Location())
	return parsedTime.After(twoPM) && parsedTime.Before(fourPM), nil
}

func IsRoundedDollarAmount(input string) (bool, error) {
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		customError := models.CustomError {
			Message: "Invalid Input",
			DebugMessage: "Cannot parse float",
			HttpCode: 404,
		}
		return false, customError
	}
	return number == math.Trunc(number), nil // Checking if the float number is a whole number
}

func IsMultipleOfQuarter(input string) (bool, error) {
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		customError := models.CustomError {
			Message: "Invalid Input",
			DebugMessage: "Cannot parse float",
			HttpCode: 404,
		}
		return false, customError
	}
	numInt := int(num * 100) 
  return numInt%25 == 0, nil
}

func CountAlphanumeric(input string) int {
	count := 0
	for _, char := range input {
			if unicode.IsLetter(char) || unicode.IsDigit(char) { // Checking if the character is alphnumeric
					count++
			}
	}
	return count
}

func GetDayFromDate(date string) (int, error)  {
	dateComponents := strings.Split(date, "-") // Splitting date components
	if len(dateComponents) != 3{
		customError := models.CustomError {
			Message: "Invalid Input",
			DebugMessage: "Cannot parse float",
			HttpCode: 404,
		}
		return -1, customError
	}
	day, err := strconv.Atoi(dateComponents[2])
	if err != nil {
		customError := models.CustomError {
			Message: "Invalid Input",
			DebugMessage: "Cannot parse float",
			HttpCode: 404,
		}
		return -1, customError
		
	} 
	return day, nil
}
