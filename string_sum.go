package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	runes := []rune(strings.ReplaceAll(input, " ", ""))

	if len(runes) == 0 {
		return "", errorEmptyInput
	}

	operands, ok := CheckTwoOperands(string(runes))

	if !ok {
		return "", errorNotTwoOperands
	}

	sum := 0
	for _, r := range operands {
		num, err := strconv.Atoi(r)
		if err != nil {
			return "", err
		}
		sum += num
	}

	return strconv.Itoa(sum), nil
}

func IsRuneSign(r rune) bool {
	return r == '-' || r == '+'
}

func CheckTwoOperands(input string) ([]string, bool) {
	runes := []rune(input)
	start := 0
	var twoOperands []string
	str := ""

	if IsRuneSign(runes[0]) {
		start = 1
		if runes[0] == '-' {
			str = "-"
		}
	}
	// fmt.Println(string(runes))
	for i := start; i < len(runes); i++ {
		if IsRuneSign(runes[i]) {
			// fmt.Println(str)
			twoOperands = append(twoOperands, str)
			str = string(runes[i])

			continue
		} else {
			str += string(runes[i])
			// fmt.Println("dd", str)
		}
	}
	twoOperands = append(twoOperands, str)
	fmt.Println(twoOperands)
	if len(twoOperands) != 2 {
		return nil, false
	}

	return twoOperands, true
}
