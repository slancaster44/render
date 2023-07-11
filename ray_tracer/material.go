package ray_tracer

import "image/color"

type Material struct {
	Color              color.RGBA
	AmbientReflection  float64
	DiffuseReflection  float64
	SpecularReflection float64
	Shininess          float64
}
