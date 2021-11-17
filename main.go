package main

import (
	"context"
	"github.com/skoptew/gotracing101/material"
	. "github.com/skoptew/gotracing101/math101"
	"github.com/skoptew/gotracing101/rendering"
	"github.com/skoptew/gotracing101/scene"
	"math/rand"
	"os/signal"

	"flag"
	"fmt"
	"image/png"
	"os"
	"time"
)

func main() {
	// handle SIGTERM: stop rendering, save rendered part of image
	ctx, cancelFunc := context.WithCancel(context.Background())
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt)

	defer func() {
		signal.Stop(termChan)
		cancelFunc()
	}()

	go func() {
		select {
		case <-termChan:
			fmt.Println("rendering terminated")
			cancelFunc()
		}
	}()

	// setup rendering parameters
	width, height, spp, maxBounces, numObjects, fname := ParseFlags()
	cam := CreateCamera(float64(width) / float64(height))
	sc := CreateTestScene(numObjects)

	// render image
	fmt.Printf("rendering %vx%v image (%v objects, %v samples per pixel, up to %v bounces)...\n", width, height, numObjects, spp, maxBounces)
	startTime := time.Now()
	img := rendering.RenderImage(ctx, sc, cam, width, height, spp, maxBounces)
	fmt.Printf("done for %s\n", time.Since(startTime))

	// save image
	fmt.Printf("saving to %s...\n", fname)
	file, err := os.Create(fname)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := png.Encode(file, img); err != nil {
		fmt.Println(err)
	}
}

func ParseFlags() (width, height, spp, maxBounces, numObjects int, fname string) {
	widthFlag := flag.Int("width", 960, "image width")
	heightFlag := flag.Int("height", 540, "image height")
	sppFlag := flag.Int("spp", 64, "samples per pixel")
	maxBouncesFlag := flag.Int("bounces", 32, "max bounces per path")
	numObjectsFlag := flag.Int("objnum", 24, "num of objects in test scene")
	fnameFlag := flag.String("filename", "out", "out file name (without extension, forced to .png)")

	flag.Parse()

	width = clamp(*widthFlag, 1, 4096)
	height = clamp(*heightFlag, 1, 4096)
	spp = clamp(*sppFlag, 1, 1024)
	maxBounces = clamp(*maxBouncesFlag, 1, 512)
	numObjects = clamp(*numObjectsFlag, 0, 64)
	fname = *fnameFlag + ".png"

	return
}

func CreateTestScene(numObjects int) *scene.Scene {
	sc := scene.NewScene()
	randSrc := rand.New(rand.NewSource(time.Now().Unix()))

	// room
	{
		matGround := material.NewMatReflective(Vec3{0.5, 0.5, 0.5}, 0.005)
		sc.Add(scene.NewSphere(Vec3{0, -51, -1}, 50.0, matGround))

		matCeiling := material.NewMatLambertian(Vec3{1, 1, 1})
		sc.Add(scene.NewSphere(Vec3{0, 54, -1}, 50.0, matCeiling))

		matRight := material.NewMatReflective(Vec3{0.9, 0.05, 0.05}, 0)
		sc.Add(scene.NewSphere(Vec3{29, 0, -1}, 25.0, matRight))

		matLeft := material.NewMatReflective(Vec3{0.05, 0.9, 0.05}, 0)
		sc.Add(scene.NewSphere(Vec3{-29, 0, -1}, 25.0, matLeft))

		matBack := material.NewMatReflective(Vec3{0.5, 0.5, 0.5}, 0)
		sc.Add(scene.NewSphere(Vec3{0, 0, -35}, 25.0, matBack))

		matEmissiveRed := material.NewMatEmissive(Vec3{1, 0.9, 0.9}, 1.4)
		sc.Add(scene.NewSphere(Vec3{26, -1, -25}, 25.0, matEmissiveRed))

		matEmissiveGreen := material.NewMatEmissive(Vec3{0.9, 1.0, 0.9}, 1.4)
		sc.Add(scene.NewSphere(Vec3{-26, -1, -25}, 25.0, matEmissiveGreen))
	}

	// random, not intersected spheres
	{
		const rad = 0.35
		const diam2 = 4 * rad * rad

		centers := make([]Vec3, numObjects)
		for i := 0; i < numObjects; i++ {
			for {
				pt := Vec3{Rand(randSrc, -3, 3),
					Rand(randSrc, -1+rad, 1.0),
					Rand(randSrc, -2.5, -4.0)}

				intersection := false
				for j := 0; j < i; j++ {
					v := Sub(centers[j], pt)
					if v.Length2() < diam2 {
						intersection = true
						break
					}
				}

				if !intersection {
					centers[i] = pt
					break
				}
			}
		}

		for _, center := range centers {
			var mat material.Material

			switch matType := randSrc.Intn(100); {
			case matType < 70: // reflective material
				albedo := Vec3{0.7, 0.7, 0.7}
				switch c := randSrc.Intn(100); {
				case c < 45:
					albedo = Vec3{0.17, 0.24, 0.60}
				case c < 90:
					albedo = Vec3{0.7, 0.15, 0.06}
				}
				mat = material.NewMatReflective(albedo, 0)

			default:
				mat = material.NewMatRefractive(1.7, Vec3{0.95, 0.95, 0.95})
			}

			sc.Add(scene.NewSphere(center, rad, mat))
		}
	}
	return sc
}

func CreateCamera(aspect float64) *rendering.Camera {
	cam := rendering.NewCamera(aspect)
	cam.Set(
		Vec3{X: 0, Y: 0, Z: 0},
		Vec3{X: 0, Y: 0, Z: -1},
		60.0)

	return cam
}

func clamp(x, min, max int) int {
	if x < min {
		x = min
	}
	if x > max {
		x = max
	}
	return x
}
