package life

import (
	"fmt"
	"time"
)

// Timer object tracks elapsed and average times in milliseconds
type Timer struct {
	msElapsed, totElapsed float64
	avgElapsed            float64
	timerStart            time.Time
	timesCalled           int64
	prevAvg               string
	stableCount           int
}

// Start must always be called before Stop; initialises Timer
func (t *Timer) Start() {
	t.timerStart = time.Now()
}

// Stop is called once code to be timed is completed
func (t *Timer) Stop() {
	timerStop := time.Now()

	t.msElapsed = timerStop.Sub(t.timerStart).Seconds() * 1000
	t.timesCalled++
	t.totElapsed += t.msElapsed

	t.prevAvg = fmt.Sprintf("%.3f", t.avgElapsed)
	t.avgElapsed = t.totElapsed / float64(t.timesCalled)
}

// Elapsed returns elapsed milliseconds between Start() and Stop()
func (t *Timer) Elapsed() string {
	return fmt.Sprintf("%.3f", t.msElapsed)
}

// AverageElapsed returns average duration over all timings
func (t *Timer) AverageElapsed() string {
	return fmt.Sprintf("%.3f", t.avgElapsed)
}

// TotalElapsed returns total duration of all timings
func (t *Timer) TotalElapsed() string {
	return fmt.Sprintf("%.3f", t.totElapsed)
}

// String provides automatic stringification for Timer
func (t *Timer) String() string {
	return fmt.Sprintf("Loop %sms (Avg %sms)", t.Elapsed(), t.AverageElapsed())
}

// IsAverageStable returns true when AverageElapsed() remains constant n times
func (t *Timer) IsAverageStable(n int) bool {
	currAvg := fmt.Sprintf("%.3f", t.avgElapsed)
	if currAvg == t.prevAvg {
		t.stableCount++
	} else {
		t.stableCount = 0
	}
	return t.stableCount >= n
}
