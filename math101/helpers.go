package math101

import (
	"math"
	"math/rand"
)

func ToRadians(angle float64) float64 {
	return angle * math.Pi / 180.0
}

func ToDegrees(angle float64) float64 {
	return angle * 180.0 / math.Pi
}

func RandUnitVectorSphere() Vec3 {
	//-- Marsaglia, George. Choosing a Point from the Surface of a Sphere. Ann. Math. Statist. 43 (1972), no. 2, 645--646
	for {
		p := Vec3{
			X: rand.Float64() * 2.0 - 1.0, // -1..1
			Y: rand.Float64() * 2.0 - 1.0,
			Z: rand.Float64() * 2.0 - 1.0,
		}

		if length2 := p.Length2(); length2 < 1.0 {
			return DivC(p, math.Sqrt(length2))
		}
	}
}

func RandUnitVectorHemisphere(normal Vec3) Vec3 {
	v := RandUnitVectorSphere()
	if Dot(v, normal) < 0.0 {
		v.Negate()
	}
	return v
}

const gamma    float64 = 2.2
const gammaInv float64 = 1.0 / gamma

func LinearToSrgb(c Vec3) Vec3 {
	c.X = math.Pow(c.X, gammaInv)
	c.Y = math.Pow(c.Y, gammaInv)
	c.Z = math.Pow(c.Z, gammaInv)
	return c
}

func SrgbToLinear(c Vec3) Vec3 {
	c.X = math.Pow(c.X, gamma)
	c.Y = math.Pow(c.Y, gamma)
	c.Z = math.Pow(c.Z, gamma)
	return c
}