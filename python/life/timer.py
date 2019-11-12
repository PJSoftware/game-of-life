import time

class Timer:
    msElapsed = 0
    timesCalled = 0
    timerStart = 0
    totElapsed = 0

    def start(self):
        self.timerStart = time.time_ns() / 1000000.0

    def stop(self):
        timerStop = time.time_ns() / 1000000.0
        if self.timerStart != 0:
            self.msElapsed = timerStop - self.timerStart
            self.timerStart = 0
            self.timesCalled += 1
            self.totElapsed += self.msElapsed

    def elapsed(self):
        return "{:.3f}".format(self.msElapsed)

    def averageElapsed(self):
        average = 0 if self.timesCalled == 0 else self.totElapsed / self.timesCalled
        return "{:.3f}".format(average)

    def totalElapsed(self):
        return "{:.3f}".format(self.totElapsed)

    def toString(self):
        return "Loop " + self.elapsed() + "(Avg " + self.averageElapsed() + ")"
