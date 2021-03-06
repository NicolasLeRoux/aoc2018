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

    numbers := strings.Split(string(dat), " ")

    start := time.Now()
    resultPartOne := solver.SolvePartOne(numbers) // 43996 (735.881µs)
    t := time.Now()
    elapsed := t.Sub(start)
    fmt.Printf("Answer part one: %d (%+v)\n", resultPartOne, elapsed)

    start = time.Now()
    resultPartTwo := solver.SolvePartTwo(numbers) // 35189 (679.103µs)
    t = time.Now()
    elapsed = t.Sub(start)
    fmt.Printf("Answer part two: %d (%+v)\n", resultPartTwo, elapsed)
}
