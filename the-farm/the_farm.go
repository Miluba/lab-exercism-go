package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows string
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %s cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0, errors.New("division by zero")
	} else if cows < 0 {
		return 0, &SillyNephewError{
			cows: fmt.Sprint(cows),
		}
	}
	c, err := weightFodder.FodderAmount()
	if c < 0 && (err == nil || err.Error() == "sensor error") {
		return 0, errors.New("negative fodder")
	}
	if err == nil {
		return c / float64(cows), err
	}
	switch err.Error() {
	case "sensor error":
		if c >= 0 {
			return c * 2 / float64(cows), nil
		}
		fallthrough
	default:
		return 0, err
	}
}
