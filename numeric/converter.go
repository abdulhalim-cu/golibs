package numeric

import (
	"fmt"
	"math"
)

// If we want to perform safe downsizing conversions we can always create
// suitable functions
// This function accepts an int argument and returns a uint8 and nil if the int is in
// range, or 0 and an error otherwise.

func Uint8FromInt(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", x)
}

// EqualFloat() function compares two float64 s to the given accuracy—or to
// the greatest accuracy the machine can achieve if a negative number (e.g., -1 )
// is passed as the limit. It relies on functions (and a constant) from the standard
// library’s math package

func EqualFloat(x, y, limit float64) bool {
	if limit <= 0.0 {
		limit = math.SmallestNonzeroFloat64
	}
	return math.Abs(x-y) <=
		(limit * math.Min(math.Abs(x), math.Abs(y)))
}

// Floating-point numbers can be converted to integers using the standard Go
// syntax (e.g., int(float) ), in which case the fractional part is simply discarded.
// Of course, if the floating-point value exceeds the range of the integer type
// converted to, the resultant integer will have an unpredictable value. We can
// address this problem using a safe conversion function.

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}
