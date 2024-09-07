package sqlgen

import (
	"fmt"
	"strconv"
	"strings"

	//"github.com/twpayne/cockroach/geo"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkbhex"
)

func CreatePointHexRepresentation(longitude, latitude float64) string {

	point := geom.NewPoint(geom.XY).MustSetCoords([]float64{longitude, latitude}).SetSRID(4326)

	hex, _ := ewkbhex.Encode(point, ewkbhex.NDR)
	return hex
}

func HexToPoint(s string) (float64, float64) {
	decoded, _ := ewkbhex.Decode(s)
	txt := fmt.Sprintf("%v", decoded)
	first := strings.Index(txt, "[")
	prefix := txt[first+1:]
	last := strings.Index(prefix, "]")
	if last == -1 {
		return 0, 0
	}
	tokens := strings.Split(prefix[0:last], " ")
	lonFloat, _ := strconv.ParseFloat(tokens[0], 64)
	latFloat, _ := strconv.ParseFloat(tokens[1], 64)
	return latFloat, lonFloat
}
