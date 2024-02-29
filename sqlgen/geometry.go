package sqlgen

import (
	"fmt"

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
	fmt.Println(decoded)
	//point, ok := decoded.(*geom.Point)
	//coords := point.Coords()
	//return coords[0], coords[1]
	return 0, 0
}
