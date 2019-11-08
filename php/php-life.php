<?php

error_reporting(E_ALL);

require_once 'life/world.php';
require_once 'life/timer.php';

$c_timer = new Timer;
$d_timer = new Timer;

function it_lives()
{
    global $c_timer, $d_timer;

    $world = new World;
    echo $world->to_string();
    display_timings($world->step());

     while(true) {
        $c_timer->start();
        $world->calculate();
        $c_timer->stop();

        $d_timer->start();
        echo $world->to_string();
        $d_timer->stop();

        display_timings($world->step());
    }
}

function display_timings($step)
{
    global $c_timer, $d_timer;
    echo "PHP: #$step | Calc " . $c_timer->to_string()." | Disp ".$d_timer->to_string()."\n";
}

it_lives();
