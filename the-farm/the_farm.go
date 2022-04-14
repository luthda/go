package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
  cows int
}
func (e *SillyNephewError) Error() string {
  return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
  fodderAmount, err := weightFodder.FodderAmount()

  if cows == 0 {
    return 0.0, errors.New("division by zero")
  } else if cows < 0 {
    return 0.0, &SillyNephewError{ cows: cows }
  } else if fodderAmount >= 0 && err == ErrScaleMalfunction {
    return fodderAmount / float64(cows) * 2.0, nil 
  } else if err != nil && err != ErrScaleMalfunction {
    return 0, err
  } else if err == nil  && fodderAmount >= 0{
    return fodderAmount / float64(cows), nil
  } else if fodderAmount < 0 {
    return 0.0, errors.New("negative fodder")
  } else {
		return 0, err
  }
}