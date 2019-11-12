from life import world
from life import timer

cTimer = timer.Timer()
dTimer = timer.Timer()

def itLives():
    myWorld = world.World()
    print(myWorld)
    displayTimings(myWorld.step)

    while myWorld.step < 1000 or not cTimer.isAverageStable(20):
        cTimer.start()
        myWorld.calculate()
        cTimer.stop()

        dTimer.start()
        print(myWorld, end='')
        dTimer.stop()

        displayTimings(myWorld.step)

def displayTimings(step):
    print(f"Python: #{step} | Calc " + str(cTimer) + " | Disp " + str(dTimer))

itLives()
