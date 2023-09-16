package util

import "fmt"

func IntComma(i int) string {
	if i < 0 {
		return "-" + IntComma(-i)
	}
	if i < 1000 {
		return fmt.Sprintf("%d", i)
	}
	return IntComma(i/1000) + "," + fmt.Sprintf("%03d", i%1000)
}
