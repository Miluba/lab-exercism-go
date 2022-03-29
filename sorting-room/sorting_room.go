package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	fs := fmt.Sprintf("This is the number %0.1f", f)
	return fs
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	nbs := fmt.Sprintf("This is a box containing the number %0.1f", float64(nb.Number()))
	return nbs
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
	if _, ok := fnb.(FancyNumber); !ok {
		return 0
	}
	v, err := strconv.Atoi(fnb.Value())
	if err != nil {
		return 0
	}
	return v
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	n := ExtractFancyNumber(fnb)
	dfnb := fmt.Sprintf("This is a fancy box containing the number %0.1f", float64(n))
	return dfnb
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch i.(type) {
	case int:
		v, _ := i.(int)
		return DescribeNumber(float64(v))
	case float64:
		v, _ := i.(float64)
		return DescribeNumber(v)
	case NumberBox:
		v, _ := i.(NumberBox)
		return DescribeNumberBox(v)
	case FancyNumberBox:
		v, _ := i.(FancyNumberBox)
		return DescribeFancyNumberBox(v)

	default:
		return fmt.Sprint("Return to sender")
	}
}
