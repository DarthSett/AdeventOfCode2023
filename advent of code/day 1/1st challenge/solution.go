package main

import (
	"strconv"
	"unicode"
)

func main() {
	strArr := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	sum := 0
	for _, v := range strArr {
		firstDigit, lastDigit := 0, 0
		firstDone := false
		for _, char := range v {
			if unicode.IsDigit(char) {
				if firstDone {
					lastDigit, _ = strconv.Atoi(string(char))
				} else {
					firstDigit, _ = strconv.Atoi(string(char))
					firstDone = true
				}
			}
		}
		if lastDigit == 0 {
			lastDigit = firstDigit
		}
		println((firstDigit * 10) + lastDigit)
		sum += (firstDigit * 10) + lastDigit
	}
	println(sum)
}

