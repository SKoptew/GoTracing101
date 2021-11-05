package scene

import (
	. "gotracing101/math101"

	"gotracing101/material"
)

type Hitable interface {
	Hit(ray *Ray, tMin float64, tMax float64) *material.HitRecord
}
