package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "./solver"
    "time"
)

func main() {
    dat, err := ioutil.ReadFile("./input.txt")

    if err != nil {
        panic(err)
    }

    coords := strings.Split(string(dat), "\n")

    start := time.Now()
    resultPartOne := solver.SolvePartOne(coords) // 4754 (285.910531ms)
    t := time.Now()
    elapsed := t.Sub(start)
    fmt.Printf("Answer part one: %d (%+v)\n", resultPartOne, elapsed)

    start = time.Now()
    resultPartTwo := solver.SolvePartTwo(coords, 10000) // 42344 (190.770034ms)
    t = time.Now()
    elapsed = t.Sub(start)
    fmt.Printf("Answer part two: %d (%+v)\n", resultPartTwo, elapsed)
}
