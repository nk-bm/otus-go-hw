package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	slen := len(str)
	var res string

	var shield bool
	var buf string

	for i := 0; i < slen; i++ {
		currentSymbol := str[i]
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
				buf = "\\" + string(currentSymbol)
			} else {
				buf = string(currentSymbol)
			}
		}

		shield = false
		if i+1 < slen {
			if repeatCount, err := strconv.Atoi(string(str[i+1])); err == nil {
				if repeatCount > 0 {
					res += strings.Repeat(buf, repeatCount)
				}
				i++
				continue
			}
		}

		res += string(currentSymbol)
	}
	return res, nil
}
