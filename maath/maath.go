package maath

import "time"

// A Point represent two dimmensional Cartesian plane
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func SecondHand(t time.Time) Point {
	return Point{}
}
