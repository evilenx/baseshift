package computations

import (
	"math"
)

// Converts binary number to decimal.
func BinaryToDecimal(binary int64) int64 {
  binaryToArray := NumberArray(binary)
  decimal := int64(0)

  for i := 0; i < len(binaryToArray); i++ {
    element := binaryToArray[i]
    decimal = decimal + (element * int64(math.Pow(2, float64(i))))
  }

  return decimal
}

