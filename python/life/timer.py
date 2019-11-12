import time

class Timer:
    msElapsed = 0
    totElapsed = 0
    avgElapsed = 0
    timerStart = 0
    timesCalled = 0
    prevAvg = ""
    stableCount = 0

    def start(self):
        self.timerStart = time.time_ns() / 1000000.0

    def stop(self):
        timerStop = time.time_ns() / 1000000.0
        if self.timerStart != 0:
            self.msElapsed = timerStop - self.timerStart
            self.timesCalled += 1
            self.totElapsed += self.msElapsed

            self.prevAvg = "{:.3f}".format(self.avgElapsed)
            self.avgElapsed = self.totElapsed / self.timesCalled

    def elapsed(self):
        return "{:.3f}".format(self.msElapsed)

    def averageElapsed(self):
        return "{:.3f}".format(self.avgElapsed)

    def totalElapsed(self):
        return "{:.3f}".format(self.totElapsed)

    def isAverageStable(self, n):
        currAvg = self.averageElapsed()
        if currAvg == self.prevAvg:
            self.stableCount += 1
        else:
            self.stableCount = 0
        return self.stableCount >= n

    def __str__(self):
        return "Loop " + self.elapsed() + "(Avg " + self.averageElapsed() + ")"
