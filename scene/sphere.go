package scene

import (
	"gotracing101/material"
	. "gotracing101/math101"
	"math"
)

type Sphere struct {
	center    Vec3
	radius    float64
	radiusInv float64
	material  material.Material
}

func NewSphere(center Vec3, radius float64, material material.Material) *Sphere {
	sphere := Sphere{
		center:    center,
		radius:    radius,
		radiusInv: 1.0 / radius,
		material:  material,
	}
	return &sphere
}


func (sphere *Sphere) Hit(ray *Ray, tMin float64, tMax float64) *material.HitRecord {
	oc    := Sub(ray.Origin, sphere.center)
	halfB := Dot(ray.Direction, oc)
	c     := Dot(oc, oc) - sphere.radius*sphere.radius
	discr := halfB*halfB - c				// 1/4 * discriminant

	intersect := func(t float64) *material.HitRecord {
		hitPoint := ray.GetPointAt(t)
		hit := material.HitRecord{
			T:   t,
			Pt:  hitPoint,
			Nrm: MulC(Sub(hitPoint, sphere.center), sphere.radiusInv),
			Mat: sphere.material,
		}
		return &hit
	}

	if discr > 0 {
		root := math.Sqrt(discr)

		// outer surface of sphere
		t := -halfB - root
		if t > tMin && t < tMax {
			return intersect(t)
		}

		// inner surface of sphere
		t = -halfB + root
		if t > tMin && t < tMax {
			return intersect(t)
		}
	}
	return nil
}
