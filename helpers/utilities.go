package helpers

import (
	"unicode"
	"strconv"
	"strings"
	"regexp"
	"time"
	"fmt"
)

func IsTimeBetween2And4PM(inputTimeStr string) bool {
	layout := "13:01"
	inputTime, err := time.Parse(layout, inputTimeStr)
	if err != nil {
			return false
	}
	startTime := time.Date(0, 0, 0, 14, 0, 0, 0, time.UTC)
	endTime := time.Date(0, 0, 0, 16, 0, 0, 0, time.UTC)
	return inputTime.After(startTime) && inputTime.Before(endTime)
}

func IsRoundedDollarAmount(input string) bool {
	pattern := `^\d+\.\d{2}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(input)
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