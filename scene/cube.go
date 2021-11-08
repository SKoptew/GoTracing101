package scene

import (
	"gotracing101/material"
	. "gotracing101/math101"
	"math"
)

type Cube struct {
	center   Vec3
	pmin     Vec3
	pmax     Vec3
	material material.Material
}

func NewCube(center Vec3, side float64, material material.Material) Hitable {
	d := side

	return &Cube{
		center: center,
		pmin: Sub(center, Vec3{d, d, d}),
		pmax: Add(center, Vec3{d, d, d}),
		material: material,
	}
}

func (cube *Cube) Hit(ray *Ray, tMin float64, tMax float64) *material.HitRecord {
	rayInvDir := Vec3{1/ray.Direction.X, 1/ray.Direction.Y, 1/ray.Direction.Z}

	// intersect ray with YZ planes
	t0 := (cube.pmin.X - ray.Origin.X) * rayInvDir.X
	t1 := (cube.pmax.X - ray.Origin.X) * rayInvDir.X
	if ray.Direction.X < 0 {
		t0, t1 = t1, t0
	}

	tMin = math.Max(tMin, t0)
	tMax = math.Min(tMax, t1)

	if tMin >= tMax {
		return nil
	}

	// intersect ray with XZ planes
	t0 = (cube.pmin.Y - ray.Origin.Y) * rayInvDir.Y
	t1 = (cube.pmax.Y - ray.Origin.Y) * rayInvDir.Y
	if ray.Direction.Y < 0 {
		t0, t1 = t1, t0
	}

	tMin = math.Max(tMin, t0)
	tMax = math.Min(tMax, t1)

	if tMin >= tMax {
		return nil
	}
	
	// intersect with XY planes
	t0 = (cube.pmin.Z - ray.Origin.Z) * rayInvDir.Z
	t1 = (cube.pmax.Z - ray.Origin.Z) * rayInvDir.Z
	if ray.Direction.Z < 0 {
		t0, t1 = t1, t0
	}

	tMin = math.Max(tMin, t0)
	tMax = math.Min(tMax, t1)

	if tMin >= tMax {
		return nil
	}

	hitPoint := ray.GetPointAt(tMin)
	hitNrm := Vec3{}
	nrm := Sub(hitPoint, cube.center)
	nrm.Normalize()
	nx, ny, nz := math.Abs(nrm.X), math.Abs(nrm.Y), math.Abs(nrm.Z)

	switch {
	case nx > ny && nx > nz:
		hitNrm = Vec3{Sign(nrm.X), 0, 0}
	case ny > nx && ny > nz:
		hitNrm = Vec3{0, Sign(nrm.Y), 0}
	default:
		hitNrm = Vec3{0, 0, Sign(nrm.Z)}
	}
	
	return &material.HitRecord{
		T:   tMin,
		Pt:  hitPoint,
		Nrm: hitNrm,
		Mat: cube.material,
	}
}