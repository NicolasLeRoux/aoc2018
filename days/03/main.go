package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "aoc2018/days/03/solver"
)

func main() {
    dat, err := ioutil.ReadFile("./input.txt")

    if err != nil {
        panic(err)
    }

    claims := strings.Split(string(dat), "\n")

    resultPartOne := solver.SolvePartOne(claims)
    fmt.Printf("Answer part one: %d\n", resultPartOne)
}
