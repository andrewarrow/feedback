package markup

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func DivsAndDivs() {
	rand.Seed(time.Now().UnixNano())

	count := 0
	fmt.Println("div")
	spaces := "  "
	maxIndent := 4
	for {
		count++
		childIndent := 2
		childSpaces := strings.Repeat(" ", childIndent)
		fmt.Printf("%sdiv\n", spaces+childSpaces)

		// Randomly decide whether to increase, decrease, or stay the same
		action := rand.Intn(3) - 1
		maxIndent += 2 * action

		spaces = moreOrLess(len(spaces), maxIndent)
		if count > 20 {
			break
		}
	}
}

func moreOrLess(currLen, maxLen int) string {
	if currLen < maxLen {
		return strings.Repeat(" ", currLen+2)
	} else {
		if currLen-2 < 0 {
			return "  "
		}
		return strings.Repeat(" ", currLen-2)
	}
}
