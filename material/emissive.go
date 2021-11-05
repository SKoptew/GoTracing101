package material

import (
	. "gotracing101/math101"
	"math/rand"
)

type Emissive struct {
	color     Vec3
	intensity float64
}

func NewMatEmissive(color Vec3, intensity float64) Material {
	return &Emissive {
		color:     color,
		intensity: intensity,
	}
}

func (mat *Emissive) Scatter(rayIn *Ray, hit *HitRecord, randSrc *rand.Rand) (attenuation Vec3, rayOut *Ray) {
	return Vec3{}, nil
}

func (mat *Emissive) Emitted(hit *HitRecord) (emitted Vec3) {
	return MulC(mat.color, mat.intensity)
}