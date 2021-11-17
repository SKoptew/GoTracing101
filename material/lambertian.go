package material

import (
	. "github.com/skoptew/gotracing101/math101"
	"math/rand"
)

type Lambertian struct {
	albedo Vec3
}

func NewMatLambertian(albedo Vec3) Material {
	return &Lambertian{albedo: albedo}
}

func (mat *Lambertian) Scatter(rayIn *Ray, hit *HitRecord, randSrc *rand.Rand) (attenuation Vec3, rayOut *Ray) {
	rayOut = &Ray{
		Origin:    hit.Pt,
		Direction: RandUnitVectorHemisphere(hit.Nrm, randSrc),
	}
	attenuation = mat.albedo
	return
}

func (mat *Lambertian) Emitted(*HitRecord) (emitted Vec3) {
	return Vec3{}
}
