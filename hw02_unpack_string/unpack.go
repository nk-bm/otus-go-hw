package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	runes := []rune(str)
	runesNum := len(runes)

	var res strings.Builder
	var shield bool
	var buf string

	for i := 0; i < runesNum; i++ {
		currentSymbol := runes[i]
		if currentSymbol == 92 && !shield {
			if i+1 >= runesNum {
				return "", ErrInvalidString
			}
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

		if i+1 < runesNum {
			shield = false
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
