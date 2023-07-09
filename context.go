package render

import "image"

type Context interface {
	Render() image.Image
}
