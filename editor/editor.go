package main

import (
	"render/renderer"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Editor struct {
	fyneApp fyne.App

	rasterWindow  fyne.Window
	raster        *canvas.Raster
	renderContext *renderer.Renderer

	editToolsWindow fyne.Window
	shapesTree      *widget.List

	modelShapes  map[string]renderer.Shape
	editorShapes map[string]renderer.Shape
}

func NewEditor() *Editor {
	e := &Editor{}

	e.renderContext = renderer.New()

	e.fyneApp = app.New()
	e.rasterWindow = e.fyneApp.NewWindow("Render Canvas")

	e.raster = canvas.NewRaster(e.RefreshRaster)
	e.rasterWindow.SetContent(e.raster)

	e.editToolsWindow = e.fyneApp.NewWindow("Edit")
	e.shapesTree = widget.NewList(
		e.NumberOfModelShapes,
		e.CreateModelShapeListEntry,
		e.UpdateModelShapeListEntry)
	e.editToolsWindow.SetContent(e.shapesTree)

	return e
}

func (e *Editor) Go() {
	go func() {
		for _ = range time.Tick(time.Millisecond * 17) {
			e.raster.Refresh()
		}
	}()

	e.rasterWindow.Resize(fyne.NewSize(900, 550))
	e.rasterWindow.Show()

	e.editToolsWindow.CenterOnScreen()
	e.editToolsWindow.ShowAndRun()
}
