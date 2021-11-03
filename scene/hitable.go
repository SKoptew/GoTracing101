package scene

import (
	. "gotracing101/math101"
)

type HitRecord struct {
	T   float64
	Pt  Vec3
	Nrm Vec3
}

type Hitable interface {
	Hit(ray *Ray, tMin float64, tMax float64) (*HitRecord, bool)
}

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (ray *Ray) GetPointAt(t float64) Vec3 {
	return Add(ray.Origin, MulC(ray.Direction, t))
}