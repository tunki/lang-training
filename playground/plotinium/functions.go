package main

import (
        "code.google.com/p/plotinum/plot"
        "code.google.com/p/plotinum/plotter"
        "code.google.com/p/plotinum/vg"
        "image/color"
        "math"
)

func main() {
        p, err := plot.New()
        if err != nil {
                panic(err)
        }
        p.Title.Text = "Functions"
        p.X.Label.Text = "X"
        p.Y.Label.Text = "Y"

        // A quadratic function x^2
        quad := plotter.NewFunction(func(x float64) float64 { return x * x })
        quad.Color = color.RGBA{B: 255, A: 255}

        // An exponential function 2^x
        exp := plotter.NewFunction(func(x float64) float64 { return math.Pow(2, x) })
        exp.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
        exp.Width = vg.Points(2)
        exp.Color = color.RGBA{G: 255, A: 255}

        // The sine function, shifted and scaled
        // to be nicely visible on the plot.
        sin := plotter.NewFunction(func(x float64) float64 { return 10*math.Sin(x) + 50 })
        sin.Dashes = []vg.Length{vg.Points(4), vg.Points(5)}
        sin.Width = vg.Points(4)
        sin.Color = color.RGBA{R: 255, A: 255}

        // Add the functions and their legend entries.
        p.Add(quad, exp, sin)
        p.Legend.Add("x^2", quad)
        p.Legend.Add("2^x", exp)
        p.Legend.Add("10*sin(x)+50", sin)
        p.Legend.ThumbnailWidth = vg.Inches(0.5)

        // Set the axis ranges.  Unlike other data sets,
        // functions don't set the axis ranges automatically
        // since functions don't necessarily have a
        // finite range of x and y values.
        p.X.Min = 0
        p.X.Max = 10
        p.Y.Min = 0
        p.Y.Max = 100

        // Save the plot to a PNG file.
        if err := p.Save(4, 4, "functions.png"); err != nil {
                panic(err)
        }
}
