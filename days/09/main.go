package main

import (
    "fmt"
    "io/ioutil"
    "./solver"
    "time"
)

func main() {
    dat, err := ioutil.ReadFile("./input.txt")

    if err != nil {
        panic(err)
    }

    sentence := string(dat)

    start := time.Now()
    resultPartOne := solver.SolvePartOne(sentence) // 413188 (10.150475ms)
    t := time.Now()
    elapsed := t.Sub(start)
    fmt.Printf("Answer part one: %d (%+v)\n", resultPartOne, elapsed)
}
