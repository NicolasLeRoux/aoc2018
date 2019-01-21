package main

import (
    "fmt"
    "io/ioutil"
    "./solver"
    "time"
)

func main() {
    dat, err := ioutil.ReadFile("./input_01.txt")

    if err != nil {
        panic(err)
    }

    sentence := string(dat)

    start := time.Now()
    resultPartOne := solver.SolvePartOne(sentence) // 413188 (10.150475ms)
    t := time.Now()
    elapsed := t.Sub(start)
    fmt.Printf("Answer part one: %d (%+v)\n", resultPartOne, elapsed)

    dat, err = ioutil.ReadFile("./input_02.txt")

    if err != nil {
        panic(err)
    }

    sentence = string(dat)
    start = time.Now()
    resultPartTwo := solver.SolvePartOne(sentence) // 3377272893 (1.507907367s)
    t = time.Now()
    elapsed = t.Sub(start)
    fmt.Printf("Answer part two: %d (%+v)\n", resultPartTwo, elapsed)
}
