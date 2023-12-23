package markup

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var attrs = []string{"flex", "w-%d", "h-%d", "items-center", "justify-center"}

func makeAttrs() string {
	buff := []string{}
	r := rand.Intn(100)
	if r > 50 {
		buff = append(buff, "flex")
	}
	r = rand.Intn(100)
	if r > 50 {
		buff = append(buff, "bg-r")
	}
	r = rand.Intn(100)
	if r > 50 {
		buff = append(buff, "w-64")
		buff = append(buff, "h-64")
	}
	r = rand.Intn(100)
	if r > 50 {
		buff = append(buff, "w-96")
		buff = append(buff, "h-96")
	}
	r = rand.Intn(100)
	if r > 50 {
		buff = append(buff, "w-full")
	}
	r = rand.Intn(100)
	if r > 50 {
		buff = append(buff, "w-1/2")
	}
	r = rand.Intn(100)
	if r > 50 {
		buff = append(buff, "rounded")
	}
	r = rand.Intn(100)
	if r > 50 {
		buff = append(buff, "rounded-full")
	}
	r = rand.Intn(100)
	if r > 50 {
		buff = append(buff, "flex-grow")
	}

	return strings.Join(buff, " ")
}

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
		fmt.Printf("%sdiv %s\n", spaces+childSpaces, makeAttrs())

		r := rand.Intn(100)
		action := 1
		if r > 60 {
			action = -1
		}
		r = rand.Intn(100)
		if r > 60 {
			action = 0
		}
		if maxIndent > 20 {
			action = rand.Intn(3) - 1
		}
		maxIndent += 2 * action

		spaces = moreOrLess(len(spaces), maxIndent)
		if count > 200 {
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
