package util

import (
	"crypto/rand"
	"fmt"
)

func PseudoUuid() string {

	b := make([]byte, 16)
	rand.Read(b)

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
