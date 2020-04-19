package style

import (
	"image/color"

	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"

	"go-hep.org/x/hep/hplot"
	"go-hep.org/x/hep/hplot/htex"
)

var (
	SmoothBlack = color.NRGBA{R: 35, G: 55, B: 57, A: 255}
)

// Apply nice options to Figure (currently type hplot.P)
func ApplyToFigure(p *hplot.P, latex htex.Handler) {

	// Plot borders
	p.Border.Right = 15
	p.Border.Left = 5
	p.Border.Top = 10
	p.Border.Bottom = 5

	// LaTex compilation setup
	p.Latex = latex
}

// Apply nice style for the plot
func ApplyToPlot(p *hplot.Plot){

	// Specify title style
	p.Title.TextStyle.Font.Size = 16
	p.Title.TextStyle.Color = SmoothBlack
	p.Title.Padding = 10

	// Specify axis ranges and padding (ie distance to (0,0) point
	p.X.Padding = 5
	p.Y.Padding = 5
	
	// Specify axis label & fontsize
	p.X.Label.Padding = 5
	p.Y.Label.Padding = 5
	p.X.Label.TextStyle.Font.Size = 14
	p.Y.Label.TextStyle.Font.Size = 14
	// p.X.Label.Position = draw.PosRight  // doesn't work with LaTeX --> disable for now
	// p.Y.Label.Position = draw.PosTop    // doesn't work with LaTeX --> disable for now
	p.X.Label.TextStyle.Color = SmoothBlack
	p.Y.Label.TextStyle.Color = SmoothBlack

	// Specify axis style
	p.X.LineStyle.Width = 1.05
	p.Y.LineStyle.Width = 1.05
	p.X.LineStyle.Color = SmoothBlack
	p.Y.LineStyle.Color = SmoothBlack

	// Specify ticks style
	p.X.Tick.LineStyle.Width = 1.05
	p.Y.Tick.LineStyle.Width = 1.05
	p.X.Tick.LineStyle.Color = SmoothBlack
	p.Y.Tick.LineStyle.Color = SmoothBlack
	p.X.Tick.Label.Font.Size = 11
	p.Y.Tick.Label.Font.Size = 11
	p.X.Tick.Label.Color = SmoothBlack
	p.Y.Tick.Label.Color = SmoothBlack

	// Specify tick position
	p.X.Tick.Marker = hplot.Ticks{N: 10}
	p.Y.Tick.Marker = hplot.Ticks{N: 10}
	
	// Specify text style of the legend
	p.Legend.TextStyle.Font.Size = 14
	p.Legend.TextStyle.Color = SmoothBlack
	p.Legend.Padding = 0.1 * vg.Inch
}

func ApplyToBottomPlot(p *hplot.Plot){
	ApplyToPlot(p)
	
	p.Y.Tick.Marker = hplot.Ticks{N: 3}
	p.Y.Tick.Label.Font.Size = 10
}

// Apply cirle markers
func SetCircleMarkersTo(h *hplot.H1D){

	// Dots as marker
	h.GlyphStyle = draw.GlyphStyle{
		Shape: draw.CircleGlyph{},
		Color: SmoothBlack,
		Radius: vg.Points(3)}

	// Y error bars
	if h.YErrs != nil {
		h.YErrs.LineStyle.Color = SmoothBlack
		h.YErrs.LineStyle.Width = 2.5
		h.YErrs.CapWidth = 7
	}
}

// Apply nominal style for histogram representing data
func ApplyToDataHist(hData *hplot.H1D){

	// Remove basic stat info
	hData.Infos.Style = hplot.HInfoNone
	
	// No line
	hData.LineStyle.Width = 0
	
	// Apply circles as marker
	SetCircleMarkersTo(hData)

}
