package material

import (
	. "github.com/skoptew/gotracing101/math101"
	"math/rand"
)

type Emissive struct {
	color     Vec3
	intensity float64
}

func NewMatEmissive(color Vec3, intensity float64) Material {
	return &Emissive{
		color:     color,
		intensity: intensity,
	}
}

func (mat *Emissive) Scatter(*Ray, *HitRecord, *rand.Rand) (attenuation Vec3, rayOut *Ray) {
	return Vec3{}, nil
}

func (mat *Emissive) Emitted(*HitRecord) (emitted Vec3) {
	return MulC(mat.color, mat.intensity)
}
