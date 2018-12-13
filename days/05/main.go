package main

import (
    "fmt"
    "io/ioutil"
    "aoc2018/days/05/solver"
)

func main() {
    dat, err := ioutil.ReadFile("./input.txt")

    if err != nil {
        panic(err)
    }

    polymers := string(dat)

    resultPartOne := solver.SolvePartOne(polymers)
    fmt.Printf("Answer part one: %d\n", len(resultPartOne))
}