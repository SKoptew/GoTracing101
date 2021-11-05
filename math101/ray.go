package math101

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (ray *Ray) GetPointAt(t float64) Vec3 {
	return Add(ray.Origin, MulC(ray.Direction, t))
}
