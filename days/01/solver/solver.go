package solver

import (
    "strconv"
)

// SolvePartOne is the chronal calibration solver of the part one of the day 01
// puzzle.
func SolvePartOne(nbrs []string) int {
    var result int = 0

    for i := 0; i < len(nbrs); i++ {
        nb, err := strconv.Atoi(nbrs[i])

        if err == nil {
            result += nb
        } else {
            panic(err)
        }
    }

    return result
}

// SolvePartTwo is the chronal calibration solver of the part two of the day 01
// puzzle.
func SolvePartTwo(nbrs []string) int {
    var m = make(map[int]bool)
    var i, sum int = 0, 0
    m[0] = true

    for {
        nb, err := strconv.Atoi(nbrs[i])

        if err == nil {
            sum += nb
        } else {
            panic(err)
        }

        // Make it circular
        if i < len(nbrs) - 1 {
            i++
        } else {
            i = 0
        }

        // Exit at the first duplicated sum
        if m[sum] {
            break
        }

        m[sum] = true
    }

    return sum
}
