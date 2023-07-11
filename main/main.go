package main

import (
	"image/color"
	"render/ray_tracer"
	"render/tuple"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

const WIDTH int = 900
const HEIGHT int = 550

func main() {
	renderApp := app.New()
	renderWindow := renderApp.NewWindow("Render Canvas")
	renderContext := ray_tracer.NewRayTracer(WIDTH, HEIGHT)
	renderContext.ReadFile("teapot.obj")

	l := ray_tracer.NewLight(tuple.NewPnt3(100, 100, 100), color.RGBA{255, 255, 255, 255})
	renderContext.AddLight(l)

	m := ray_tracer.Material{color.RGBA{128, 255, 255, 255}, 0.1, 0.9, 0.9, 200}
	t := ray_tracer.Triangle{
		tuple.NewPnt3(0, 0, 0),
		tuple.NewPnt3(0, 1, 4),
		tuple.NewPnt3(5, 0, 0),
		&m,
	}
	renderContext.AddShape(&t)

	renderWindow.Resize(fyne.NewSize(float32(WIDTH), float32(HEIGHT)))
	renderRaster := canvas.NewRaster(renderContext.Render)
	renderWindow.SetContent(renderRaster)

	/*go func() {
		for range time.Tick(time.Second / 60) {
			renderRaster.Refresh()
		}
	}()*/

	renderWindow.CenterOnScreen()
	renderWindow.ShowAndRun()

}
