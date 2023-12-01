package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var onlyNumbersRegex = regexp.MustCompile(`[a-zA-Z]+`)
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	stringContents := string(contents[:])
	test := strings.Split(stringContents, "\n")
	var total int = 0
	for _, calibrationValue := range test {
		onlyNumbersString := onlyNumbersRegex.ReplaceAllString(calibrationValue, "")
		calibrationNumber, bad := strconv.Atoi(onlyNumbersString)
		if bad != nil {
			fmt.Println("Can't convert string to num", bad)
			return
		}
		if len(onlyNumbersString) == 1 {
			singleStringNumCombined := fmt.Sprintf("%d%d", calibrationNumber, calibrationNumber)
			numToAdd, numError := strconv.Atoi(singleStringNumCombined)
			if numError != nil {
				fmt.Println("Failure to convert string to num", numError)
				return
			}
			total += numToAdd
		} else {
			firstAndLast := fmt.Sprintf("%s%s", onlyNumbersString[0:1], onlyNumbersString[len(onlyNumbersString)-1:])
			numToAdd, numError := strconv.Atoi(firstAndLast)
			if numError != nil {
				fmt.Println("Failure to convert string to num", numError)
				return
			}
			total += numToAdd
		}
	}
	fmt.Println("Done!")
	fmt.Println(total)
	fmt.Println("Done!")
}
