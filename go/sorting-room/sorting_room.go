package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %0.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %0.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if n, ok := fnb.(FancyNumber); ok {
		n, _ := strconv.Atoi(n.Value())
		return n
	}

	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %0.1f", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	if n, ok := i.(NumberBox); ok {
		return DescribeNumberBox(n)
	}

	if n, ok := i.(FancyNumberBox); ok {
		return DescribeFancyNumberBox(n)
	}

	if n, ok := i.(int); ok {
		return DescribeNumber(float64(n))
	}

	if n, ok := i.(float64); ok {
		return DescribeNumber(n)
	}

	return "Return to sender"
}
