package geo_ext

import (
	"math"
)

type Position struct {
	Lon float64 // 经度
	Lat float64 // 维度
}

func Distence(p, q Position) float64 {
	var α1, β1, α2, β2 = rad(p.Lon), rad(p.Lat), rad(q.Lon), rad(q.Lat)

	// L = R*arccos(cos(β1)cos(β2)cos(α1-α2)+sin(β1)sin(β2)) , (
	const r = 6378137
	return r * math.Acos(math.Cos(β1)*math.Cos(β2)*math.Cos(α1-α2)+math.Sin(β1)*math.Sin(β2))
}

func rad(d float64) float64 {
	return d * math.Pi / 180.0
}
