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

    claims := strings.Split(string(dat), "\n")

    resultPartOne := solver.SolvePartOne(claims)
    fmt.Printf("Answer part one: %d\n", resultPartOne)

    resultPartTwo := solver.SolvePartTwo(claims)
    fmt.Printf("Answer part two: %d\n", resultPartTwo)
}
