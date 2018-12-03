package solver

import (
    "strconv"
)

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

func SolvePartTwo(nbrs []string) int {
    return 0
}