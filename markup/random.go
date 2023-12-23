package markup

import (
	"fmt"
	"math/rand"
	"strings"
)

func DivsAndDivs() {
	count := 0
	fmt.Println("div")
	spaces := "  "
	vals := []int{2, 4}
	max := 4
	for {
		fmt.Printf("%sdiv\n", spaces)
		count++
		spaces = moreOrLess(len(spaces), vals)
		if len(spaces) == max {
			max = max + 2
			vals = append(vals, max)
		}
		if count > 10 {
			break
		}
	}
}

func moreOrLess(size int, vals []int) string {
	val := rand.Intn(len(vals))
	n := vals[val]
	buff := []string{}
	for i := 0; i < n; i++ {
		buff = append(buff, " ")
	}
	return strings.Join(buff, "")
}

func moreOrLess2(size int) string {
	if size == 2 {
		return "    "
	} else if size == 4 {
		return "      "
	} else if size == 6 {
		return "        "
	}
	return "  "
}
