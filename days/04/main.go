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

    records := strings.Split(string(dat), "\n")

    start := time.Now()
    resultPartOne := solver.SolvePartOne(records) // 67558 (598.217µs)
    t := time.Now()
    elapsed := t.Sub(start)
    fmt.Printf("Answer part one: %d (%+v)\n", resultPartOne, elapsed)

    start = time.Now()
    resultPartTwo := solver.SolvePartTwo(records) // 78990 (1.864759ms)
    t = time.Now()
    elapsed = t.Sub(start)
    fmt.Printf("Answer part two: %d (%+v)\n", resultPartTwo, elapsed)
}
