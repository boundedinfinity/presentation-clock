package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx               context.Context
	clockTicker       *time.Ticker
	wg                *sync.WaitGroup
	timeFormat        string
	clockAll          []*Clock
	clockCurrent      *Clock
	clockSelectCh     chan (string)
	countdownAll      []*Countdown
	countdownCurrent  *Countdown
	countdownSelectCh chan (int64)
	countdownValue    int64
	debug             bool
}

func newElemet[T any](options []T) Item[T] {
	return Item[T]{
		Options: options,
		Ch:      make(chan T),
	}
}

type Item[T any] struct {
	Options []T
	Current T
	Ch      chan (T)
}

// NewApp creates a new App application struct
func NewApp() *App {
	clockTz := []string{"EST", "CST", "MST", "PST", "IST"}
	countdowns := []int64{1, 10, 15, 20, 30, 45, 60, 90}
	ticker := time.NewTicker(1)
	ticker.Stop()

	app := &App{
		debug: true,
		// https://gobyexample.com/tickers
		wg: &sync.WaitGroup{},
		// https://pkg.go.dev/time#example-Time.Format
		clockTicker:       time.NewTicker(time.Second),
		timeFormat:        "03:04:05 PM",
		clockAll:          []*Clock{},
		clockSelectCh:     make(chan string),
		countdownSelectCh: make(chan int64),
		countdownAll:      []*Countdown{},
		countdownValue:    0,
	}

	for _, backin := range countdowns {
		app.countdownAll = append(app.countdownAll, &Countdown{
			Value:    backin,
			Selected: false,
		})
	}

	for _, tz := range clockTz {
		if id, ok := tz2Loc[tz]; !ok {
			panic(fmt.Sprintf("unable to find location for clock %s", tz))
		} else {
			if loc, err := time.LoadLocation(id); err != nil {
				panic(fmt.Sprintf("unable to find location for clock %s", tz))
			} else {
				app.clockAll = append(app.clockAll, &Clock{
					loc:      loc,
					Tz:       tz,
					Selected: false,
				})
			}
		}
	}

	return app
}

func (this *App) startup(ctx context.Context) {
	this.ctx = ctx
	this.wg.Add(1)

	go func() {
		defer this.wg.Done()

		for {
			select {
			case countdown := <-this.countdownSelectCh:
				if this.debug {
					fmt.Printf("countdown received: %d\n", countdown)
				}

				this.countdownValue = 0

				if this.countdownCurrent != nil && this.countdownCurrent.Value == countdown {
					this.countdownCurrent = nil
				} else {
					this.countdownCurrent = nil
					for _, item := range this.countdownAll {
						if item.Value == countdown {
							item.Selected = true
							this.countdownCurrent = item
							this.countdownValue = countdown * Min
						} else {
							item.Selected = false
						}
					}
				}
			case tz := <-this.clockSelectCh:
				if this.debug {
					fmt.Printf("clock received: %s\n", tz)
				}

				if this.clockCurrent != nil && this.clockCurrent.Tz == tz {
					this.clockCurrent = nil
				} else {
					this.clockCurrent = nil
					for _, item := range this.clockAll {
						if item.Tz == tz {
							item.Selected = true
							this.clockCurrent = item
						} else {
							item.Selected = false
						}
					}
				}

			case <-this.ctx.Done():
				this.clockTicker.Stop()
				return
			case <-this.clockTicker.C:
				if current, err := this.GetCurrentTimes(); err != nil {
					panic(err)
				} else {
					// godump.Dump(current)
					runtime.EventsEmit(this.ctx, "data", current)
				}
			}
		}
	}()
}

func (this *App) shutdown() {
	this.wg.Wait()
}

func (this *App) SelectClock(name string) {
	go func() { this.clockSelectCh <- name }()
}

func (this *App) SelectCountdown(backin int64) {
	go func() { this.countdownSelectCh <- backin }()
}

func (this *App) GetCurrentTimes() (Data, error) {
	data := Data{
		Clocks:     this.clockAll,
		Clock:      this.clockCurrent,
		Countdowns: this.countdownAll,
		Countdown:  this.countdownCurrent,
	}

	now := time.Now()
	for _, clock := range data.Clocks {
		clock.Value = now.In(clock.loc).Format(this.timeFormat)
	}
	if this.countdownValue > 0 {
		this.countdownValue -= 1
		data.Segmented = NewSegmentedDuration(this.countdownValue)
	}

	// if this.debug {
	// 	godump.Dump(data.Segmented)
	// }

	return data, nil
}
