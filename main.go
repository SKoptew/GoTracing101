package main

import (
	"fmt"
	. "gotracing101/math101"
)

func main() {
	v0 := Vec3{1, 2, 3}
	v1 := Vec3{7, 7, 7}


	fmt.Println(Add(v0, v1))
	fmt.Println(Mul(v0, v1))
	fmt.Println(Sub(v0, v1))
	fmt.Println(Div(v0, v1))

	fmt.Println(Cross(Vec3{1,0,0}, Vec3{0,1,0}))
	fmt.Println(Dot(Vec3{1,2,3}, Vec3{1,2,3}))

	fmt.Println(v0.Length())
	fmt.Println((&Vec3{1,0,0}).Length())

	v0.Normalize()
	fmt.Println(v0)

	fmt.Println(v1.GetNormalized())
}