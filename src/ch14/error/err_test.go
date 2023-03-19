package err_test

import (
	"errors"
	"testing"
)

var MoreThan = errors.New("more than 100")
var LessThan = errors.New("less than 0")

func Getfibonaic(n int) ([]int, error) {
	//if n < 0 || n > 100 {
	//	return nil, errors.New("n should be in [1, 100]")
	//}

	if n < 0 {
		return nil, LessThan
	}

	if n > 100 {
		return nil, MoreThan
	}

	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}

	return fibList, nil
}

func TestGetFibonaic(t *testing.T) {
	//t.Log(Getfibonaic(10))
	//t.Log(Getfibonaic(-10))
	if v, err := Getfibonaic(-10); err == nil {
		t.Log(v)
	} else {
		if err == LessThan {
			// do something
			t.Error("less than 0!")
		} else if err == MoreThan {
			t.Error("more than 100!")
		}
	}
}
