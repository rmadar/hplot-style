package style_test

import (
	"fmt"
	"log"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"

	"gonum.org/v1/plot/vg"

	"go-hep.org/x/hep/hbook"
	"go-hep.org/x/hep/hplot"

	"github.com/rmadar/hplot-style/style"
)

func ExampleSimePlot() {

	// Create a very wide normal distribution
	dist := distuv.Normal{
		Mu:    0,
		Sigma: 5,
		Src:   rand.New(rand.NewSource(0)),
	}

	// Histogram filling
	h := hbook.NewH1D(20, -5, 5)
	for i := 0; i < 50000; i++ {
		v := dist.Rand()
		h.Fill(v, 1)
	}

	// Histogram scaling
	h.Scale(1.0 / h.Integral())

	// Making the plot
	p := hplot.New()
	style.ApplyToPlot(p)
	hist := hplot.NewH1D(h)
	p.Add(hist)
	p.Y.Min = 0

	// Saving the result
	if err := p.Save(6*vg.Inch, -1, "example_plot.png"); err != nil {
		log.Fatalf("error saving plot: %v\n", err)
	} else {
		fmt.Println("Plot saved.")
	}

	// Output:
	// Plot saved.
}
