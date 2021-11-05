package rendering

import (
	. "gotracing101/math101"
	"gotracing101/scene"
	"image"
	"image/color"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func TraceRay(scene scene.Hitable, ray *Ray, maxBounces int, randSrc *rand.Rand) Vec3 {
	accAttenuation := Vec3{1, 1, 1}
	ray.Direction.Normalize()

	for bounce := 0; bounce < maxBounces; bounce++ {
		if hit := scene.Hit(ray, 0.0001, 200); hit != nil {
			if attenuation, rayOut := hit.Mat.Scatter(ray, hit, randSrc); rayOut != nil {
				accAttenuation.Mul(attenuation)
				ray = rayOut
			} else {
				// ray hit emissive surface or absorbed
				return Mul(accAttenuation, hit.Mat.Emitted(hit))
			}
		} else {
			break
		}
	}

	// ray reaches sky
	return MissShader(ray, accAttenuation)
}

func MissShader(ray *Ray, accAttenuation Vec3) Vec3 {
	skyColor := Lerp(
		Vec3{1.0, 1.0, 1.0},
		Vec3{0.5, 0.7, 1.0},
		0.5*(1.0 + ray.Direction.Y))

	return Mul(accAttenuation, skyColor)
}

func convertColor(c Vec3) color.RGBA {
	return color.RGBA{
		R: uint8(math.Min(c.X, 1.0) * 255.99),
		G: uint8(math.Min(c.Y, 1.0) * 255.99),
		B: uint8(math.Min(c.Z, 1.0) * 255.99),
		A: 0xff,
	}
}

type Task struct {
	y int
}

func RenderImage(sc *scene.Scene, cam *Camera, width, height, spp, maxBounces int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	tasks := make(chan Task, height)
	workerCount := runtime.NumCPU()

	wg := sync.WaitGroup{}
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			RenderWorker(tasks, img, width, height, sc, cam, spp, maxBounces)
		}()
	}

	for y := 0; y < height; y++ {
		tasks <- Task{y}
	}
	close(tasks)

	wg.Wait()
	return img
}

func RenderWorker(tasks <-chan Task, img *image.RGBA, width, height int, sc *scene.Scene, cam *Camera, spp, maxBounces int) {
	invWidth, invHeight := 1.0/float64(width), 1.0/float64(height)

	randSrc := rand.New(rand.NewSource(time.Now().Unix()))

	for task := range tasks {
		for x := 0; x < width; x++ {
			color := Vec3{}
			for sampleNum := 0; sampleNum < spp; sampleNum++ {
				u :=       (float64(x)      + Rand01(randSrc)) * invWidth // center value: (x + 0.5) / width; += -0.5...0.5 half-pixel scattering for multisampling
				v := 1.0 - (float64(task.y) + Rand01(randSrc)) * invHeight

				ray := cam.GetRay(u, v)
				color.Add(TraceRay(sc, ray, maxBounces, randSrc))
			}

			if spp > 1 {
				color.DivC(float64(spp))
			}

			color = LinearToSrgb(color)
			img.Set(x, task.y, convertColor(color))
		}
	}
}