package charts

import "math"
import "github.com/keep94/gochart"

func Squares() {
    xs := gochart.NewInts(1, 1, 10)
    ys := xs.Apply(func(x int64) int64 { return x*x })
    gochart.NewChart(xs, ys).WriteTo(nil)
}

func SquareRoots() {
    xs := gochart.NewFloats(1.0, 1.0, 10)
    ys := xs.Apply(math.Sqrt)
    gochart.NewChart(xs, ys, gochart.YFormat("%.4f")).WriteTo(nil)
}
