package main

import (
	"fmt"
	"time"
)

type Data struct {
	Clocks     []*Clock           `json:"clocks"`
	Clock      *Clock             `json:"clock"`
	Countdowns []*Countdown       `json:"countdowns"`
	Countdown  *Countdown         `json:"countdown"`
	Segmented  *SegmentedDuration `json:"segmented"`
}

type Clock struct {
	loc      *time.Location
	Tz       string `json:"tz"`
	Value    string `json:"value"`
	Selected bool   `json:"selected"`
}

type Countdown struct {
	Value    int64 `json:"value"`
	Selected bool  `json:"selected"`
}

var (
	Sec  int64 = 1
	Min  int64 = Sec * 60
	Hour int64 = Min * 60
)

func NewSegmentedDuration(dur int64) *SegmentedDuration {
	sd := SegmentedDuration{}

	sd.Sec = dur
	sd.Hour = sd.Sec / Hour
	sd.Sec -= sd.Hour * Hour
	sd.Min = sd.Sec / Min
	sd.Sec -= sd.Min * Min

	sd.Hours = fmt.Sprintf("%02d", sd.Hour)
	sd.Mins = fmt.Sprintf("%02d", sd.Min)
	sd.Secs = fmt.Sprintf("%02d", sd.Sec)

	return &sd
}

type SegmentedDuration struct {
	Hour  int64  `json:"hour"`
	Hours string `json:"hours"`
	Min   int64  `json:"min"`
	Mins  string `json:"mins"`
	Sec   int64  `json:"sec"`
	Secs  string `json:"secs"`
}

func (this *SegmentedDuration) Reset() {
	this.Hour = 0
	this.Min = 0
	this.Sec = 0
}

var (
	tz2Loc = map[string]string{
		"EST": "America/New_York",
		"MST": "America/Denver",
		"CST": "America/Chicago",
		"PST": "America/Los_Angeles",
		"IST": "Asia/Kolkata",
	}
)
