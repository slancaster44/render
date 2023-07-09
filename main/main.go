package main

import (
	"render"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	renderApp := app.New()
	renderWindow := renderApp.NewWindow("Render Canvas")
	renderContext := render.NewRayTracer(900, 550)
	renderWindow.Resize(fyne.NewSize(900, 550))

	go func() {
		img := renderContext.Render()
		renderRaster := canvas.NewRasterFromImage(img)
		renderWindow.SetContent(renderRaster)

		for range time.Tick(time.Second / 60) {
			img = renderContext.Render()
			renderRaster.Refresh()
		}
	}()

	renderWindow.CenterOnScreen()
	renderWindow.ShowAndRun()

}
