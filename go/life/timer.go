package life

import (
	"fmt"
	"time"
)

type Timer struct {
	msElapsed, totElapsed int64
	timerStart            time.Time
	timesCalled           int64
}

func (t *Timer) Start() {
	t.timerStart = time.Now()
}

func (t *Timer) Stop() {
	timerStop := time.Now()

	t.msElapsed = timerStop.Sub(t.timerStart).Milliseconds()
	t.timesCalled++
	t.totElapsed += t.msElapsed
}

func (t *Timer) Elapsed() string {
	return fmt.Sprintf("%.3f", float64(t.msElapsed)/1000.0)
}

func (t *Timer) AverageElapsed() string {
	var avg float64
	avg = 0
	if t.timesCalled > 0 {
		avg = float64(t.totElapsed / t.timesCalled)
	}
	return fmt.Sprintf("%.3f", avg/1000)
}

func (t *Timer) TotalElapsed() string {
	return fmt.Sprintf("%.3f", float64(t.totElapsed/1000))
}

func (t *Timer) ToString() string {
	return fmt.Sprintf("Loop %s (Avg %s)", t.Elapsed(), t.AverageElapsed())
}
