# Day 18: Settlers of The North Pole

On the outskirts of the North Pole base construction project, many Elves are
collecting lumber.

The lumber collection area is 50 acres by 50 acres; each acre can be either
__open ground__ (`.`), __trees__ (`|`), or a __lumberyard__ (`#`). You take a
scan of the area (your puzzle input).

Strange magic is at work here: each minute, the landscape looks entirely
different. In exactly __one minute__, an open acre can fill with trees, a wooded
acre can be converted to a lumberyard, or a lumberyard can be cleared to open
ground (the lumber having been sent to other projects).

The change to each acre is based entirely on __the contents of that acre__ as
well as __the number of open, wooded, or lumberyard acres adjacent to it__ at
the start of each minute. Here, "adjacent" means any of the eight acres
surrounding that acre. (Acres on the edges of the lumber collection area might
have fewer than eight adjacent acres; the missing acres aren't counted.)

In particular:
- An __open__ acre will become filled with __trees__ if __three or more__
adjacent acres contained trees. Otherwise, nothing happens.
- An acre filled with __trees__ will become a __lumberyard__ if __three or
more__ adjacent acres were lumberyards. Otherwise, nothing happens.
- An acre containing a __lumberyard__ will remain a __lumberyard__ if it was
adjacent to __at least one other lumberyard and at least one acre containing
trees__. Otherwise, it becomes __open__.

These changes happen across all acres __simultaneously__, each of them using the
state of all acres at the beginning of the minute and changing to their new form
by the end of that same minute. Changes that happen during the minute don't
affect each other.

For example, suppose the lumber collection area is instead only 10 by 10 acres
with this initial configuration:

```txt
Initial state:
.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.

After 1 minute:
.......##.
......|###
.|..|...#.
..|#||...#
..##||.|#|
...#||||..
||...|||..
|||||.||.|
||||||||||
....||..|.

After 2 minutes:
.......#..
......|#..
.|.|||....
..##|||..#
..###|||#|
...#|||||.
|||||||||.
||||||||||
||||||||||
.|||||||||

After 3 minutes:
.......#..
....|||#..
.|.||||...
..###|||.#
...##|||#|
.||##|||||
||||||||||
||||||||||
||||||||||
||||||||||

After 4 minutes:
.....|.#..
...||||#..
.|.#||||..
..###||||#
...###||#|
|||##|||||
||||||||||
||||||||||
||||||||||
||||||||||

After 5 minutes:
....|||#..
...||||#..
.|.##||||.
..####|||#
.|.###||#|
|||###||||
||||||||||
||||||||||
||||||||||
||||||||||

After 6 minutes:
...||||#..
...||||#..
.|.###|||.
..#.##|||#
|||#.##|#|
|||###||||
||||#|||||
||||||||||
||||||||||
||||||||||

After 7 minutes:
...||||#..
..||#|##..
.|.####||.
||#..##||#
||##.##|#|
|||####|||
|||###||||
||||||||||
||||||||||
||||||||||

After 8 minutes:
..||||##..
..|#####..
|||#####|.
||#...##|#
||##..###|
||##.###||
|||####|||
||||#|||||
||||||||||
||||||||||

After 9 minutes:
..||###...
.||#####..
||##...##.
||#....###
|##....##|
||##..###|
||######||
|||###||||
||||||||||
||||||||||

After 10 minutes:
.||##.....
||###.....
||##......
|##.....##
|##.....##
|##....##|
||##.####|
||#####|||
||||#|||||
||||||||||
```

After 10 minutes, there are `37` wooded acres and `31` lumberyards. Multiplying
the number of wooded acres by the number of lumberyards gives the total
__resource value__ after ten minutes: `37 * 31 = 1147`.

__What will the total resource value of the lumber collection area be after 10
minutes?__