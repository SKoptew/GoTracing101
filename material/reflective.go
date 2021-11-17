package material

import (
	. "github.com/skoptew/gotracing101/math101"
	"math/rand"
)

type Reflective struct {
	albedo    Vec3
	fuzziness float64
}

func NewMatReflective(albedo Vec3, fuzziness float64) Material {
	return &Reflective{
		albedo:    albedo,
		fuzziness: fuzziness,
	}
}

func (mat *Reflective) Scatter(rayIn *Ray, hit *HitRecord, randSrc *rand.Rand) (attenuation Vec3, rayOut *Ray) {
	outDir := Reflect(rayIn.Direction, hit.Nrm)

	if mat.fuzziness > 0 {
		outDir.Add(MulC(RandUnitVectorHemisphere(hit.Nrm, randSrc), mat.fuzziness))
		outDir.Normalize()
	}

	rayOut = &Ray{
		Origin:    hit.Pt,
		Direction: outDir,
	}
	attenuation = mat.albedo
	return
}

func (mat *Reflective) Emitted(*HitRecord) (emitted Vec3) {
	return Vec3{}
}
