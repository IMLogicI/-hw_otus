package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	runesArr := []rune(s)
	if len(runesArr) > 0 && unicode.IsNumber(runesArr[0]) {
		return "", ErrInvalidString
	}

	lenS := len(s) - 1
	res := strings.Builder{}
	slashRune := []rune(`\`)[0]
	currRuneIsScreened := false
	currRuneIsIndex := false

	for i, currRune := range runesArr {
		runeStr := string(currRune)

		switch {
		case currRune == slashRune && !currRuneIsScreened:
			if i == lenS || !(unicode.IsNumber(runesArr[i+1]) || runesArr[i+1] == slashRune) {
				return "", ErrInvalidString
			}

			currRuneIsScreened = true
			continue
		case lenS > i && unicode.IsNumber(runesArr[i+1]):
			if lenS > i+1 && unicode.IsNumber(runesArr[i+2]) {
				return "", ErrInvalidString
			}

			num, err := strconv.Atoi(string(runesArr[i+1]))
			if err != nil {
				return "", fmt.Errorf("err when converting num str to int: %w", err)
			}

			res.WriteString(strings.Repeat(runeStr, num))
			currRuneIsIndex = true
			continue
		case !currRuneIsIndex:
			res.WriteRune(currRune)
		}

		currRuneIsScreened = false
		currRuneIsIndex = false
	}

	return res.String(), nil
}
