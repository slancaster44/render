package render

import (
	"image"
	"image/color"
	"render/tuple"
)

type RayTracer struct {
	output *image.RGBA
	trueX  int
	trueY  int
}

func NewRayTracer(frameWidth, frameHeight int) *RayTracer {
	retVal := &RayTracer{
		output: image.NewRGBA(image.Rect(0, 0, frameWidth, frameHeight)),
		trueX:  0,
		trueY:  0,
	}

	return retVal
}

func (r *RayTracer) put(x, y float64, c color.RGBA) {
	ix := int(x)
	iy := int(y)
	r.output.SetRGBA(ix, iy, c)
}

var runs uint8 = 0

func (r *RayTracer) Render() image.Image {
	runs++

	start := tuple.NewPnt3(0, 1, 0)

	velocity := tuple.ScalarSlowMultiply(tuple.Normalize(tuple.NewVec3(1, 3, 0)), 10)
	gravity := tuple.NewVec3(0, -0.1, 0)
	wind := tuple.NewVec3(-0.01, 0, 0)

	color := color.RGBA{runs, runs / 2, runs / 3, 255}

	for start.Y > 0 {
		start = tuple.Add(start, velocity)
		start = tuple.Add(start, wind)

		velocity = tuple.Add(gravity, velocity)

		r.put(start.X, start.Y, color)
	}

	return r.output
}
