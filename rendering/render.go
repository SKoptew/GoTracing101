package rendering

import (
	. "gotracing101/math101"
	"gotracing101/scene"
	"image"
	"image/color"
	"math/rand"
)

func TraceRay(scene scene.Hitable, ray *scene.Ray, maxBounces int) Vec3 {
	accAttenuation := Vec3{1, 1, 1}
	ray.Direction.Normalize()

	for bounce := 0; bounce < maxBounces; bounce++ {
		if hit := scene.Hit(ray, 0.0001, 200); hit != nil {
			// ray always scattered again (no absorbtion/emission)
			// grey diffuse material, albedo = 0.5
			accAttenuation.MulC(0.5)

			ray.Origin    = hit.Pt
			ray.Direction = RandUnitVectorHemisphere(hit.Nrm)
		} else {
			break
		}
	}

	// ray reaches sky
	return MissShader(ray, accAttenuation)
}

func MissShader(ray *scene.Ray, accAttenuation Vec3) Vec3 {
	skyColor := Lerp(
		Vec3{1.0, 1.0, 1.0},
		Vec3{0.5, 0.7, 1.0},
		0.5*(1.0 + ray.Direction.Y))

	return Mul(accAttenuation, skyColor)
}

func convertColor(c Vec3) color.RGBA {
	return color.RGBA{
		R: uint8(c.X * 255.99),
		G: uint8(c.Y * 255.99),
		B: uint8(c.Z * 255.99),
		A: 0xff,
	}
}

func RenderImage(sc scene.Hitable, cam *Camera, width, height, spp, maxBounces int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	invWidth, invHeight := 1.0/float64(width), 1.0/float64(height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := Vec3{}

			for sampleNum := 0; sampleNum < spp; sampleNum++ {
				u :=       (float64(x) + rand.Float64()) * invWidth	// center value: (x + 0.5) / width; += -0.5...0.5 half-pixel scattering for multisampling
				v := 1.0 - (float64(y) + rand.Float64()) * invHeight

				ray := cam.GetRay(u, v)
				color.Add(TraceRay(sc, ray, maxBounces))
			}

			if spp > 1 {
				color.DivC(float64(spp))
			}

			color = LinearToSrgb(color)
			img.Set(x, y, convertColor(color))
		}
	}

	return img
}