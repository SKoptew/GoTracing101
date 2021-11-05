package rendering

import (
	. "gotracing101/math101"
	"math"
)

type Camera struct {
	clipNear float64
	clipFar  float64
	aspect   float64
	fov      float64

	//-- camera basis
	origin    Vec3
	direction Vec3
	horiz     Vec3
	vert      Vec3

	//-- UV vector basis
	uvBottomLeft Vec3
	uvHoriz      Vec3
	uvVert       Vec3
}

func NewCamera(aspect float64) *Camera {
	cam := Camera{
		clipNear:  0.01,
		clipFar:   math.MaxFloat64,
		aspect:    aspect,
	}
	cam.Set(Vec3{0,0,0,}, Vec3{0, 0, -1}, 60.0)
	return &cam
}

func (cam *Camera) ClipNear() float64 {
	return cam.clipNear
}

func (cam *Camera) ClipFar() float64 {
	return cam.clipFar
}

func (cam *Camera) Set(origin Vec3, direction Vec3, fov float64) {
	cam.fov    = fov
	cam.origin = origin

	//-- view-space basis
	cam.direction = direction.GetNormalized()
	cam.horiz     = Cross(direction, Vec3{0, 1, 0})
	cam.vert      = Cross(cam.horiz, direction)

	//-- UV basis
	focusPlaneHeight := 2.0 * math.Tan(ToRadians(fov) * 0.5)
	focusPlaneWidth  := focusPlaneHeight * cam.aspect

	cam.uvHoriz = MulC(cam.horiz, focusPlaneWidth)
	cam.uvVert  = MulC(cam.vert, focusPlaneHeight)
	cam.uvBottomLeft = Sub(Sub(cam.direction, MulC(cam.uvHoriz, 0.5)), MulC(cam.uvVert, 0.5))
}

func (cam *Camera) GetRay(u, v float64) *Ray {
	ray := Ray{
		Origin: cam.origin,
		Direction: Add(Add(cam.uvBottomLeft, MulC(cam.uvHoriz, u)), MulC(cam.uvVert, v)),
	}

	return &ray
}