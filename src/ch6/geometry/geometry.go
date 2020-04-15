package geometry

import "math"

// Point struct
type Point struct{ X, Y float64 }

// Distance traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// ScaleBy factor
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
