// +build darwin linux

package main

import (
	"image"
	"image/draw"
	"math"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/geom"
)

const (
	dpi = 72
)

type TextAlign int

const (
	Center TextAlign = iota
	Left
	Right
)

type TextSprite struct {
	placeholder     string
	text            string
	font            *truetype.Font
	widthPx         int
	heightPx        int
	textColor       *image.Uniform
	backgroundColor *image.Uniform
	fontSize        float64
	xPt             geom.Pt
	yPt             geom.Pt
	align           TextAlign
}

func (ts TextSprite) Render(sz size.Event) {
	sprite := images.NewImage(ts.widthPx, ts.heightPx)

	draw.Draw(sprite.RGBA, sprite.RGBA.Bounds(), ts.backgroundColor, image.ZP, draw.Src)

	d := &font.Drawer{
		Dst: sprite.RGBA,
		Src: ts.textColor,
		Face: truetype.NewFace(ts.font, &truetype.Options{
			Size:    ts.fontSize,
			DPI:     dpi,
			Hinting: font.HintingNone,
		}),
	}

	dy := int(math.Ceil(ts.fontSize * dpi / dpi))
	var textWidth fixed.Int26_6
	if ts.placeholder == "" {
		textWidth = d.MeasureString(ts.text)
	} else {
		textWidth = d.MeasureString(ts.placeholder)
	}

	switch ts.align {
	case Center:
		d.Dot = fixed.Point26_6{
			X: fixed.I(sz.Size().X/2) - (textWidth / 2),
			Y: fixed.I(ts.heightPx/2 + dy/2),
		}
	case Left:
		d.Dot = fixed.Point26_6{
			X: fixed.I(0),
			Y: fixed.I(ts.heightPx/2 + dy/2),
		}
	case Right:
		d.Dot = fixed.Point26_6{
			X: fixed.I(sz.Size().X) - textWidth,
			Y: fixed.I(ts.heightPx/2 + dy/2),
		}
	}

	d.DrawString(ts.text)

	sprite.Upload()
	sprite.Draw(
		sz,
		geom.Point{X: ts.xPt, Y: ts.yPt},
		geom.Point{X: ts.xPt + sz.WidthPt, Y: ts.yPt},
		geom.Point{X: ts.xPt, Y: ts.yPt + sz.HeightPt},
		sz.Bounds())
	sprite.Release()

}
