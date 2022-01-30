package thefarm

import "fmt"

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()

	if cows == 0 {
		return 0, fmt.Errorf("division by zero")
	} else if cows < 0 {
		return 0, &SillyNephewError{cows: cows}
	}

	if fodder >= 0 {
		if err == ErrScaleMalfunction {
			return fodder * 2 / float64(cows), nil
		} else if err != nil {
			return 0 / float64(cows), fmt.Errorf("non-scale error")
		}
	} else if err != nil && err != ErrScaleMalfunction {
		return 0, fmt.Errorf("non-scale error")
	} else {
		return 0, fmt.Errorf("negative fodder")
	}

	return fodder / float64(cows), nil
}
