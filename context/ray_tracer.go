package context

import (
	"image"
	"image/jpeg"
	"os"
	"render/utility"
)

type RayTracer struct {
}

func NewRayTracer() *RayTracer {
	return &RayTracer{}
}

func (r *RayTracer) Render() image.Image {
	f, err1 := os.Open("./testing_materials/test.jpeg")
	utility.Check(err1)
	defer f.Close()

	img, err2 := jpeg.Decode(f)
	utility.Check(err2)

	return img
}
