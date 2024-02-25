package sqlgen

import (
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkbhex"
)

func CreatePointHexRepresentation(longitude, latitude float64) string {

	point := geom.NewPoint(geom.XY).MustSetCoords([]float64{longitude, latitude}).SetSRID(4326)

	hex, _ := ewkbhex.Encode(point, ewkbhex.NDR)
	return hex
}
