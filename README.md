# Conway's Game of Life Simulation
In 1985 I got my very first computer: a VZ-200 with a massive 8kb of RAM. I played a few games on it, messed around with BASIC for a while--and then dove in to teaching myself Assembler with a book-for-kids targeted at the ZX81 Spectrum (both machines shared the Z80A Processor.) The first program I ever wrote from scratch was an implementation of [Conway's Game of Life](http://en.wikipedia.org/wiki/Conway's_Game_of_Life), in Z80A Assembly Language, running on the 128x64 graphics mode of that clunky little machine.

Fast forward 34 years, and I found myself wanting a small programming project to brush up on my PHP. I was facing the possibility of a Tech Challenge in an upcoming job application, so I wanted the practice. Google showed me a couple of people suggesting using Conway's Game of Life as a Tech Challenge, and so I decided that I should try my hand at implementing it in PHP--and then translating it to Go while simultaneously learning that language.

Another Google search led me to [KieranP's](https://github.com/KieranP/) [Game-Of-Life-Implementations](https://github.com/KieranP/Game-Of-Life-Implementations) here on GitHub. I have not copied Kieran's code, but I *did* read through it before beginning, and no doubt borrowed a couple of design decisions therefrom. (Timing the calculation and display steps, for instance, and implementing the game grid as a one-dimensional array rather than the more obvious two-dimensional array!) However, I opted to implement a wrapped grid rather than non-wrapped; I also decided to store the dimensions, parameters, rules etc of the simulation in one single INI file which could be read by all versions of the code.

And, of course, I used a 128x64 grid, for old times' sake.

# Implementation Features

## OOP
Where it makes sense I have written Object-Oriented code. It just seems to make more sense--although maybe I'll change my mind when I run into a language which does not support OOP.

## INI file initialisation
One common ini file provides the configuration for all implementations. This includes at least one \[world-xxxxx\] section which defines the board size and configuration for the simulation, and a \[rules\] section which lists a number of different birth/survival "rulesets" that are compatible with the Game of Life automata processing rules.

If the INI file is not available--or if another option is not specified--the program should default to the following settings:
    * 128 x 64 board resolution, wrapped in both directions.
    * Initial spawn percentage of 20%.
    * A "square" grid (no other option is implemented yet.)

## Output to text console
At this point, printing to STDOUT is our only option--and the Windows CMD terminal I'm using does not give simple access to cursor positioning, colour output, or even a nice way to clear the screen, so for now I'm simply dumping to screen and allowing it to scroll.

## Calculation and Display timed independently
The code I looked at by *KieranP* timed both the calculation procedure and the display procedure. This seemed like a good idea so I have done the same in my code.

## Logging of average timing
Speaking of *KieranP*, their *README.md* file had a list of timings to compare speeds of the different language implementations. I intend to do something similar, and so to enable this, my code writes these values to a log file.

My intent (not yet implemented) is to produce something like a Perl script to read all the log files and tabulate the results automatically. I've yet to figure out all the details of how this will actually work, but it will probably append the results to the end of *this* file. (Perhaps I'll run it once for every new implementation added, for the final commit before merging back to Master.)

## Main loop continues until average timing stabilises
Initially, my code looped indefinitely. However, in order to actually automate the logging, I needed to determine an endpoint.

Rather than simply pick a number, I decided instead to run a minimum number of times (1000, as a nice round number) and then start watching the calculated *average* Calculation Timing value. Once it had remained unchanged through *n* (20 seems to work nicely) iterations, the loop (and the program) would stop, allowing the results to be logged.

## Logging of errors
**PHP** uses *exceptions*, and hence so did my PHP implementation.

**Go** does not understand *exceptions*: the Go way is to pass errors back out of functions, having them bubble up through the codebase until something either responds to it or ignores it. My response to this--especially since most of my actual error detection was around INI file processing (ie, "user input")--was to look into writing such anomalies to a log file and fall back to default values.

Which means I should probably look into adding that response to my PHP code too.

# TODO

## Improved text output
Where possible, it would be nice to be able to reset our cursor position before each pass, so we are overwriting rather than scrolling.

## Colour output
I have also considered adding more states than simply "Alive" or "Dead"; specifically, "Just Born" and "Just Died". I figured these could be different colurs, such as Green and Red.

If I ever get colour output working in my current environment, I'll look at implementing this.

## Graphical output
Likewise for graphical output.

Go in particular does not really provide any GUI capability in its core libraries, so for now this can remain on the wish list.

## Tabulate speeds
See discussion above under *Logging of average timing*. Note that this will be an implementation-independent task.

# Feature Tabulation

| Feature | PHP | Go | Python |
| --- | --- | --- | --- |
| OOP | Yes | Yes | Yes |
| INI File | Yes | Yes | No |
| Text Out | Yes | Yes | No |
| Timing | Yes | Yes | Yes |
| Times Logged | No | Yes | No |
| Auto Loop Break | No | Yes | Yes | 
| Errors Logged | No | Yes | No |
| Improved Output | No | No | No |
| Colour Output | No | No | No |
| Graphical Output | No | No | No |

# Language Speed Comparison
Not yet implemented.
