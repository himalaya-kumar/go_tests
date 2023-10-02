package clockface

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"time"
)

type (

	// A Point represent two dimmensional Cartesian plane
	Point struct {
		X float64 `yml:"x"`
		Y float64 `yml:"y"`
	}

	Svg struct {
		XMLName xml.Name `xml:"svg"`
		Text    string   `xml:",chardata"`
		Xmlns   string   `xml:"xmlns,attr"`
		Width   string   `xml:"width,attr"`
		Height  string   `xml:"height,attr"`
		ViewBox string   `xml:"viewBox,attr"`
		Version string   `xml:"version,attr"`
		Circle  Circle   `xml:"circle"`
		Line    []Line   `xml:"line"`
	}

	Circle struct {
		Cx float64 `xml:"cx,attr"`
		Cy float64 `xml:"cy,attr"`
		R  float64 `xml:"r,attr"`
	}

	Line struct {
		X1 float64 `xml:"x1,attr"`
		Y1 float64 `xml:"y1,attr"`
		X2 float64 `xml:"x2,attr"`
		Y2 float64 `xml:"y2,attr"`
	}
)

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

func SecondHand(t time.Time) Point {

	p := secondHandInPoints(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	return p
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandInPoints(t time.Time) Point {

	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	io.WriteString(w, svgEnd)
}
func secondHand(w io.Writer, t time.Time) {
	p := secondHandInPoints(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // Scale
	p = Point{p.X, -p.Y}                                      // Flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         //translate
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
