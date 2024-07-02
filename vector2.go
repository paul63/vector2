// Basic 2d vector struct and calculations for Go language
// Author Paul Brace
// July 2024

package vector2

import 	"math"

// Struct for a 2d vector with associated methods
type Vector struct {
	X float64
	Y float64
}

// Calculate the magnitude/length of the vector v
func (v *Vector) Length() float64{
	length := math.Sqrt((v.X * v.X) + (v.Y * v.Y))
	return length
}

// Removes length of vector v but retains ratio between X and Y
func (v *Vector) Normalize() {
	// calculate vector length
	length := v.Length()
	v.DivideByScalar(length)
}

// Adds vector v2 to vector v
func (v *Vector) Add(v2 Vector) {
	v.X += v2.X
	v.Y += v2.Y
}

// Deducts vector v2 from v
func (v *Vector) Subtract(v2 Vector) {
	v.X -= v2.X
	v.Y -= v2.Y
}

// Multiplies vector v by a scalar
func (v *Vector) MultiplyByScalar(scalar float64) {
	v.X *= scalar
	v.Y *= scalar
}

// Divides vector v by a scalar
func (v *Vector) DivideByScalar(scalar float64) {
	v.X /= scalar
	v.Y /= scalar
}

// Calculates the Dot value of vector v
func (v *Vector) Dot(v2 Vector) float64{
	return v.X * v2.X + v.Y * v2.Y
}

// Calculates the angle between vectors v and v2
func (v *Vector) AngleTo(v2 Vector) float64 {
	angle := math.Acos(v.Dot(v2) / (v.Length() * v2.Length()))
	return angle
}

// Calculates the distance from vector v to v2
func (v *Vector) DistanceTo(v2 Vector) float64 {
	angle := math.Sqrt(math.Pow(v2.X - v.X, 2) + math.Pow(v2.Y - v.Y, 2))
	return angle
}

// Returns the rotation required for an image at vector v to "point towards" vector v2
func (v *Vector) PointTowards (v2 Vector) float64 {
	dx := v.X - v2.X
	dy := v.Y - v2.Y
	return math.Atan2(-dx, dy)
}