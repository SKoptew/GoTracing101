package math101

import (
	"math"
	"math/rand"
)

func Sign(x float64) float64 {
	if math.Signbit(x) {
		return -1.0
	}
	return 1.0
}

func Rand01(randSrc *rand.Rand) float64 {
	return randSrc.Float64()
}

func Rand(randSrc *rand.Rand, min, max float64) float64 {
	if max < min {
		min, max = max, min
	}
	return min + (max-min)*randSrc.Float64()
}

func ToRadians(angle float64) float64 {
	return angle * math.Pi / 180.0
}

// Reflect : reflected vector. v - n * 2*dot(v, n);
func Reflect(v Vec3, n Vec3) Vec3 {
	return Sub(v, MulC(n, 2*Dot(v, n)))
}

// ReflectFast : reflected vector. Version with pre-calculated VdotN value
func ReflectFast(v Vec3, n Vec3, VdotN float64) Vec3 {
	return Sub(v, MulC(n, 2*VdotN))
}

// Refract : refracted vector. normal directed from second medium towards first medium
func Refract(v Vec3, n Vec3, VdotN, ior float64) Vec3 {
	vTan := MulC(Sub(v, MulC(n, VdotN)), ior)

	if discr := 1 - vTan.Length2(); discr > 0 {
		vNorm := MulC(n, math.Sqrt(discr))
		if math.Signbit(VdotN) {
			vNorm.Negate()
		}

		return Add(vTan, vNorm)
	}

	// total internal reflection
	return Sub(v, MulC(n, 2*VdotN))
}

func RandUnitVectorSphere(randSrc *rand.Rand) Vec3 {
	//-- Marsaglia, George. Choosing a Point from the Surface of a Sphere. Ann. Math. Statist. 43 (1972), no. 2, 645--646
	for {
		p := Vec3{
			X: Rand01(randSrc)*2.0 - 1.0, // -1..1
			Y: Rand01(randSrc)*2.0 - 1.0,
			Z: Rand01(randSrc)*2.0 - 1.0,
		}

		if length2 := p.Length2(); length2 < 1.0 {
			return DivC(p, math.Sqrt(length2))
		}
	}
}

func RandUnitVectorHemisphere(normal Vec3, randSrc *rand.Rand) Vec3 {
	v := RandUnitVectorSphere(randSrc)
	if Dot(v, normal) < 0.0 {
		v.Negate()
	}
	return v
}

const gamma float64 = 2.2
const gammaInv float64 = 1.0 / gamma

func LinearToSrgb(c Vec3) Vec3 {
	c.X = math.Pow(c.X, gammaInv)
	c.Y = math.Pow(c.Y, gammaInv)
	c.Z = math.Pow(c.Z, gammaInv)
	return c
}
