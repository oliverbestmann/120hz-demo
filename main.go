package main

import (
	"fmt"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var startTime = time.Now()

type Game struct {
	mode int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var toScreen ebiten.GeoM
	scale := float64(screen.Bounds().Dy()) / 1000.0
	toScreen.Translate(0, 500)
	toScreen.Scale(scale, scale)

	now := time.Since(startTime)
	if g.mode == 0 {
		g.mode0(screen, toScreen, now)
	}

	if g.mode == 1 {
		g.mode1(screen, toScreen, now)

	}

	text := fmt.Sprintf("FPS: %1.2f", ebiten.ActualFPS())
	ebitenutil.DebugPrintAt(screen, text, 16, 16)
}

func (g *Game) mode0(screen *ebiten.Image, toScreen ebiten.GeoM, now time.Duration) {
	var p vector.Path
	p.Arc(0, 0, 100, 0, math.Pi*2, vector.Clockwise)

	drawAt := func(x float64, scale float64) {
		idx := int(now.Seconds() * scale)

		if idx%2 != 0 {
			return
		}

		t := float64(idx) / scale
		y := math.Sin(5*t) * 400

		var model ebiten.GeoM
		model.Translate(x, y)

		model.Concat(toScreen)

		var out vector.Path
		out.AddPath(&p, &vector.AddPathOptions{GeoM: model})
		vector.FillPath(screen, &out, nil, nil)
	}

	drawAt(200, 30)
	drawAt(500, 60)
	drawAt(800, 120)
}

func (g *Game) mode1(screen *ebiten.Image, toScreen ebiten.GeoM, now time.Duration) {
	idx := int64(now.Seconds() * 120)

	var x, y float64

	switch idx % 4 {
	case 0:
		x = 200
		y = -300

	case 1:
		x = 800
		y = -300

	case 2:
		x = 800
		y = 300

	case 3:
		x = 200
		y = 300
	}

	x, y = toScreen.Apply(x, y)

	var p vector.Path
	p.Arc(float32(x), float32(y), 100, 0, math.Pi*2, vector.Clockwise)
	vector.FillPath(screen, &p, nil, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	err := ebiten.RunGame(&Game{mode: 1})
	if err != nil {
		panic(err)
	}
}
