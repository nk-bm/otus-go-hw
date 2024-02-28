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
		currentRune := runes[i]
		if currentRune == 92 && !shield {
			if i+1 >= runesNum {
				return "", ErrInvalidString
			}
			shield = true
			continue
		}

		if currentRune >= '0' && currentRune <= '9' {
			if !shield {
				return "", ErrInvalidString
			}
			buf = strconv.Itoa(int(currentRune - 48))
		} else {
			if shield && currentRune != 92 {
				return "", ErrInvalidString
			}
			buf = string(currentRune)
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
		res.WriteRune(currentRune)
	}

	return res.String(), nil
}
