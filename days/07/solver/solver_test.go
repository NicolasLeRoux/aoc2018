package solver

import (
    "fmt"
    "testing"
)

func TestSolvePartOne(t *testing.T) {
    suites := []struct {
        input []string; expected string
    }{
        {[]string{
            "Step C must be finished before step A can begin.",
            "Step C must be finished before step F can begin.",
            "Step A must be finished before step B can begin.",
            "Step A must be finished before step D can begin.",
            "Step B must be finished before step E can begin.",
            "Step D must be finished before step E can begin.",
            "Step F must be finished before step E can begin."}, "CABDFE"},
        {[]string{
            "Step A must be finished before step D can begin.",
            "Step B must be finished before step D can begin.",
            "Step C must be finished before step D can begin."}, "ABCD"},
    }

    for _, suite := range suites {
        output := SolvePartOne(suite.input[:])

        if output != suite.expected {
            t.Errorf("TestSolvePartOne('[...]') got '%s' but expect '%s'", output, suite.expected)
        } else {
            fmt.Printf("✓ TestSolvePartOne('[...]')\n")
        }
    }
}

func TestSolvePartTwo(t *testing.T) {
    suites := []struct {
        input []string; nbWorkers, stepDuration, expected int
    }{
        {[]string{
            "Step C must be finished before step A can begin.",
            "Step C must be finished before step F can begin.",
            "Step A must be finished before step B can begin.",
            "Step A must be finished before step D can begin.",
            "Step B must be finished before step E can begin.",
            "Step D must be finished before step E can begin.",
            "Step F must be finished before step E can begin."}, 2, 0, 15},
    }

    for _, suite := range suites {
        output := SolvePartTwo(suite.input[:], suite.nbWorkers, suite.stepDuration)

        if output != suite.expected {
            t.Errorf("TestSolvePartTwo('[...]') got '%d' but expect '%d'", output, suite.expected)
        } else {
            fmt.Printf("✓ TestSolvePartTwo('[...]')\n")
        }
    }
}