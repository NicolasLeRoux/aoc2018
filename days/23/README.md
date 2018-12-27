# Day 23: Experimental Emergency Teleportation

Using your torch to search the darkness of the rocky cavern, you finally locate the man's friend: a small __reindeer__.

You're not sure how it got so far in this cave. It looks sick - too sick to walk - and too heavy for you to carry all the way back. Sleighs won't be invented for another 1500 years, of course.

The only option is __experimental emergency teleportation__.

You hit the "experimental emergency teleportation" button on the device and push __I accept the risk__ on no fewer than 18 different warning messages. Immediately, the device deploys hundreds of tiny __nanobots__ which fly around the cavern, apparently assembling themselves into a very specific __formation__. The device lists the `X,Y,Z` position (`pos`) for each nanobot as well as its __signal radius__ (`r`) on its tiny screen (your puzzle input).

Each nanobot can transmit signals to any integer coordinate which is a distance away from it __less than or equal to__ its signal radius (as measured by [Manhattan distance](https://en.wikipedia.org/wiki/Taxicab_geometry)). Coordinates a distance away of less than or equal to a nanobot's signal radius are said to be __in range__ of that nanobot.

Before you start the teleportation process, you should determine which nanobot is the __strongest__ (that is, which has the largest signal radius) and then, for that nanobot, the __total number of nanobots that are in range__ of it, __including itself__.

For example, given the following nanobots:

```txt
pos=<0,0,0>, r=4
pos=<1,0,0>, r=1
pos=<4,0,0>, r=3
pos=<0,2,0>, r=1
pos=<0,5,0>, r=3
pos=<0,0,3>, r=1
pos=<1,1,1>, r=1
pos=<1,1,2>, r=1
pos=<1,3,1>, r=1
```

The strongest nanobot is the first one (position `0,0,0`) because its signal radius, `4` is the largest. Using that nanobot's location and signal radius, the following nanobots are in or out of range:

- The nanobot at `0,0,0` is distance `0` away, and so it is __in range__.
- The nanobot at `1,0,0` is distance `1` away, and so it is __in range__.
- The nanobot at `4,0,0` is distance `4` away, and so it is __in range__.
- The nanobot at `0,2,0` is distance `2` away, and so it is __in range__.
- The nanobot at `0,5,0` is distance `5` away, and so it is __not__ in range.
- The nanobot at `0,0,3` is distance `3` away, and so it is __in range__.
- The nanobot at `1,1,1` is distance `3` away, and so it is __in range__.
- The nanobot at `1,1,2` is distance `4` away, and so it is __in range__.
- The nanobot at `1,3,1` is distance `5` away, and so it is __not__ in range.

In this example, in total, __`7`__ nanobots are in range of the nanobot with the largest signal radius.

Find the nanobot with the largest signal radius. __How many nanobots are in range__ of its signals?