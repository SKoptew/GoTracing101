package rendering

import (
	. "gotracing101/math101"
	"gotracing101/scene"
	"image"
	"image/color"
)

func TraceRay(scene scene.Hitable, ray *scene.Ray) Vec3 {
	ray.Direction.Normalize()

	if hit := scene.Hit(ray, 0.0001, 200); hit != nil {
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

func RenderImage(sc scene.Hitable, cam *Camera, width int, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	invWidth, invHeight := 1.0/float64(width), 1.0/float64(height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			u :=       (float64(x) + 0.5) * invWidth
			v := 1.0 - (float64(y) + 0.5) * invHeight

			ray := cam.GetRay(u, v)
			img.Set(x, y, convertColor(TraceRay(sc, ray)))
		}
	}

	return img
}