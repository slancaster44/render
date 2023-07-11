package ray_tracer

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"os"
	"render/tuple"
	"strconv"
	"strings"
	"sync"
)

type RayTracer struct {
	output *image.RGBA
	camera *Camera
	lights []*Light
	shapes []*Triangle
}

type pixel struct {
	x int
	y int
	c color.RGBA
}

func NewRayTracer(frameWidth, frameHeight int) *RayTracer {
	retVal := &RayTracer{
		output: image.NewRGBA(image.Rect(0, 0, frameWidth, frameHeight)),
		camera: NewCamera(frameHeight, frameWidth),
	}

	return retVal
}

func (r *RayTracer) AddShape(s *Triangle) {
	r.shapes = append(r.shapes, s)
}

func (r *RayTracer) AddLight(l *Light) {
	r.lights = append(r.lights, l)
}

func (r *RayTracer) screenCoordToIndex(x, y int) int {
	return (y * r.output.Rect.Dx()) + x
}

func (r *RayTracer) screenIndexToCoord(index int) (int, int) {
	return index % r.output.Rect.Dx(), index / r.output.Rect.Dx()
}

var RENDER_THREADS int = 4

func (r *RayTracer) Render(_, _ int) image.Image {
	fmt.Println("Entering render routine")

	var wg sync.WaitGroup

	helper := func(start, finish int, ch chan pixel) {
		defer wg.Done()

		for i := 0; i < finish; i++ {
			cameraX, cameraY := r.screenIndexToCoord(i)
			ray := r.camera.RayForPixel(cameraX, cameraY)

			for _, s := range r.shapes {
				hit, nv := s.Intersect(ray)

				if hit != nil && hit.TValue > 0 {
					p := ray.PositionAt(hit.TValue)
					ev := tuple.Negate(ray.Direction)

					c := color.RGBA{0, 0, 0, 0}

					for _, light := range r.lights {
						c_delta := r.CalculateLighting(s.Material, light, p, ev, nv)
						c.R += c_delta.R
						c.G += c_delta.G
						c.B += c_delta.B
						c.A += c_delta.A
					}

					ch <- pixel{cameraX, cameraY, c}
				}

			}
		}

	}

	thread_chunk_size := (r.output.Rect.Dx() * r.output.Rect.Dy()) / RENDER_THREADS
	ch := make(chan pixel, r.output.Rect.Dx()*r.output.Rect.Dy())

	for i := 0; i < RENDER_THREADS; i++ {
		wg.Add(1)
		go helper(i*thread_chunk_size, (i*thread_chunk_size)+thread_chunk_size, ch)
	}
	wg.Wait()

	for {
		select {
		case result_pixel := <-ch:
			r.output.SetRGBA(result_pixel.x, result_pixel.y, result_pixel.c)
		default:
			goto exit
		}
	}

exit:
	fmt.Println("Exiting Render Routine")
	return r.output
}

func (r *RayTracer) CalculateLighting(
	material *Material,
	light *Light, pos *tuple.Tuple3,
	eye_vector *tuple.Tuple3,
	normal_vector *tuple.Tuple3) color.RGBA {

	effective_color := tuple.ScalarMultiply(tuple.ColorToVec(material.Color), light.Intensity)

	light_vector := tuple.Normalize(tuple.Subtract(light.Location, pos))
	light_dot_normal := tuple.DotProduct(light_vector, normal_vector)

	ambient := tuple.ScalarMultiply(effective_color, material.AmbientReflection)

	var diffuse *tuple.Tuple3
	var specular *tuple.Tuple3

	if light_dot_normal < 0 {
		diffuse = tuple.NewVec3(0, 0, 0)
		specular = tuple.NewVec3(0, 0, 0)
	} else {
		diffuse = tuple.ScalarMultiply(effective_color, material.DiffuseReflection*light_dot_normal)

		reflect_vector := tuple.Reflect(tuple.Negate(light_vector), normal_vector)
		reflect_dot_eye := tuple.DotProduct(reflect_vector, eye_vector)

		if reflect_dot_eye <= 0 {
			specular = tuple.NewVec3(0, 0, 0)
		} else {
			factor := math.Pow(reflect_dot_eye, material.Shininess)
			specular = tuple.ScalarMultiply(tuple.ColorToVec(light.Color), material.SpecularReflection*factor)
		}
	}

	final := tuple.Add(diffuse, tuple.Add(specular, ambient))
	return tuple.VecToColor(final)
}

func (r *RayTracer) ReadFile(filename string) {
	vertexes := []*tuple.Tuple3{}

	text, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	UNKOWN := 0
	VERTEX := 1
	FACE := 2
	MODE := UNKOWN

	curItem := [3]float64{}

	split_text := strings.Fields(string(text))
	for i := 0; i < len(split_text); i++ {
		switch split_text[i] {
		case "v":
			MODE = VERTEX
		case "f":
			MODE = FACE
		default:
			if v1, err := strconv.ParseFloat(split_text[i], 64); err == nil {
				curItem[0] = v1
				i++

				curItem[1], err = strconv.ParseFloat(split_text[i], 64)
				i++

				if err != nil {
					panic(err)
				}

				curItem[2], err = strconv.ParseFloat(split_text[i], 64)
				if err != nil {
					panic(err)
				}

				if MODE == VERTEX {
					vertexes = append(vertexes, tuple.NewPnt3(curItem[0], curItem[1], curItem[2]))
				} else if MODE == FACE {
					p1 := vertexes[int(curItem[0])-1]
					p2 := vertexes[int(curItem[0])-1]
					p3 := vertexes[int(curItem[0])-1]

					r.AddShape(NewTriangle(p1, p2, p3))
				}
			}
		}
	}

}
