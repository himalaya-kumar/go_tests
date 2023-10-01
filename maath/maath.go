package clockface

import "time"

// A Point represent two dimmensional Cartesian plane
type Point struct {
	X float64 `yml:"x"`
	Y float64 `yml:"y"`
}

func SecondHand(t time.Time) Point {
	return Point{X: 150, Y: 60}
}
