package scene

import (
	"gotracing101/material"
	. "gotracing101/math101"
)

type Hitable interface {
	Hit(ray *Ray, tMin float64, tMax float64) *material.HitRecord
}
