package math101

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}
//-------------------------------------------------------------

func (v *Vec3) Negate() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vec3) Length2() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vec3) Normalize() {
	t := 1.0 / math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	v.X *= t
	v.Y *= t
	v.Z *= t
}

func (v *Vec3) GetNormalized() Vec3 {
	t := 1.0 / math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	return Vec3{
		X: v.X * t,
		Y: v.Y * t,
		Z: v.Z * t,
	}
}

//-------------------------------------------------------------
// binary functions, modify receiver

func (v *Vec3) Add(rhs Vec3) {
	v.X += rhs.X
	v.Y += rhs.Y
	v.Z += rhs.Z
}

func (v *Vec3) Sub(rhs Vec3) {
	v.X -= rhs.X
	v.Y -= rhs.Y
	v.Z -= rhs.Z
}

func (v *Vec3) Mul(rhs Vec3) {
	v.X *= rhs.X
	v.Y *= rhs.Y
	v.Z *= rhs.Z
}

func (v *Vec3) Div(rhs Vec3) {
	v.X /= rhs.X
	v.Y /= rhs.Y
	v.Z /= rhs.Z
}

func (v *Vec3) MulC(c float64) {
	v.X *= c
	v.Y *= c
	v.Z *= c
}

func (v *Vec3) DivC(c float64) {
	v.X /= c
	v.Y /= c
	v.Z /= c
}
//-------------------------------------------------------------
// binary functions, return new Vec3

func Add(a Vec3, b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func Sub(a Vec3, b Vec3) Vec3 {
	return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func Mul(a Vec3, b Vec3) Vec3 {
	return Vec3{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

func Div(a Vec3, b Vec3) Vec3 {
	return Vec3{a.X / b.X, a.Y / b.Y, a.Z / b.Z}
}

func MulC(a Vec3, c float64) Vec3 {
	return Vec3{a.X * c, a.Y * c, a.Z * c}
}

func DivC(a Vec3, c float64) Vec3 {
	return Vec3{a.X / c, a.Y / c, a.Z / c}
}

func Dot(a Vec3, b Vec3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Cross(a Vec3, b Vec3) Vec3 {
	return Vec3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func Lerp(a Vec3, b Vec3, t float64) Vec3 {
	return Add(MulC(a, 1.0 - t), MulC(b, t))
}
//-------------------------------------------------------------

func (v Vec3) String() string {
	return fmt.Sprintf("[%v, %v, %v]", v.X, v.Y, v.Z)
}

func ToRadians(angle float64) float64 {
	return angle * math.Pi / 180.0
}

func ToDegrees(angle float64) float64 {
	return angle * 180.0 / math.Pi
}