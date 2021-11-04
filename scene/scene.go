package scene

import "math"

type Scene struct {
	surfaces []Hitable
}

func NewScene() *Scene {
	scene := Scene{}
	return &scene
}

func (scene *Scene) Hit(ray *Ray, tMin float64, tMax float64) *HitRecord {

	var closestHit *HitRecord
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
