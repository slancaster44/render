package renderer

import (
	"image/color"
	"image/jpeg"
	"os"
	"render/tuple"
	"testing"
)

func TestRenderer(t *testing.T) {
	r := New()
	r.NumThreads = 4

	shapes := []Shape{}
	p1 := tuple.NewPnt3(50, 50, 0)
	p2 := tuple.NewPnt3(50, 300, 0)
	p3 := tuple.NewPnt3(300, 300, 0)
	tri := NewIsometricTriangle(p1, p2, p3, color.RGBA{0, 128, 255, 255})
	shapes = append(shapes, tri)

	l := NewIsometricLine(p1, p2, color.RGBA{255, 0, 0, 255})
	shapes = append(shapes, l)

	img := r.BakeAndDrawShapes(900, 550, shapes)

	red, green, blue, alpha := img.At(50, 75).RGBA()
	if (red / 256) != 255 {
		t.Errorf("Render fail %v %v %v %v", red/255, green, blue, alpha)
	}

	os.Remove("output.jpg")
	f, err := os.Create("output.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err = jpeg.Encode(f, img, nil); err != nil {
		panic(err)
	}
}

var NUM_TRIANGLES int = 10000

func BenchmarkBakeAndDrawShapes(b *testing.B) {
	r := New()
	r.NumThreads = 4

	triangles := []Shape{}

	for i := 0; i < NUM_TRIANGLES; i++ {
		p1 := tuple.NewPnt3(50, 50, 0)
		p2 := tuple.NewPnt3(50, 100, 0)
		p3 := tuple.NewPnt3(100, 100, 0)
		tri := NewIsometricTriangle(p1, p2, p3, color.RGBA{128, 255, 255, 255})
		triangles = append(triangles, tri)
	}

	b.ResetTimer()
	img := r.BakeAndDrawShapes(900, 550, triangles)
	b.StopTimer()

	os.Remove("output.jpg")
	f, err := os.Create("output.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err = jpeg.Encode(f, img, nil); err != nil {
		panic(err)
	}
}
