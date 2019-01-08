package main

import (
    "fmt"
    "io/ioutil"
    "./solver"
    "time"
)

func main() {
    dat, err := ioutil.ReadFile("./input.txt")

    if err != nil {
        panic(err)
    }

    polymers := string(dat)

    start := time.Now()
    resultPartOne := solver.SolvePartOne(polymers) // 11264 (5.098028ms)
    t := time.Now()
    elapsed := t.Sub(start)
    fmt.Printf("Answer part one: %d (%+v)\n", len(resultPartOne), elapsed)

    start = time.Now()
    resultPartTwo := solver.SolvePartTwo(polymers) // 4552 (164.257525ms)
    t = time.Now()
    elapsed = t.Sub(start)
    fmt.Printf("Answer part two: %d (%+v)\n", resultPartTwo, elapsed)
}
