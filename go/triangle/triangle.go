// package triangle contains methods for determining if a triangle is equilateral, isosceles, or scalene
package triangle

import (
    "math"
    "sort"
)

// Kind of triangle
type Kind int

// Triangle kinds
const (
    NaT = iota // not a triangle
    Equ        // equilateral
    Iso        // isosceles
    Sca        // scalene
)

// KindFromSides checks if a triangle is equilateral, isosceles, or scalene
func KindFromSides(sides []float64) (k Kind) {
    k = Sca

    a := sides[0]
    b := sides[1]
    c := sides[2]

    if !IsTriangle(sides) {
        k = NaT
    } else if a == b && a == c {
        k = Equ
    } else if a == b || a == c || b == c {
        k = Iso
    }
    return k
}

// IsTriangle checks for triangle inequality
func IsTriangle(sides []float64) bool {
    // Check if side values are not NaN or Infinity
    for _, value := range sides {
        if math.IsNaN(value) || math.IsInf(value, 0) {
            return false
        }
    }

    sort.Float64s(sides)

    a, b, c := sides[0], sides[1], sides[2]

    // Check if sides are not negative values
    if !(a > float64(0) && b > float64(0) && c > float64(0)) {
        return false
    }

    if a + b < c || b + c < a || c + a < b {
        return false
    }

    return true
}
