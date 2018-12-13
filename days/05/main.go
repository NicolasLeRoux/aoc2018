package main

import (
    "fmt"
    "io/ioutil"
    "aoc2018/days/05/solver"
    "time"
)

func main() {
    dat, err := ioutil.ReadFile("./input.txt")

    if err != nil {
        panic(err)
    }

    polymers := string(dat)

    start := time.Now()
    resultPartOne := solver.SolvePartOne(polymers)
    t := time.Now()
    elapsed := t.Sub(start)
    fmt.Printf("Answer part one: %d (%+v)\n", len(resultPartOne), elapsed)

    start = time.Now()
    resultPartTwo := solver.SolvePartTwo(polymers)
    t = time.Now()
    elapsed = t.Sub(start)
    fmt.Printf("Answer part two: %d (%+v)\n", resultPartTwo, elapsed)
}