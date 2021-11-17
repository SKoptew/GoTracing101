package material

import (
	. "github.com/skoptew/gotracing101/math101"
	"math/rand"
)

type Material interface {
	Scatter(rayIn *Ray, hit *HitRecord, randSrc *rand.Rand) (attenuation Vec3, rayOut *Ray)
	Emitted(hit *HitRecord) (emitted Vec3)
}
