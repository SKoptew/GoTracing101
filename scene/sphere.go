package scene

import (
	. "gotracing101/math101"
	"math"
)

type Sphere struct {
	center    Vec3
	radius    float64
	radiusInv float64
}

func NewSphere(center Vec3, radius float64) *Sphere {
	sphere := Sphere{
		center:    center,
		radius:    radius,
		radiusInv: 1.0 / radius,
	}
	return &sphere
}


func (sphere *Sphere) Hit(ray *Ray, tMin float64, tMax float64) *HitRecord {
	oc    := Sub(ray.Origin, sphere.center)
	halfB := Dot(ray.Direction, oc)
	c     := Dot(oc, oc) - sphere.radius*sphere.radius
	discr := halfB*halfB - c				// 1/4 * discriminant

	intersect := func(t float64) *HitRecord {
		hitPoint := ray.GetPointAt(t)
		hit := HitRecord{
			T:   t,
			Pt:  hitPoint,
			Nrm: MulC(Sub(hitPoint, sphere.center), sphere.radiusInv),
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
