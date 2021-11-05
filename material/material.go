package material

import (
	. "gotracing101/math101"
	"math/rand"
)

type Material interface {
	Scatter(rayIn *Ray, hit *HitRecord, randSrc *rand.Rand) (attenuation Vec3, rayOut *Ray)
}