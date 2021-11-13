# GoTracing101

Simple path tracer written in Golang, mostly for learning pusposes.

Command line flags:
-width=960		width of resulting image
-height=540		height of resulting image
-spp=64			samples per pixel
-bounces=32		maximum bounces per ray
-objnum=24		number of objects (small spheres) at test scene
-filename=out 	out file name without extension

Interruption (Ctrl+C) will stop rendering, but rendered part of image will be saved.


![Sample_output](Readme_test_render.jpg)