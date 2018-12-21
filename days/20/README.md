# Day 20: A Regular Map


While you were learning about instruction pointers, the Elves made considerable progress. When you look up, you discover that the North Pole base construction project has completely surrounded you.

The area you are in is made up entirely of rooms and doors. The rooms are arranged in a grid, and rooms only connect to adjacent rooms when a door is present between them.

For example, drawing rooms as `.`, walls as `#`, doors as `|` or `-`, your current position as `X`, and where north is up, the area you're in might look like this:

```txt
#####
#.|.#
#-###
#.|X#
#####
```

You get the attention of a passing construction Elf and ask for a map. "I don't have time to draw out a map of this place - it's __huge__. Instead, I can give you directions to every room in the facility!" He writes down some directions on a piece of parchment and runs off. In the example above, the instructions might have been `^WNE$`, a regular expression or __"regex"__ (your puzzle input).

The regex matches routes (like `WNE` for "west, north, east") that will take you from your current room through various doors in the facility. In aggregate, the routes will take you through __every door in the facility at least once__; mapping out all of these routes will let you build a proper map and find your way around.

`^` and `$` are at the beginning and end of your regex; these just mean that the regex doesn't match anything outside the routes it describes. (Specifically, `^` matches the start of the route, and `$` matches the end of it.) These characters will not appear elsewhere in the regex.

The rest of the regex matches various sequences of the characters `N` (north), `S` (south), `E` (east), and `W` (west). In the example above, `^WNE$` matches only one route, `WNE`, which means you can move __west, then north, then east__ from your current position. Sequences of letters like this always match that exact route in the same order.

Sometimes, the route can __branch__. A branch is given by a __list of options__ separated by pipes (`|`) and wrapped in parentheses. So, `^N(E|W)N$` contains a branch: after going north, you must choose to go __either east or west__ before finishing your route by going north again. By tracing out the possible routes after branching, you can determine where the doors are and, therefore, where the rooms are in the facility.

For example, consider this regex: `^ENWWW(NEEE|SSE(EE|N))$`

This regex begins with `ENWWW`, which means that from your current position, all routes must begin by moving east, north, and then west three times, in that order. After this, there is a branch. Before you consider the branch, this is what you know about the map so far, with doors you aren't sure about marked with a `?`:

```txt
#?#?#?#?#
?.|.|.|.?
#?#?#?#-#
    ?X|.?
    #?#?#
```

After this point, there is `(NEEE|SSE(EE|N))`. This gives you exactly two options: `NEEE` and `SSE(EE|N)`. By following `NEEE`, the map now looks like this:

```txt
#?#?#?#?#
?.|.|.|.?
#-#?#?#?#
?.|.|.|.?
#?#?#?#-#
    ?X|.?
    #?#?#
```

Now, only `SSE(EE|N)` remains. Because it is in the same parenthesized group as `NEEE`, it starts from the same room `NEEE` started in. It states that starting from that point, there exist doors which will allow you to move south twice, then east; this ends up at another branch. After that, you can either move east twice or north once. This information fills in the rest of the doors:

```txt
#?#?#?#?#
?.|.|.|.?
#-#?#?#?#
?.|.|.|.?
#-#?#?#-#
?.?.?X|.?
#-#-#?#?#
?.|.|.|.?
#?#?#?#?#
```

Once you've followed all possible routes, you know the remaining unknown parts are all walls, producing a finished map of the facility:

```txt
#########
#.|.|.|.#
#-#######
#.|.|.|.#
#-#####-#
#.#.#X|.#
#-#-#####
#.|.|.|.#
#########
```

Sometimes, a list of options can have an __empty option__, like `(NEWS|WNSE|)`. This means that routes at this point could effectively skip the options in parentheses and move on immediately. For example, consider this regex and the corresponding map:

```txt
^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$

###########
#.|.#.|.#.#
#-###-#-#-#
#.|.|.#.#.#
#-#####-#-#
#.#.#X|.#.#
#-#-#####-#
#.#.|.|.|.#
#-###-###-#
#.|.|.#.|.#
###########
```

This regex has one main route which, at three locations, can optionally include additional detours and be valid: `(NEWS|)`, `(WNSE|)`, and `(SWEN|)`. Regardless of which option is taken, the route continues from the position it is left at after taking those steps. So, for example, this regex matches all of the following routes (and more that aren't listed here):

- `ENNWSWWSSSEENEENNN`
- `ENNWSWW`__`NEWS`__`SSSEENEENNN`
- `ENNWSWW`__`NEWS`__`SSSEENEE`__`SWEN`__`NNN`
- `ENNWSWWSSSEEN`__`WNSE`__`EENNN`

By following the various routes the regex matches, a full map of all of the doors and rooms in the facility can be assembled.

To get a sense for the size of this facility, you'd like to determine which room is __furthest__ from you: specifically, you would like to find the room for which the __shortest path to that room would require passing through the most doors__.

- In the first example (`^WNE$`), this would be the north-east corner __`3`__ doors away.
- In the second example (`^ENWWW(NEEE|SSE(EE|N))$`), this would be the south-east corner __`10`__ doors away.
- In the third example (`^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$`), this would be the north-east corner __`18`__ doors away.

Here are a few more examples:

```txt
Regex: ^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$
Furthest room requires passing 23 doors

#############
#.|.|.|.|.|.#
#-#####-###-#
#.#.|.#.#.#.#
#-#-###-#-#-#
#.#.#.|.#.|.#
#-#-#-#####-#
#.#.#.#X|.#.#
#-#-#-###-#-#
#.|.#.|.#.#.#
###-#-###-#-#
#.|.#.|.|.#.#
#############
```

```txt
Regex: ^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$
Furthest room requires passing 31 doors

###############
#.|.|.|.#.|.|.#
#-###-###-#-#-#
#.|.#.|.|.#.#.#
#-#########-#-#
#.#.|.|.|.|.#.#
#-#-#########-#
#.#.#.|X#.|.#.#
###-#-###-#-#-#
#.|.#.#.|.#.|.#
#-###-#####-###
#.|.#.|.|.#.#.#
#-#-#####-#-#-#
#.#.|.|.|.#.|.#
###############
```

__What is the largest number of doors you would be required to pass through to reach a room?__ That is, find the room for which the shortest path from your starting location to that room would require passing through the most doors; what is the fewest doors you can pass through to reach it?