package raja

import (
	"net/http"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

type Query string

type Stations []Station
type Station struct {
	Id          int    `json:"Code"`
	PersianName string `json:"Name"`
	EnglishName string `json:"EnglishName"`
}

type xrs struct {
	As          string        `json:"As"`
	Message     string        `json:"Message"`
	LinkMessage string        `json:"LinkMessage"`
	Slider      []interface{} `json:"slider"`
	Key         string        `json:"key"`
}

type TrainList struct {
	Trains       []GoTrains `json:"GoTrains"`
	ReturnTrains any        `json:"ReturnTrains"`
}
type GoTrains struct {
	ArrivalDate               time.Time `json:"ArrivalDate"`
	FullPrice                 float64   `json:"FullPrice"`
	FromStation               int       `json:"FromStation"`
	ToStation                 int       `json:"ToStation"`
	RetStatus                 int       `json:"RetStatus"`
	Remain                    int       `json:"Remain"`
	TrainNumber               int       `json:"TrainNumber"`
	WagonType                 int       `json:"WagonType"`
	WagonName                 string    `json:"WagonName"`
	PathCode                  int       `json:"PathCode"`
	CircularPeriod            int       `json:"CircularPeriod"`
	MoveDate                  string    `json:"MoveDate"`
	ShamsiMoveDate            string    `json:"ShamsiMoveDate"`
	ExitDate                  string    `json:"ExitDate"`
	ExitTime                  string    `json:"ExitTime"`
	TimeOfArrival             string    `json:"TimeOfArrival"`
	ShamsiDateofArrival       string    `json:"ShamsiDateofArrival"`
	Counting                  float64   `json:"Counting"`
	SoldCount                 int       `json:"SoldCount"`
	Degree                    int       `json:"degree"`
	AvaliableSellCount        int       `json:"AvaliableSellCount"`
	Cost                      float64   `json:"Cost"`
	CostDisplay               string    `json:"CostDisplay"`
	CompartmentCapicity       int       `json:"CompartmentCapicity"`
	IsCompartment             int       `json:"IsCompartment"`
	WagonInfo                 string    `json:"WagonInfo"`
	TimeFilter                string    `json:"TimeFilter"`
	WagonFilter               string    `json:"WagonFilter"`
	CircularNumberSerial      int       `json:"CircularNumberSerial"`
	CountingAll               int       `json:"CountingAll"`
	RateCode                  int       `json:"RateCode"`
	AirConditioning           bool      `json:"AirConditioning"`
	AirConditioningDisply     string    `json:"AirConditioningDisply"`
	Media                     bool      `json:"Media"`
	MediaDisply               string    `json:"MediaDisply"`
	RowID                     int       `json:"RowId"`
	ButtonRowType             any       `json:"ButtonRowType"`
	RationCode                int       `json:"RationCode"`
	Soldcounting              int       `json:"soldcounting"`
	SeatType                  int       `json:"SeatType"`
	Owner                     int       `json:"Owner"`
	Ischarter                 bool      `json:"ischarter"`
	AxleCode                  int       `json:"AxleCode"`
	ShowFullPrice             bool      `json:"ShowFullPrice"`
	Tooltip                   string    `json:"Tooltip"`
	SexCode                   int       `json:"SexCode"`
	CompanyName               string    `json:"CompanyName"`
	ShamsiDate                string    `json:"ShamsiDate"`
	DayName                   string    `json:"DayName"`
	TimeDuration              int       `json:"TimeDuration"`
	ExitDateTime              string    `json:"ExitDateTime"`
	RequestIsExclusive        bool      `json:"RequestIsExclusive"`
	RequestNumberOfPassengers int       `json:"RequestNumberOfPassengers"`
	IsDisabled                bool      `json:"IsDisabled"`
	BackgroundColor           string    `json:"BackgroundColor"`
	DiscountPercent           string    `json:"DiscountPercent"`
}

type TrainInfo struct {
	Source      Station
	Destination Station
	ShamsiDate  ptime.Time
}

type GetTrainListOpt struct {
	HttpClient *http.Client
	ApiKey     string
}
