package scientific

import (
	"math/big"
)

// converts a big.Int to scientific notation if it is larger than 1,000,000
func BigIntToScientific(n *big.Int) string {
	// Threshold to compare the value with 1,000,000
	threshold := big.NewInt(1000000)

	// Convert big.Int to big.Float
	f := new(big.Float).SetInt(n)

	// Check if the number is greater than or equal to 1,000,000
	if n.Cmp(threshold) >= 0 {
		// Use big.Float.Text to format in scientific notation
		return f.Text('e', 3) // 'e' format for scientific notation with 3 significant digits
	}

	// If the number is less than 1,000,000, return it as a string
	return n.String()
}
