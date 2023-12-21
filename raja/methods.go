package raja

import (
	"fmt"

	ptime "github.com/yaa110/go-persian-calendar"
)

func NewTrainInfo(src string, dst string, date ptime.Time) (TrainInfo, error) {
	ti := new(TrainInfo)
	stSrc, err := FindStation(src)
	if err != nil {
		return *ti, err
	}
	stDst, err := FindStation(dst)
	if err != nil {
		return *ti, err
	}
	ti.Source = stSrc
	ti.Destination = stDst
	ti.ShamsiDate = date
	return *ti, nil
}

func (ti TrainInfo) encode() string {
	return fmt.Sprintf(
		"%d-%d-Family-1-%s--1-false-0-0-L1",
		ti.Source.Code,
		ti.Destination.Code,
		ti.ShamsiDate.Format("yyyyMMdd"),
	)
}
