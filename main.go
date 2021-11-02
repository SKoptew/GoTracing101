package main

import (
	. "gotracing101/math101"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Ray struct {
	origin    Vec3
	direction Vec3
}

func TraceRay(ray *Ray) Vec3 {
	return Lerp(
		Vec3{0.05, 0.05, 0.2},
		Vec3{0.7, 0.8, 0.9},
		0.5*(1.0 + ray.direction.Y))
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

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			u := (float64(x) + 0.5) * invWidth
			v := (float64(y) + 0.5) * invHeight

			ray := Ray {
				origin: Vec3{},
				direction: Vec3{
					X: 2.0*u - 1.0,
					Y: (2.0*v - 1.0) * aspect,
					Z: -2.0,
				},
			}

			col := TraceRay(&ray)
			img.Set(x, y, convertColor(col))
		}
	}

	return img
}

func main() {
	img := RenderImage(512, 256)

	file, _ := os.Create("out.png")
	png.Encode(file, img)
}