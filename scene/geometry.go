package scene

import (
	. "gotracing101/math101"
	"math"
)

type Sphere struct {
	Center Vec3
	Radius float64
}


func (sphere *Sphere) Hit(ray *Ray, tMin float64, tMax float64) (*HitRecord, bool) {
	oc    := Sub(ray.Origin, sphere.Center)
	halfB := Dot(ray.Direction, oc)
	c     := Dot(oc, oc) - sphere.Radius*sphere.Radius
	discr := halfB*halfB - c				// 1/4 * discriminant

	intersect := func(t float64) (*HitRecord, bool) {
		hitPoint := ray.GetPointAt(t)
		hit := HitRecord{
			T:   t,
			Pt:  hitPoint,
			Nrm: DivC(Sub(hitPoint, sphere.Center), sphere.Radius), // $$$ replace to MulC(invR)
		}
		return &hit, true
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
	return nil, false
}
