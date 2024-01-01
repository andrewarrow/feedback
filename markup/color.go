package markup

import (
	"fmt"
	"math/rand"
)

var Colors = []string{"gray",
	"red",
	"yellow",
	"green",
	"blue",
	"indigo",
	"purple",
	"pink",
	"rose",
	"teal",
}

func RandomColor() string {
	randInt := rand.Intn(8) + 1
	text := "gray-900"
	if randInt > 5 {
		text = "white"
	}
	return fmt.Sprintf("bg-%s-%d00 text-%s",
		Colors[rand.Intn(len(Colors))], randInt, text)
}
