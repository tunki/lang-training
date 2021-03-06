package main

import (
        "code.google.com/p/plotinum/plot"
        "code.google.com/p/plotinum/plotter"
        "code.google.com/p/plotinum/vg"
        "image/color"
        "math"
        "math/rand"
)

func main() {
        // Draw some random values from the standard
        // normal distribution.
        rand.Seed(int64(0))
        v := make(plotter.Values, 10000)
        for i := range v {
                v[i] = rand.NormFloat64()
        }

        // Make a plot and set its title.
        p, err := plot.New()
        if err != nil {
                panic(err)
        }
        p.Title.Text = "Histogram"

        // Create a histogram of our values drawn
        // from the standard normal.
        h, err := plotter.NewHist(v, 16)
        if err != nil {
                panic(err)
        }
        // Normalize the area under the histogram to
        // sum to one.
        h.Normalize(1)
        p.Add(h)

        // The normal distribution function
        norm := plotter.NewFunction(stdNorm)
        norm.Color = color.RGBA{R: 255, A: 255}
        norm.Width = vg.Points(2)
        p.Add(norm)

        // Save the plot to a PNG file.
        if err := p.Save(4, 4, "hist.png"); err != nil {
                panic(err)
        }
}

// stdNorm returns the probability of drawing a
// value from a standard normal distribution.
func stdNorm(x float64) float64 {
        const sigma = 1.0
        const mu = 0.0
        const root2π = 2.50662827459517818309
        return 1.0 / (sigma * root2π) * math.Exp(-((x-mu)*(x-mu))/(2*sigma*sigma))
}
