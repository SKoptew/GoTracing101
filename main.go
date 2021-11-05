package main

import (
	"gotracing101/material"
	. "gotracing101/math101"
	"gotracing101/rendering"
	"gotracing101/scene"

	"flag"
	"fmt"
	"image/png"
	"os"
	"time"
)

func main() {
	width, height, spp, maxBounces, fname := ParseFlags()

	cam := CreateCamera(float64(width)/float64(height))
	sc  := CreateTestScene()

	fmt.Printf("rendering %vx%v image (%v samples per pixel, up to %v bounces)...\n", width, height, spp, maxBounces)
	startTime := time.Now()
	img := rendering.RenderImage(sc, cam, width, height, spp, maxBounces)
	fmt.Printf("done for %s\n", time.Since(startTime))

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

func ParseFlags() (width, height, spp, maxBounces int, fname string) {
	widthFlag      := flag.Int("width",  1024, "image width")
	heightFlag     := flag.Int("height", 768, "image height")
	sppFlag        := flag.Int("spp", 32, "samples per pixel")
	maxBouncesFlag := flag.Int("bounces", 32, "max bounces per path")
	fnameFlag      := flag.String("filename", "out", "out file name (without extension, forced to .png)")

	flag.Parse()

	width      = clamp(*widthFlag, 1, 4096)
	height     = clamp(*heightFlag, 1, 4096)
	spp        = clamp(*sppFlag, 1, 1024)
	maxBounces = clamp(*maxBouncesFlag, 1, 512)
	fname      = *fnameFlag + ".png"

	return
}

func CreateTestScene() *scene.Scene {
	sc := scene.NewScene()

	matGround := material.NewMatLambertian(Vec3{0.1, 0.4, 0.1})
	matRed    := material.NewMatLambertian(Color(250, 20, 20))
	//matMirror := material.NewMatReflective(Vec3{0.95,0.95,0.95}, 0.0)
	matGold   := material.NewMatReflective(Vec3{0.8,0.6,0.2}, 0.025)
	matGlass  := material.NewMatRefractive(1.5, Vec3{1,1,1})

	sc.Add(scene.NewSphere(Vec3{ 0, -100.5, -1}, 100.0, matGround))

	sc.Add(scene.NewSphere(Vec3{ 0, 0, -1}, 0.5, matRed))
	sc.Add(scene.NewSphere(Vec3{-1, 0, -1}, 0.5, matGlass))
	sc.Add(scene.NewSphere(Vec3{ 1, 0, -1}, 0.5, matGold))

	return sc
}

func CreateCamera(aspect float64) *rendering.Camera {
	cam := rendering.NewCamera(aspect)
	cam.Set(
		Vec3{X: 0, Y: 0, Z: 0},
		Vec3{X: 0, Y: 0, Z: -1},
		90.0)

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