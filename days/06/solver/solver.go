package solver

import (
    "strconv"
    "strings"
    "fmt"
)

type Coord struct {
    X int
    Y int
}

func SolvePartOne(lines []string) int {
    grid := make(map[string]int)
    coords := make(map[int]Coord)
    var xMin, xMax, yMin, yMax int = 0, 0, 0, 0

    for i := 0; i < len(lines); i++ {
        grid[lines[i]] = i

        x, y := parseCoord(lines[i])
        coords[i] = Coord{x, y}
        if x > xMax {
            xMax = x
        }
        if y > yMax {
            yMax = y
        }
    }

    for x := xMin; x < xMax; x++ {
        for y := yMin; y < yMax; y++ {
            key := strings.Join([]string{strconv.Itoa(x), strconv.Itoa(y)}, ", ")
            _, ok := grid[key]

            if !ok {
                grid[key] = -1
            }

            // Closest coord to set the value in the grid...
            // I also need a map name to coord to loop over...
        }
    }

    fmt.Printf("%d - %d; %d - %d\n", xMin, xMax, yMin, yMax)
    fmt.Printf("%+v\n", grid)

    return 0
}

func parseCoord(coord string) (int, int) {
    splitted := strings.Split(string(coord), ", ")
    x, _ := strconv.Atoi(splitted[0])
    y, _ := strconv.Atoi(splitted[1])

    return x, y
}

func closestCoord(current Coord, coords map[int]Coord) (int, bool) {
    var minDist, idx int = 1000, 0
    var single bool = true

    for key, val := range coords {
        dist := (val.X - current.X) + (val.Y - current.Y)

        if dist < minDist {
            minDist = dist
            idx = key
            single = true
        } else if dist == minDist {
            single = false
        }
    }

    return idx, single
}