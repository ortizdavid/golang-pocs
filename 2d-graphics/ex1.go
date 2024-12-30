package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "image/color"
    "log"
)

type Game struct{}

func (g *Game) Update() error {
    // Update the game logic here
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    // Clear the screen with a solid color
    screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})

    // Draw a rectangle
    x, y := 50, 50
    w, h := 100, 50
    col := color.RGBA{0x00, 0x00, 0xff, 0xff} // Blue color
    ebitenutil.DrawRect(screen, float64(x), float64(y), float64(w), float64(h), col)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    // Return the screen dimensions
    return 640, 480
}

func main() {
    game := &Game{}
    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("2D Graphics with Ebiten")

    // Start the game loop
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
