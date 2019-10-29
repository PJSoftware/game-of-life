<?php

error_reporting(E_ALL);

class Timer
{

    private $ms_elapsed = 0;
    private $times_called = 0;
    private $timer_start = 0;
    private $tot_elapsed = 0;

    public function start()
    {
        $this->timer_start = microtime(true);
    }

    public function stop()
    {
        $timer_stop = microtime(true);
        if ($this->timer_start != 0) {
            $this->ms_elapsed = ($timer_stop - $this->timer_start) * 1000;
            $this->timer_start = 0;
            $this->times_called++;
            $this->tot_elapsed += $this->ms_elapsed;
        }
    }

    public function elapsed()
    {
        return sprintf("%.3f", $this->ms_elapsed);
    }

    public function average_elapsed()
    {
        return sprintf("%.3f", $this->times_called == 0 ? 0 : $this->tot_elapsed / $this->times_called);
    }

    public function total_elapsed()
    {
        return sprintf("%.3f", $this->tot_elapsed);
    }

    public function to_string()
    {
        return "Loop " . $this->elapsed() . " (Avg " . $this->average_elapsed() . ")";
    }

}