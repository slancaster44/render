package main

import (
	"render/context"
	"render/tuple"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	tuple.Add_x86_Simd(tuple.NewVec3(1, 1, 1), tuple.NewVec3(2, 2, 2))
}
func runWindow() {
	renderApp := app.New()
	renderWindow := renderApp.NewWindow("Render Canvas")
	renderContext := context.NewRayTracer()

	img := renderContext.Render()

	renderRaster := canvas.NewRasterFromImage(img)
	renderWindow.SetContent(renderRaster)

	img_bounds := img.Bounds()
	width := img_bounds.Max.X - img_bounds.Min.X
	height := img_bounds.Max.Y - img_bounds.Min.Y

	renderWindow.Resize(fyne.NewSize(float32(width), float32(height)))
	renderWindow.ShowAndRun()

}
