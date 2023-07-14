package main

import (
	"image"
	"render/utility"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (e *Editor) RefreshRaster(w, h int) image.Image {
	return e.renderContext.DrawBakedShapes(w, h, append(utility.MapValues(e.modelShapes), utility.MapValues(e.editorShapes)...))
}

func (e *Editor) NumberOfModelShapes() int {
	return len(e.modelShapes)
}

func (e *Editor) CreateModelShapeListEntry() fyne.CanvasObject {
	return widget.NewLabel("New Shape")
}

func (e *Editor) UpdateModelShapeListEntry(w widget.ListItemID, o fyne.CanvasObject) {
	o.(*widget.Label).SetText("Updated Shape")
}
