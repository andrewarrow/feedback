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

func randomColor() string {
	randInt := rand.Intn(8) + 1
	return fmt.Sprintf("bg-%s-%d00", Colors[rand.Intn(len(Colors))], randInt)
}
