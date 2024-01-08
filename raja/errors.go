package raja

import "errors"

var (
	ErrGetTrains       = errors.New("get train list error")
	ErrGetTrainsDecode = errors.New("decode train list json error")
	ErrTrainsNotFound  = errors.New("trains not found")
	ErrBadStatus       = errors.New("unknown status code")
)
