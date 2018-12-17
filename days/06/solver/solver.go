package solver

import (
    "strconv"
    "strings"
)

func SolvePartOne(coords []string) int {
    return 0
}

func parseCoord(coord string) (int, int) {
    splitted := strings.Split(string(coord), ", ")
    x, _ := strconv.Atoi(splitted[0])
    y, _ := strconv.Atoi(splitted[1])

    return x, y
}