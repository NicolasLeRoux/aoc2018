package solver

import (
    "strconv"
    "strings"
    "math"
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

    for x := xMin; x <= xMax; x++ {
        for y := yMin; y <= yMax; y++ {
            key := strings.Join([]string{strconv.Itoa(x), strconv.Itoa(y)}, ", ")
            _, ok := grid[key]

            if !ok {
                closest, single := closestCoord(Coord{x, y}, coords)

                if single {
                    grid[key] = closest
                }
            }
        }
    }
    areaBefore := getArea(grid)

    for x := xMin - 1; x <= xMax + 1; x++ {
        for y := yMin - 1; y <= yMax + 1; y++ {
            key := strings.Join([]string{strconv.Itoa(x), strconv.Itoa(y)}, ", ")
            _, ok := grid[key]

            if !ok {
                closest, single := closestCoord(Coord{x, y}, coords)

                if single {
                    grid[key] = closest
                }
            }
        }
    }
    areaAfter := getArea(grid)
    area := sameArea(areaBefore, areaAfter)
    max, _ := maxArea(area)

    return max
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
        dist := int(math.Abs(float64(val.X - current.X)) +
            math.Abs(float64(val.Y - current.Y)))

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

func getArea(grid map[string]int) map[int]int {
    area := make(map[int]int)

    for _, val := range grid {
        sum, ok := area[val]

        if ok {
            area[val] = sum + 1
        } else {
            area[val] = 1
        }
    }

    return area
}

func sameArea(areaA, areaB map[int]int) map[int]int {
    area := make(map[int]int)

    for key, valA := range areaA {
        valB := areaB[key]

        if valA == valB {
            area[key] = valA
        }
    }

    return area
}

func maxArea(area map[int]int) (int, int) {
    var maxVal, maxKey int = 0, 0

    for key, val := range area {
        if maxVal < val {
            maxVal = val
            maxKey = key
        }
    }

    return maxVal, maxKey
}

func SolvePartTwo(lines []string, limit int) int {
    grid := make(map[string]int)
    coords := make(map[int]Coord)
    var xMin, xMax, yMin, yMax int = 0, 0, 0, 0

    for i := 0; i < len(lines); i++ {
        x, y := parseCoord(lines[i])
        coords[i] = Coord{x, y}
        if x > xMax {
            xMax = x
        }
        if y > yMax {
            yMax = y
        }
    }

    for x := xMin; x <= xMax; x++ {
        for y := yMin; y <= yMax; y++ {
            key := strings.Join([]string{strconv.Itoa(x), strconv.Itoa(y)}, ", ")
            _, ok := grid[key]

            if !ok {
                dist := getAbsoluteValue(Coord{x, y}, coords)

                grid[key] = dist
            }
        }
    }

    var result int = 0
    for _, val := range grid {

        if val < limit {
            result += 1
        }
    }

    return result
}

func getAbsoluteValue(current Coord, coords map[int]Coord) int {
    var absoluteValue int = 0

    for _, val := range coords {
        dist := int(math.Abs(float64(val.X - current.X)) +
            math.Abs(float64(val.Y - current.Y)))

        absoluteValue += dist
    }

    return absoluteValue
}