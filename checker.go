package main

import (
	"strconv"
	"strings"
)

var operators = []byte{'+', '-', '*', '/'}

func checkString(s string) (n1, n2 int, operator byte, isRoman bool, err error) {
	// Check operator
	var j int
	for _, v := range operators {
		if j = strings.IndexByte(s, v); j > 0 {
			break
		}
	}
	if j < 1 {
		return 0, 0, 0, false, errWrongInput
	}

	// Check is arabic numbers
	n1, err1 := strconv.Atoi(s[:j])
	n2, err2 := strconv.Atoi(s[j+1:])
	if err1 == nil && err2 == nil {
		return n1, n2, s[j], false, nil
	}

	// Check is roman numbers
	n1, err1 = romanToInt(s[:j])
	n2, err2 = romanToInt(s[j+1:])
	if err1 == nil && err2 == nil {
		return n1, n2, s[j], true, nil
	}

	return 0, 0, 0, false, errWrongInput
}

func checkLimits(n1, n2 int, operator byte, isRoman bool) error {
	if isRoman {
		if (operator == '-' && n2 >= n1) || (operator == '/' && n2 > n1) {
			return errWrongInput
		}
	} else {
		if n1 < 1 || n1 > 10 || n2 < 1 || n2 > 10 {
			return errWrongInput
		}
	}
	return nil
}
