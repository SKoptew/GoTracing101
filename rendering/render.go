package rendering

import (
	. "gotracing101/math101"
	"gotracing101/scene"
	"image"
	"image/color"
)

func TraceRay(scene scene.Hitable, ray *scene.Ray) Vec3 {
	ray.Direction.Normalize()

	if hit, isHit := scene.Hit(ray, 0.0001, 200); isHit {
		return Add(MulC(hit.Nrm, 0.5), Vec3{0.5, 0.5, 0.5})
	}

	// no hit, return sky imitation
	return Lerp(
		Vec3{0.05, 0.05, 0.2},
		Vec3{0.7, 0.8, 0.9},
		0.5*(1.0 + ray.Direction.Y))
}

func convertColor(c Vec3) color.RGBA {
	return color.RGBA{
		R: uint8(c.X * 255.99),
		G: uint8(c.Y * 255.99),
		B: uint8(c.Z * 255.99),
		A: 0xff,
	}
}

func RenderImage(width int, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	invWidth, invHeight := 1.0/float64(width), 1.0/float64(height)
	aspect := float64(height) * invWidth

	//-- test simple scene
	testSphere := scene.Sphere{Center: Vec3{0, 0, -5}, Radius: 1.0}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			u := (float64(x) + 0.5) * invWidth
			v := (float64(y) + 0.5) * invHeight

			ray := scene.Ray {
				Origin: Vec3{},
				Direction: Vec3{
					X: 2.0*u - 1.0,
					Y: (2.0*v - 1.0) * aspect,
					Z: -2.0,
				},
			}

			col := TraceRay(&testSphere, &ray)
			img.Set(x, y, convertColor(col))
		}
	}

	return img
}