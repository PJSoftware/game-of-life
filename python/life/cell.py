class Cell:
    
    def __init__(self, alive = False):
        self.currentState = alive
        self.futureState = alive

    def updateState(self, state):
        self.futureState = state

    def refresh(self):
        self.currentState = self.futureState

    def isAlive(self):
        return self.currentState

    def __str__(self):
        return "O" if self.currentState else "."
