package main

import (
    "image"
    "image/color"

    "github.com/hajimehoshi/go-cairo"
)

func main() {
    // Define image size and colors
    width, height := 600, 400
    surface := cairo.NewImageSurface(cairo.FormatRGB24, width, height)
    ctx := cairo.NewContext(surface)

    // Set background color
    ctx.SetSourceColor(color.White)
    ctx.Rectangle(0, 0, float64(width), float64(height))
    ctx.Fill()

    // Draw ellipse with specific center, radii, and color
    centerX, centerY := width/2, height/2
    radiusX, radiusY := 100.0, 50.0
    ctx.SetSourceColor(color.RGBA{R: 0, G: 0, B: 255, A: 255}) // Set blue color
    ctx.Arc(float64(centerX), float64(centerY), radiusX, radiusY, 0, 2*math.Pi)
    ctx.Fill()

    // Save the image to a PNG file
    err := surface.WriteToPNG("ellipse.png")
    if err !=  nil {
        panic(err)
    }
}
