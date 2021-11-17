package material

import (
	. "github.com/skoptew/gotracing101/math101"
)

type HitRecord struct {
	T   float64
	Pt  Vec3
	Nrm Vec3
	Mat Material
}
