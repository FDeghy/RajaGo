package raja

import (
	"net/http"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

const (
	BASE_URL    = "https://www.raja.ir"
	SERVICE_URL = "https://hostservice.raja.ir"
	USER_AGENT  = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"
)

var (
	Client = http.Client{
		Timeout: 30 * time.Second,
	}
	Stations []Station
)

type Station struct {
	Code        int    `json:"Code"`
	Name        string `json:"Name"`
	EnglishName string `json:"EnglishName"`
	TelCode     string `json:"TelCode"`
	Group       string `json:"Group"`
	IsCarTicket bool   `json:"IsCarTicket"`
	IsExclusion bool   `json:"IsExclusion"`
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
