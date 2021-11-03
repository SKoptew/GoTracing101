package main

import (
	"gotracing101/rendering"

	"flag"
	"fmt"
	"image/png"
	"os"
	"time"
)

func clamp(x, min, max int) int {
	if x < min {
		x = min
	}
	if x > max {
		x = max
	}
	return x
}

func main() {
	widthFlag  := flag.Int("width",  1024, "image width")
	heightFlag := flag.Int("height", 768, "image height")
	fnameFlag  := flag.String("filename", "out", "out file name (without extension, forced to .png)")

	flag.Parse()

	width  := clamp(*widthFlag, 1, 2048)
	height := clamp(*heightFlag, 1, 2048)
	fname  := *fnameFlag + ".png"

	fmt.Printf("rendering %vx%v image...\n", width, height)
	startTime := time.Now()
	img := rendering.RenderImage(width, height)
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