# vector2
## Go package for a basic 2d vector and calculation methods

### To use:

`Import "github.com/paulSB/vector2"`

### provides a 2 vector struct:

`type Vector struct {
	X float64
	Y float64
}`

### With the following methods:

```
// Calculate the magnitude/length of the vector v
func (v *Vector) Length() float64

// Removes length of vector v but retains ratio between X and Y
func (v *Vector) Normalize()

// Adds vector v2 to vector v
func (v *Vector) Add(v2 Vector)

// Deducts vector v2 from v
func (v *Vector) Subtract(v2 Vector)

// Multiplies vector v by a scalar
func (v *Vector) MultiplyByScalar(scalar float64)

// Divides vector v by a scalar
func (v *Vector) DivideByScalar(scalar float64)

// Calculates the Dot value of vector v
func (v *Vector) Dot(v2 Vector) float64

// Calculates the angle between vectors v and v2
func (v *Vector) AngleTo(v2 Vector) float64

// Calculates the distance from vector v to v2
func (v *Vector) DistanceTo(v2 Vector) float64

// Returns the rotation required for an image at vector v to "point towards" vector v2
func (v *Vector) PointTowards (v2 Vector) float64
```

Author Paul Brace
July 2024

Please message me with any questions or suggestions/corrections.
