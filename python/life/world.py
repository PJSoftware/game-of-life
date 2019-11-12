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
        for x in range(1, self.param.width):
            for y in range(1, self.param.height):
                alive = True if x == 5 and (4 <= y <= 6) else False
                self._addCell(x, y, alive)

    def _prepopulate(self):
        for x in range(1, self.param.width):
            for y in range(1, self.param.height):
                alive = True if random.randint(0, 99) <= self.param.spawnPercent else False
                self._addCell(x, y, alive)
        
        
    def _addCell(self, x, y, alive):
        self.grid[self._gridReference(x, y)] = cell.Cell(alive)

    def _gridReference(self, x, y):
        x = self._wrapCoord(x, 1, self.param.width, self.param.wrapX)
        y = self._wrapCoord(y, 1, self.param.height, self.param.wrapY)
        return x + "|" + y
    
    def _wrapCoord(self, val, min, max, wrapEnabled):
        if val < min:
            val = max - val if wrapEnabled else min - 1
        elif val > max:
            val = val - max if wrapEnabled else max + 1
        return val

    def calculate(self):
        return


"""

    public function to_string()
    {
        $output = "";
        for ($y = 1; $y <= $this->param->height; $y++) {
            for ($x = 1; $x <= $this->param->width; $x++) {
                $output .= $this->cell_at($x, $y)->to_string();
            }
            $output .= "\n";
        }
        return $output;
    }

    public function calculate()
    {
        for ($y = 1; $y <= $this->param->height; $y++) {
            for ($x = 1; $x <= $this->param->width; $x++) {
                $cell = $this->cell_at($x, $y);
                $neighbours = $this->count_neighbours($x, $y);
                if ($cell->is_alive()) {
                    if (!array_key_exists("s$neighbours",$this->param->survive_values)) {
                        $cell->update_state(false);
                    }

                } else {
                    if (array_key_exists("b$neighbours",$this->param->birth_values)) {
                        $cell->update_state(true);
                    }
                }
            }
        }

        for ($x = 1; $x <= $this->param->width; $x++) {
            for ($y = 1; $y <= $this->param->height; $y++) {
                $this->cell_at($x, $y)->refresh();
            }
        }

        $this->step++;
    }

    public function step()
    {
        return $this->step;
    }

    private function count_neighbours($x, $y)
    {
        $total_neighbours = 0;
        foreach ($this->neighbours_at as $offset) {
            if ($this->cell_at($x + $offset[0], $y + $offset[1])->is_alive()) {
                $total_neighbours++;
            }
        }
        return $total_neighbours;
    }


    private function cell_at($x, $y)
    {
        $gridref = $this->grid_reference($x, $y);
        if (array_key_exists($gridref,$this->grid)) {
            return $this->grid[$this->grid_reference($x, $y)];
        } else {
            return $this->EMPTY_CELL;
        }
    }
    }
}
"""