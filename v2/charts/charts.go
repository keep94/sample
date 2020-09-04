package charts

import "github.com/keep94/gochart"

func Squares() {
    xs := gochart.NewInts(1, 1, 10)
    ys := xs.Apply(func(x int64) int64 { return x*x })
    gochart.NewChart(xs, ys).WriteTo(nil)
}
