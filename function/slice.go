package function

import "golang.org/x/exp/constraints"

// CheckInSlice  check value in slice
func CheckInSlice[T constraints.Ordered](a T, s []T) bool {
	for _, val := range s {
		if a == val {
			return true
		}
	}
	return false
}

// DelOneInSlice  delete one element of slice  left->right
func DelOneInSlice[T constraints.Ordered](a T, old []T) (new []T) {
	for i, val := range old {
		if a == val {
			new = append(old[:i], old[i+1:]...)
			return
		}
	}
	return old
}
