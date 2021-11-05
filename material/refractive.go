package material

import (
	. "gotracing101/math101"
	"math"
	"math/rand"
)

type Refractive struct {
	ior         float64
	attenuation Vec3
}

func NewMatRefractive(ior float64, attenuation Vec3) Material {
	return &Refractive {
		ior: ior,
		attenuation: attenuation,
	}
}

func (mat *Refractive) Scatter(rayIn *Ray, hit *HitRecord, randSrc *rand.Rand) (attenuation Vec3, rayOut *Ray) {
	VdotN   := Dot(rayIn.Direction, hit.Nrm)
	fresnel := getSchlickFresnelApprox(math.Abs(VdotN), mat.ior)

	rayOut = &Ray{ Origin: hit.Pt }

	if Rand01(randSrc) > fresnel {
		ior := mat.ior
		if VdotN < 0 {
			ior = 1 / ior
		}
		rayOut.Direction = Refract(rayIn.Direction, hit.Nrm, VdotN, ior)
	} else {
		rayOut.Direction = ReflectFast(rayIn.Direction, hit.Nrm, VdotN)
	}

	attenuation = mat.attenuation
	return
}

// calc probability of ray reflection
func getSchlickFresnelApprox(cosine, ior float64) float64 {
	r0 := (1 - ior) / (1 + ior)
	r0 *= r0
	return r0 + (1 - r0) * math.Pow(1 - cosine, 5)
}
