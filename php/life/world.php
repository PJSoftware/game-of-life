<?php

error_reporting(E_ALL);

require_once 'cell.php';
require_once 'setup.php';

class World
{

    private $EMPTY_CELL;
    private $param;
    private $step = 0;
    private $grid = array();    // Implementing as hashed array provides more flexibility
    private $neighbours_at = array(
        array(-1, 1),  array(0, 1),  array(1, 1), // above
        array(-1, 0),  /*  self  */  array(1, 0), // beside
        array(-1, -1), array(0, -1), array(1, -1) // below
    );

    function __construct($predefined = false)
    {
        $this->EMPTY_CELL = new Cell;
        $this->param = new Setup;
        if ($predefined) {
            $this->predefined();
        } else {
            $this->prepopulate();
        }
    }

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

    private function prepopulate()
    {
        for ($x = 1; $x <= $this->param->width; $x++) {
            for ($y = 1; $y <= $this->param->height; $y++) {
                $alive = (rand(1, 100) <= $this->param->spawn_percent);
                $this->add_cell($x, $y, $alive);
            }
        }
    }

    private function predefined()   // One oscillator to test that we are working
    {
        for ($x = 1; $x <= $this->param->width; $x++) {
            for ($y = 1; $y <= $this->param->height; $y++) {
                $alive = false;
                if ($x==5 && ($y == 4 || $y == 5 || $y == 6)) {
                    $alive = true;
                }
                $this->add_cell($x, $y, $alive);
            }
        }
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

    private function add_cell($x, $y, $alive)
    {
        $this->grid[$this->grid_reference($x, $y)] = new Cell($alive);
    }

    private function grid_reference($x, $y)
    {
        $x = $this->wrap_coord($x, 1, $this->param->width, $this->param->wrapx);
        $y = $this->wrap_coord($y, 1, $this->param->height, $this->param->wrapy);
        return "$x|$y";
    }

    private function wrap_coord($val, $min, $max, $wrap_enabled) {
        if ($val < $min) {
            $val = $wrap_enabled ? $max - $val : $min - 1;
        } elseif ($val > $max) {
            $val = $wrap_enabled ? $val - $max : $max + 1;
        }
        return $val;
    }
}