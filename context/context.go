package context

import "image"

type Context interface {
	Render() image.Image
}
