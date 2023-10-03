package helpers

import (
	"unicode"
	"strconv"
	"strings"
	"time"
	"fmt"
	"math"
)

func IsTimeBetween2And4PM(inputTimeStr string) bool {
	parsedTime, err := time.Parse("15:04", inputTimeStr)
	if err != nil {
		return false
	}
	twoPM := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 14, 0, 0, 0, parsedTime.Location())
	fourPM := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 16, 0, 0, 0, parsedTime.Location())
	return parsedTime.After(twoPM) && parsedTime.Before(fourPM)
}

func IsRoundedDollarAmount(input string) bool {
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return false
	}
	return number == math.Trunc(number)
}

func IsMultipleOfQuarter(input string) bool {
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
			return false
	}
	numInt := int(num * 100) 
  return numInt%25 == 0
}

func CountAlphanumeric(input string) int {
	count := 0
	for _, char := range input {
			if unicode.IsLetter(char) || unicode.IsDigit(char) {
					count++
			}
	}
	return count
}

func GetDayFromDate(date string) int {
	dateComponents := strings.Split(date, "-")
	day, err := strconv.Atoi(dateComponents[2])
    if err != nil {
        panic(err)
    } 
	fmt.Println(day)
	return day
}