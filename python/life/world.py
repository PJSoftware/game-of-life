from life import param
from life import cell
import random

class World:
    neighboursAt = [
        [-1, 1], [0, 1], [1, 1],    # above
        [-1, 0],         [1, 0],    # beside
        [-1, -1], [0, -1], [1, -1]  # below
    ]

    def __init__(self, useTestWorld = False):
        self.EMPTY_CELL = cell.Cell()
        self.step = 0
        self.grid = {}
        self.param = param.Param()
        if useTestWorld:
            self._predefined()
        else:
            self._prepopulate()

    def _predefined(self):
        for x in range(1, self.param.width+1):      # range(1,j) steps from i to j-1
            for y in range(1, self.param.height+1):
                alive = True if x == 5 and (4 <= y <= 6) else False
                self._addCell(x, y, alive)

    def _prepopulate(self):
        for x in range(1, self.param.width+1):
            for y in range(1, self.param.height+1):
                alive = True if random.randint(0, 99) <= self.param.spawnPercent else False
                self._addCell(x, y, alive)
        
        
    def _addCell(self, x, y, alive):
        self.grid[self._gridReference(x, y)] = cell.Cell(alive)

    def _gridReference(self, x, y):
        x = self._wrapCoord(x, 1, self.param.width, self.param.wrapX)
        y = self._wrapCoord(y, 1, self.param.height, self.param.wrapY)
        return str(x) + "|" + str(y)
    
    def _wrapCoord(self, val, min, max, wrapEnabled):
        if val < min:
            val = max - val if wrapEnabled else min - 1
        elif val > max:
            val = val - max if wrapEnabled else max + 1
        return val

    def calculate(self):
        for y in range(1, self.param.height+1):
            for x in range(1, self.param.width+1):
                myCell = self._cellAt(x, y)
                neighbours = self._countNeighbours(x, y)
                if myCell.isAlive():
                    if not f"s{neighbours}" in self.param.surviveValues:
                        myCell.updateState(False)
                else:
                    if f"b{neighbours}" in self.param.birthValues:
                        myCell.updateState(True)

        for y in range(1, self.param.height+1):
            for x in range(1, self.param.width+1):
                self._cellAt(x, y).refresh()

        self.step += 1

    def _cellAt(self, x, y):
        gridRef = self._gridReference(x, y)
        return self.grid[gridRef] if gridRef in self.grid else self.EMPTY_CELL

    def _countNeighbours(self, x, y):
        totalNeighbours = 0
        for offset in self.neighboursAt:
            if self._cellAt(x + offset[0], y + offset[1]).isAlive():
                totalNeighbours += 1
        return totalNeighbours

    def __str__(self):
        output = ""
        for y in range(1, self.param.height+1):
            for x in range(1, self.param.width+1):
                output += str(self._cellAt(x, y))
            output += "\n"
        return output
