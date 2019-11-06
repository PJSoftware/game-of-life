package life

import (
	"fmt"
	"time"
)

type Timer struct {
	msElapsed, totElapsed float64
	avgElapsed            float64
	timerStart            time.Time
	timesCalled           int64
	prevAvg               string
	stableCount           int
}

func (t *Timer) Start() {
	t.timerStart = time.Now()
}

func (t *Timer) Stop() {
	timerStop := time.Now()

	t.msElapsed = timerStop.Sub(t.timerStart).Seconds() * 1000
	t.timesCalled++
	t.totElapsed += t.msElapsed

	t.prevAvg = fmt.Sprintf("%.3f", t.avgElapsed)
	t.avgElapsed = t.totElapsed / float64(t.timesCalled)
}

func (t *Timer) Elapsed() string {
	return fmt.Sprintf("%.3f", t.msElapsed)
}

func (t *Timer) AverageElapsed() string {
	return fmt.Sprintf("%.3f", t.avgElapsed)
}

func (t *Timer) TotalElapsed() string {
	return fmt.Sprintf("%.3f", t.totElapsed)
}

func (t *Timer) ToString() string {
	return fmt.Sprintf("Loop %sms (Avg %sms)", t.Elapsed(), t.AverageElapsed())
}

func (t *Timer) IsAverageStable() bool {
	currAvg := fmt.Sprintf("%.3f", t.avgElapsed)
	if currAvg == t.prevAvg {
		t.stableCount++
	} else {
		t.stableCount = 0
	}
	return t.stableCount >= 20
}
