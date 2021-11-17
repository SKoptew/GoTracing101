package scene

import (
	"github.com/skoptew/gotracing101/material"
	. "github.com/skoptew/gotracing101/math101"
)

type Hitable interface {
	Hit(ray *Ray, tMin float64, tMax float64) *material.HitRecord
}
