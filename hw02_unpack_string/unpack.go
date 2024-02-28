package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	runes := []rune(str)

	var res strings.Builder
	var shield bool
	var buf string

	for i := 0; i < len(runes); i++ {
		currentSymbol := runes[i]
		if currentSymbol == 92 && !shield {
			shield = true
			continue
		}

		if currentSymbol >= '0' && currentSymbol <= '9' {
			if !shield {
				return "", ErrInvalidString
			}
			buf = strconv.Itoa(int(currentSymbol - 48))
		} else {
			if shield && currentSymbol != 92 {
				return "", ErrInvalidString
			}
			buf = string(currentSymbol)
		}

		shield = false
		if i+1 < len(runes) {
			if repeatCount, err := strconv.Atoi(string(runes[i+1])); err == nil {
				if repeatCount > 0 {
					res.WriteString(strings.Repeat(buf, repeatCount))
				}
				i++
				continue
			}
		}

		res.WriteRune(currentSymbol)
	}

	return res.String(), nil
}
