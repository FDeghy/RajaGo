package raja

import "errors"

var (
	ErrGetTrains       = errors.New("get train list error")
	ErrGetTrainsDecode = errors.New("decode train list json error")
)
