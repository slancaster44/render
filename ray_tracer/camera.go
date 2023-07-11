package ray_tracer

import (
	"math"
	"render/transformation"
	"render/tuple"
	"render/utility"
)

type Camera struct {
	canvasHeight float64
	canvasWidth  float64
	Fov          float64
	location     *tuple.Tuple3
	orientation  *tuple.Tuple3

	half_width  float64
	half_height float64
	pixel_size  float64
}

func NewCamera(h, w int) *Camera {
	c := &Camera{
		canvasHeight: float64(h),
		canvasWidth:  float64(w),
		Fov:          math.Pi / 2.0,
		location:     tuple.NewPnt3(0, 0, 10),
		orientation:  tuple.NewVec3(0, 0, 1),
	}

	c.computeRayParamaters()
	return c
}

func (c *Camera) TransformLocation(t transformation.Transformation) {
	c.location = t(c.location)
	c.computeRayParamaters()
}

func (c *Camera) TransformOrientation(t transformation.Transformation) {
	c.orientation = t(c.orientation)
	c.computeRayParamaters()
}

func (c *Camera) computeRayParamaters() {
	half_view := utility.Tan(c.Fov / 2)
	aspect_ratio := c.canvasHeight / c.canvasWidth

	if aspect_ratio >= 1.0 {
		c.half_width = half_view
		c.half_height = half_view / aspect_ratio
	} else {
		c.half_width = half_view * aspect_ratio
		c.half_height = half_view
	}

	c.pixel_size = (c.half_width * 2) / c.canvasHeight
}

func (c *Camera) RayForPixel(x, y int) *Ray {
	xoffset := (float64(x) + 0.5) * c.pixel_size
	yoffset := (float64(y) + 0.5) * c.pixel_size

	worldX := c.half_width - xoffset
	worldY := c.half_height - yoffset
	rayDirection := tuple.NewVec3(worldX, worldY, 0)
	rayDirection = tuple.Normalize(tuple.Add(rayDirection, c.orientation))

	return &Ray{
		Origin:    c.location,
		Direction: rayDirection,
	}
}
