package helpers

import (
	"unicode"
	"strconv"
	"strings"
	"time"
	"math"
	"fmt"
)

func IsTimeBetween2And4PM(inputTimeStr string) (bool, string) {
	parsedTime, err:= time.Parse("15:04", inputTimeStr)
	errorMsg := "Invalid Purchase Time"
	if err != nil {
		return false, errorMsg
	}
	twoPM := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 14, 0, 0, 0, parsedTime.Location())
	fourPM := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 16, 0, 0, 0, parsedTime.Location())
	return parsedTime.After(twoPM) && parsedTime.Before(fourPM), ""
}

func IsRoundedDollarAmount(input string) (bool, string) {
	number, err := strconv.ParseFloat(input, 64)
	errorMsg := "Invalid Total Amount"
	if err != nil {
		return false, errorMsg
	}
	return number == math.Trunc(number), "" // Checking if the float number is a whole number
}

func IsMultipleOfQuarter(input string) (bool, string) {
	num, err := strconv.ParseFloat(input, 64)
	errorMsg := "Invalid Total Amount"
	if err != nil {
		return false, errorMsg
	}
	numInt := int(num * 100) 
  return numInt%25 == 0, ""
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

func GetDayFromDate(date string) (int, string)  {
	dateComponents := strings.Split(date, "-") // Splitting date components
	errorMsg := "Invalid Total Amount"
	if len(dateComponents) != 3{
		return -1, errorMsg
	}
	day, err := strconv.Atoi(dateComponents[2])
	if err != nil {
		fmt.Println("Caught")
		return -1, errorMsg
	} 
	return day, ""
}