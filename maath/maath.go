package clockface

import (
	"math"
	"time"
)

// A Point represent two dimmensional Cartesian plane
type Point struct {
	X float64 `yml:"x"`
	Y float64 `yml:"y"`
}

func SecondHand(t time.Time) Point {
	return Point{X: 150, Y: 60}
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandInPoints(time time.Time) Point {
	return Point{0, -1}
}
