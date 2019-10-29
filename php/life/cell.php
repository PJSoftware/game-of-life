<?php

error_reporting(E_ALL);

class Cell
{
    private $current_state;
    private $future_state;

    function __construct($alive = false)
    {
        $this->current_state = $alive;
        $this->future_state = $this->current_state;
    }

    public function update_state($state)
    {
        $this->future_state = $state;
    }

    public function refresh()
    {
        $this->current_state = $this->future_state;
    }

    public function is_alive()
    {
        return $this->current_state;
    }

    public function to_string()
    {
        return $this->current_state ? "O" : ".";
    }

}