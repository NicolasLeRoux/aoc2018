package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "aoc2018/days/07/solver"
    "time"
)

func main() {
    dat, err := ioutil.ReadFile("./input.txt")

    if err != nil {
        panic(err)
    }

    instructions := strings.Split(string(dat), "\n")

    start := time.Now()
    resultPartOne := solver.SolvePartOne(instructions) // ABDCJLFMNVQWHIRKTEUXOZSYPG (140.637µs)
    t := time.Now()
    elapsed := t.Sub(start)
    fmt.Printf("Answer part one: %s (%+v)\n", resultPartOne, elapsed)

    start = time.Now()
    resultPartTwo := solver.SolvePartTwo(instructions, 5, 60) // 896 (1.561888ms)
    t = time.Now()
    elapsed = t.Sub(start)
    fmt.Printf("Answer part two: %d (%+v)\n", resultPartTwo, elapsed)
}
