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
    resultPartOne := solver.SolvePartOne(points)
    t := time.Now()
    elapsed := t.Sub(start)
    fmt.Printf("Answer part one: %d (%+v)\n", resultPartOne, elapsed)
}