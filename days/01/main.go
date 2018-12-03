package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "aoc2018/days/01/solver"
)

func main() {
    dat, err := ioutil.ReadFile("./input.txt")

    if err != nil {
        panic(err)
    }

    operations := strings.Split(string(dat), "\n")

    result := solver.SolvePartOne(operations[:len(operations) - 1])
    fmt.Printf("Answer part one: %d\n", result)
}
