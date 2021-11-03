package scene

type Scene struct {
	objects []Hitable
}

func (scene *Scene) Hit(ray *Ray, tMin float64, tMax float64) (*HitRecord, bool) {
	return nil, false
}
