// Package parser
// @description
// @author      张盛钢
// @datetime    2023/4/10 13:48
package parser

import (
	"os"
	"strings"
)

func isSafeRune(r rune) bool {
	return isLetter(r) || isNumber(r) || r == '_'
}

func isLetter(r rune) bool {
	return 'A' <= r && r <= 'z'
}

func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}

// SafeString converts the input string into a safe naming style in golang
func SafeString(in string) string {
	if len(in) == 0 {
		return in
	}

	data := strings.Map(func(r rune) rune {
		if isSafeRune(r) {
			return r
		}
		return '_'
	}, in)

	headRune := rune(data[0])
	if isNumber(headRune) {
		return "_" + data
	}
	return data
}

// MkdirIfNotExist makes directories if the input path is not exists
func MkdirIfNotExist(dir string) error {
	if len(dir) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}
