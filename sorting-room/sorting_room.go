package sorting

import (
	"fmt"
	"strconv"
)

func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
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

func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
		case FancyNumber:
			i, _ := strconv.Atoi(fnb.Value())
			return i
		default:
			return 0
	}
}

func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	fn := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(fn))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch nb := i.(type) { 
		case int:
			return DescribeNumber(float64(nb))
		case float64:
			return DescribeNumber(nb)
		case NumberBox:
			return DescribeNumberBox(nb)
		case FancyNumberBox:
			return DescribeFancyNumberBox(nb)
		default:
			return "Return to sender"
	}
}
