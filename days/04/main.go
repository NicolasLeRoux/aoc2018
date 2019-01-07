package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "./solver"
)

func main() {
    dat, err := ioutil.ReadFile("./input.txt")

    if err != nil {
        panic(err)
    }

    records := strings.Split(string(dat), "\n")

    resultPartOne := solver.SolvePartOne(records)
    fmt.Printf("Answer part one: %d\n", resultPartOne)

    resultPartTwo := solver.SolvePartTwo(records)
    fmt.Printf("Answer part two: %d\n", resultPartTwo)
}
