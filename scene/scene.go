package scene

import (
	"github.com/skoptew/gotracing101/material"
	. "github.com/skoptew/gotracing101/math101"
	"math"
)

type Scene struct {
	surfaces []Hitable
}

func NewScene() *Scene {
	scene := Scene{}
	return &scene
}

func (scene *Scene) Hit(ray *Ray, tMin float64, tMax float64) *material.HitRecord {

	var closestHit *material.HitRecord
	closestHitT := math.MaxFloat64

	for _, surface := range scene.surfaces {
		if hit := surface.Hit(ray, tMin, tMax); hit != nil {
			if hit.T < closestHitT {
				closestHit = hit
				closestHitT = hit.T
			}
		}
	}

	return closestHit
}

func (scene *Scene) Add(surface Hitable) {
	scene.surfaces = append(scene.surfaces, surface)
}
