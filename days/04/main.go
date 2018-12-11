package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "aoc2018/days/04/solver"
)

func main() {
    dat, err := ioutil.ReadFile("./input.txt")

    if err != nil {
        panic(err)
    }

    records := strings.Split(string(dat), "\n")

    resultPartOne := solver.SolvePartOne(records)
    fmt.Printf("Answer part one: %d\n", resultPartOne)
}
