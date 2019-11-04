package life

import (
	"fmt"
	"time"
)

type Timer struct {
	msElapsed, totElapsed float64
	timerStart            time.Time
	timesCalled           int64
	lastAvg				  string
	stableCount			  int
}

func (t *Timer) Start() {
	t.timerStart = time.Now()
}

func (t *Timer) Stop() {
	timerStop := time.Now()

	t.msElapsed = timerStop.Sub(t.timerStart).Seconds()*1000
	t.timesCalled++
	t.totElapsed += t.msElapsed
}

func (t *Timer) Elapsed() string {
	return fmt.Sprintf("%.3f", float64(t.msElapsed))
}

func (t *Timer) AverageElapsed() string {
	avgString := "-----"
	if t.timesCalled > 0 {
		avg := t.totElapsed / float64(t.timesCalled)
		avgString = fmt.Sprintf("%.3f", avg)
	}
	if avgString == t.lastAvg {
		t.stableCount++
	} else {
		fmt.Printf("Avg(%s) vs Last(%s) %d\n", avgString, t.lastAvg, t.timesCalled)
		//t.stableCount = 0
		t.lastAvg = avgString
	}
	return t.lastAvg
}

func (t *Timer) TotalElapsed() string {
	return fmt.Sprintf("%.3f", float64(t.totElapsed))
}

func (t *Timer) ToString() string {
	return fmt.Sprintf("Loop %sms (Avg %sms)", t.Elapsed(), t.AverageElapsed())
}

func (t *Timer) IsAverageStable() bool {
	fmt.Printf("%d\n", t.stableCount)
	return (t.stableCount >= 10)
}
