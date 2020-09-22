package geometry

import (
	"errors"
)

func CubeVolume(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	} else {
		return 0, errors.New("Zero length is not allowed")
	}
}

//vMAJOR.MINOR.PATCH -> this is the numbering scheme used go modules
