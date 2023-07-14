package renderer

import (
	"image"
	"render/utility"
)

type Renderer struct {
	NumThreads                   int
	ScreenOriginX, ScreenOriginY int
}

func New() *Renderer {
	return &Renderer{NumThreads: 1, ScreenOriginX: 0, ScreenOriginY: 0}
}

func (r *Renderer) BakeAndDrawShapes(frame_w, frame_h int, shapes []Shape) image.Image {
	if r.NumThreads == 1 {
		for _, t := range shapes {
			t.Bake()
		}
	} else {
		utility.MethodOn(shapes, r.NumThreads, "Bake")
	}

	return r.DrawBakedShapes(frame_w, frame_h, shapes)
}

func (r *Renderer) DrawBakedShapes(frame_w, frame_h int, shapes []Shape) image.Image {
	output := image.NewRGBA(image.Rect(r.ScreenOriginX, r.ScreenOriginY, r.ScreenOriginX+frame_w, r.ScreenOriginY+frame_h))

	for _, t := range shapes {
		for _, pix := range t.Pixels() {
			output.SetRGBA(int(pix.X), int(pix.Y), t.Color())
		}
	}

	return output
}
