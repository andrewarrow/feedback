package sqlgen

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

func CreatePointHexRepresentation(longitude, latitude float64) string {
	lonHex, _ := float64ToHex(longitude)
	latHex, _ := float64ToHex(latitude)
	pointHex := lonHex + latHex
	return pointHex
}

func float64ToHex(value float64) (string, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf.Bytes()), nil
}
