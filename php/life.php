<?php

error_reporting(E_ALL);

/*
 *  Our main file should:
 *      Read arguments to get our starting point configuration
 *      Read INI file to initialise rules, world size and layout
 *      Initialise world per settings
 *      Populate world randomly or per specified starting layout
 *      Commence looping:
 *          Calculate new world position (with timing)
 *          Display new world position (with timing)
 *          Display timings (current and total)
 */

/*
 *  Required:
 *      World class to handle the World initialisation, calculation, display
 *          Setup class to handle all initial setup
 *          Cell class to handle individual cell operations (??)
 *      Timer class to handle timing calculations etc
 *      MAIN loop (does this need to be a class?)
 */

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

    echo "#$step | Calc " . $c_timer->to_string()." | Disp ".$d_timer->to_string()."\n";
}

it_lives();
