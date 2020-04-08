package style

import (
	"image/color"

	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"

	"go-hep.org/x/hep/hplot"
)

var defaultBlack color.NRGBA = color.NRGBA{R: 30, G: 30, B: 30, A: 255}

// Apply nice style for the plot
func ApplyToPlot(p *hplot.Plot){

	// Plot borders
	p.Border.Right = 15
	p.Border.Left = 5
	p.Border.Top = 10
	p.Border.Bottom = 5

	// Specify title style
	p.Title.TextStyle.Font.Size = 16
	p.Title.TextStyle.Color = defaultBlack
	p.Title.Padding = 10

	// Specify axis ranges and padding (ie distance to (0,0) point
	p.X.Padding = 5
	p.Y.Padding = 5
	
	// Specify axis label & fontsize
	p.X.Label.TextStyle.Font.Size = 14
	p.Y.Label.TextStyle.Font.Size = 14
	p.X.Label.TextStyle.Color = defaultBlack
	p.Y.Label.TextStyle.Color = defaultBlack

	// Specify axis style
	p.X.LineStyle.Width = 1.05
	p.Y.LineStyle.Width = 1.05
	p.X.LineStyle.Color = defaultBlack
	p.Y.LineStyle.Color = defaultBlack

	// Specify ticks style
	p.X.Tick.LineStyle.Width = 1.05
	p.Y.Tick.LineStyle.Width = 1.05
	p.X.Tick.LineStyle.Color = defaultBlack
	p.Y.Tick.LineStyle.Color = defaultBlack
	p.X.Tick.Label.Font.Size = 11
	p.Y.Tick.Label.Font.Size = 11
	p.X.Tick.Label.Color = defaultBlack
	p.Y.Tick.Label.Color = defaultBlack

	// Specify tick position
	p.X.Tick.Marker = hplot.Ticks{N: 10}
	p.Y.Tick.Marker = hplot.Ticks{N: 10}

	// Specify text style of the legend
	p.Legend.TextStyle.Font.Size = 14
	p.Legend.TextStyle.Color = defaultBlack
}

// Apply nominal style for histogram representing data
func ApplyToDataHist(hData *hplot.H1D){

	// Remove basic stat info
	hData.Infos.Style = hplot.HInfoNone
	
	// No line
	hData.LineStyle.Width = 0

	// Y error bars
	hData.YErrs.LineStyle.Color = color.NRGBA{R: 10, G: 10, B: 10, A: 255}
	hData.YErrs.LineStyle.Width = 2.5
	hData.YErrs.CapWidth = 7

	// Dots as marker
	hData.GlyphStyle = draw.GlyphStyle{
		Shape:  draw.CircleGlyph{},
		Color:  color.NRGBA{R: 10, G: 10, B: 10, A: 255},
		Radius: vg.Points(3)}
}
