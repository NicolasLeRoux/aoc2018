package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "aoc2018/days/02/solver"
)

func main() {
    dat, err := ioutil.ReadFile("./input.txt")

    if err != nil {
        panic(err)
    }

    words := strings.Split(string(dat), "\n")

    resultPartOne := solver.SolvePartOne(words)
    fmt.Printf("Answer part one: %d\n", resultPartOne)

    resultPartTwo := solver.SolvePartTwo(words)
    fmt.Printf("Answer part one: %s\n", resultPartTwo)
}
