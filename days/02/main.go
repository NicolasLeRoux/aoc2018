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

    array := strings.Split(string(dat), "\n")
    words := array[:len(array) - 1]

    resultPartOne := solver.SolvePartOne(words)
    fmt.Printf("Answer part one: %d\n", resultPartOne)

    resultPartTwo := solver.SolvePartTwo(words)
    fmt.Printf("Answer part one: %s\n", resultPartTwo)
}
