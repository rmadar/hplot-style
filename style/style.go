package style

import (
	"image/color"

	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"

	"go-hep.org/x/hep/hplot"
)

var (
	SmoothBlack = color.NRGBA{R: 35, G: 55, B: 57, A: 255}
)

// Apply nice options to Figure
func ApplyToFigure(fig *hplot.Fig) {

	// Plot borders
	fig.Border.Right = 15
	fig.Border.Left = 5
	fig.Border.Top = 10
	fig.Border.Bottom = 5

	// DPI for png
	fig.DPI = 300
}

// Apply nice style for the plot
func ApplyToPlot(p *hplot.Plot) {

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
	p.Legend.TextStyle.Font.Size = 12
	p.Legend.TextStyle.Color = SmoothBlack
	p.Legend.Padding = 0.1 * vg.Inch
	p.Legend.ThumbnailWidth = 25
}

// Apply a style to a ratio plot, from the top plot
func ApplyToRatioPlot(rp *hplot.RatioPlot, top *hplot.Plot) {

	// Bottom plot style
	ApplyToPlot(rp.Bottom)
	rp.Bottom.Y.Tick.Marker = hplot.Ticks{N: 3}
	rp.Bottom.Y.Tick.Label.Font.Size = 10
	rp.Bottom.X = top.X

	// To plot style
	rp.Top = top
	rp.Top.X.Tick.Label.Font.Size = 0
	rp.Top.X.Tick.Label.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
	rp.Top.X.Tick.LineStyle.Width = 0.5
	rp.Top.X.Tick.LineStyle.Color = color.NRGBA{R: 120, G: 120, B: 120, A: 255}
	rp.Top.X.Tick.Length = 5
	rp.Top.X.LineStyle.Width = 0.8
	rp.Top.X.LineStyle.Color = color.NRGBA{R: 120, G: 120, B: 120, A: 255}
	rp.Top.X.Label.Text = ""	
}

// Apply cirle markers
func SetCircleMarkersToHist(h *hplot.H1D) {

	// Dots as marker
	h.GlyphStyle = draw.GlyphStyle{
		Shape:  draw.CircleGlyph{},
		Color:  SmoothBlack,
		Radius: vg.Points(3)}

	// Y error bars
	if h.YErrs != nil {
		h.YErrs.LineStyle.Color = SmoothBlack
		h.YErrs.LineStyle.Width = 2.5
		h.YErrs.CapWidth = 7
	}
}

// Apply cirle markers
func SetCircleMarkersToScatter(s *hplot.S2D) {

	// Dots as marker
	s.GlyphStyle = draw.GlyphStyle{
		Shape:  draw.CircleGlyph{},
		Color:  SmoothBlack,
		Radius: vg.Points(3)}

	// Y error bars
	if s.YErrs != nil {
		s.YErrs.LineStyle.Color = SmoothBlack
		s.YErrs.LineStyle.Width = 2.5
		s.YErrs.CapWidth = 7
	}
}

// Apply nominal style for histogram representing data
func ApplyToDataHist(hData *hplot.H1D) {

	// Remove basic stat info
	hData.Infos.Style = hplot.HInfoNone

	// Remove line
	hData.LineStyle.Width = 0

	// Remove fill
	hData.FillColor = color.NRGBA{}

	// Apply circles as marker
	SetCircleMarkersToHist(hData)

}

// Apply nominal style for scatter2D representing data
func ApplyToDataS2D(sData *hplot.S2D) {

	// No line
	sData.LineStyle.Width = 0

	// Apply circles as marker // to be written for scatter2D
	SetCircleMarkersToScatter(sData)

}

// Helper to copy histogram style
func CopyStyleH1DtoH1D(dst, src *hplot.H1D) {
	dst.LineStyle = src.LineStyle
	dst.FillColor = src.FillColor
	dst.GlyphStyle = src.GlyphStyle
	if src.YErrs != nil {
		dst.YErrs.LineStyle = src.YErrs.LineStyle
		dst.YErrs.CapWidth = src.YErrs.CapWidth
	}
	if src.Band != nil && dst.Band != nil {
		dst.Band.FillColor = src.Band.FillColor
		dst.Band.LineStyle = src.Band.LineStyle
	}
}

func CopyStyleH1DtoS2D(dst *hplot.S2D, src *hplot.H1D) {
	dst.LineStyle = src.LineStyle
	dst.GlyphStyle = src.GlyphStyle
	if src.YErrs != nil {
		dst.YErrs.LineStyle = src.YErrs.LineStyle
		dst.YErrs.CapWidth = src.YErrs.CapWidth
	}
	if src.Band != nil && dst.Band != nil {
		dst.Band.FillColor = src.Band.FillColor
		dst.Band.LineStyle = src.Band.LineStyle
	}
}

// Helper function to change color opacity
func ChangeOpacity(c color.NRGBA, alpha uint8) color.NRGBA {
	return color.NRGBA{R: c.R, G: c.G, B: c.B, A: alpha}
}
