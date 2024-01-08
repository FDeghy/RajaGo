package raja

import (
	"fmt"
)

func (ti TrainInfo) encode() string {
	return fmt.Sprintf(
		"%d-%d-Family-1-%s--1-false-0-0-L1",
		ti.Source.Id,
		ti.Destination.Id,
		ti.ShamsiDate.Format("yyyyMMdd"),
	)
}

func (q Query) String() string {
	return string(q)
}
