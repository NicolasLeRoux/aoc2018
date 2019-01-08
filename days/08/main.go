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
    resultPartOne := solver.SolvePartOne(numbers) // 43996 (1.112693ms)
    t := time.Now()
    elapsed := t.Sub(start)
    fmt.Printf("Answer part one: %d (%+v)\n", resultPartOne, elapsed)

    start = time.Now()
    resultPartTwo := solver.SolvePartTwo(numbers) // 35189 (1.029503ms)
    t = time.Now()
    elapsed = t.Sub(start)
    fmt.Printf("Answer part two: %d (%+v)\n", resultPartTwo, elapsed)
}
