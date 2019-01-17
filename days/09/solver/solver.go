package solver

import (
    "strings"
    "strconv"
)

func SolvePartOne(input string) int {
    return 0
}

func parse(input string) (int, int) {
    words := strings.Split(input, " ")
    nbPlayers, _ := strconv.Atoi(words[0])
    lastMarblePoints, _ := strconv.Atoi(words[6])

    return nbPlayers, lastMarblePoints
}

func getCircularIndex(length, index int) int {
    if length == index {
        return index
    } else {
        return index % length
    }
}
