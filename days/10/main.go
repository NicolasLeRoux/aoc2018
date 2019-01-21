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

    points := strings.Split(string(dat), "\n")

    start := time.Now()
    resultPartOne, resultPartTwo := solver.SolvePartOne(points)
    t := time.Now()
    elapsed := t.Sub(start)
    fmt.Printf("Answer part one: %s (%+v)\n", resultPartOne, elapsed)
    fmt.Printf("Answer part one: %d (Idem)\n", resultPartTwo)
}