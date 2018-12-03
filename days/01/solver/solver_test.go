package solver

import (
    "fmt"
    "testing"
)

func TestSolvePartOne(t *testing.T) {
    suites := []struct {
        input []string; expected int
    }{
        {[]string{"+1", "-2", "+3", "+1"}, 3},
        {[]string{"+1", "+1", "+1"}, 3},
        {[]string{"+1", "+1", "-2"}, 0},
        {[]string{"-1", "-2", "-3"}, -6},
    }

    for _, suite := range suites {
        output := SolvePartOne(suite.input[:])

        if output != suite.expected {
            t.Errorf("TestSolvePartOne(%s) got %d but expect %d", suite.input, output, suite.expected)
        } else {
            fmt.Printf("✓ TestSolvePartOne(%s)\n", suite.input)
        }
    }
}

func TestSolvePartTwo(t *testing.T) {
    suites := []struct {
        input []string; expected int
    }{
        {[]string{"+1", "-2", "+3", "+1"}, 2},
        {[]string{"+1", "-1"}, 0},
        {[]string{"+3", "+3", "+4", "-2", "-4"}, 10},
        {[]string{"-6", "+3", "+8", "+5", "-6"}, 5},
        {[]string{"+7", "+7", "-2", "-7", "-4"}, 14},
    }

    for _, suite := range suites {
        output := SolvePartTwo(suite.input[:])

        if output != suite.expected {
            t.Errorf("TestSolvePartTwo(%s) got %d but expect %d", suite.input, output, suite.expected)
        } else {
            fmt.Printf("✓ TestSolvePartTwo(%s)\n", suite.input)
        }
    }
}
