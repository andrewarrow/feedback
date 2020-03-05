// +build darwin linux

package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"strings"
	"time"

	"github.com/golang/freetype/truetype"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/gl"
)

type Game struct {
	lastCalc   clock.Time
	touchCount uint64
	font       *truetype.Font
}

func NewGame() *Game {
	var g Game
	g.reset()
	return &g
}

func (g *Game) reset() {
	var err error
	g.font, err = LoadCustomFont()
	if err != nil {
		log.Fatalf("error parsing font: %v", err)
	}
}

func (g *Game) Touch(down bool) {
	if down {
		go hitApi()
		g.touchCount++
	}
}

func (g *Game) Update(now clock.Time) {
	for ; g.lastCalc < now; g.lastCalc++ {
		g.calcFrame()
	}
}

func (g *Game) calcFrame() {

}

func (g *Game) Render(sz size.Event, glctx gl.Context, images *glutil.Images) {
	headerHeightPx, footerHeightPx := 100, 100

	header := &TextSprite{
		text:            fmt.Sprintf("%vpx * %vpx", sz.WidthPx, sz.HeightPx),
		font:            g.font,
		widthPx:         sz.WidthPx,
		heightPx:        headerHeightPx * 2,
		textColor:       image.White,
		backgroundColor: image.NewUniform(color.RGBA{0x31, 0xA6, 0xA2, 0xFF}),
		fontSize:        30,
		xPt:             10,
		yPt:             -5,
		align:           Center,
	}
	header.Render(sz)

	loading := &TextSprite{
		placeholder:     "feedback",
		text:            display + strings.Repeat(".", int(time.Now().Unix()%4)),
		font:            g.font,
		widthPx:         sz.WidthPx,
		heightPx:        sz.HeightPx - headerHeightPx - footerHeightPx,
		textColor:       image.White,
		backgroundColor: image.NewUniform(color.RGBA{0x35, 0x67, 0x99, 0xFF}),
		fontSize:        96,
		xPt:             0,
		yPt:             PxToPt(sz, headerHeightPx),
	}
	loading.Render(sz)

	footer := &TextSprite{
		text:            fmt.Sprintf("%d", g.touchCount),
		font:            g.font,
		widthPx:         sz.WidthPx,
		heightPx:        footerHeightPx,
		textColor:       image.White,
		backgroundColor: image.NewUniform(color.RGBA{0x31, 0xA6, 0xA2, 0xFF}),
		fontSize:        30,
		xPt:             0,
		yPt:             PxToPt(sz, sz.HeightPx-footerHeightPx-20),
		align:           Center,
	}
	footer.Render(sz)
}

func PxToPt(sz size.Event, sizePx int) geom.Pt {
	return geom.Pt(float32(sizePx) / sz.PixelsPerPt)
}
