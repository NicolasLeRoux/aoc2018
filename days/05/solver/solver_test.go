package solver

import (
    "fmt"
    "testing"
)

func TestSolvePartOne(t *testing.T) {
    suites := []struct {
        input, expected string
    }{
        {"aA", ""},
        {"abBA", ""},
        {"abAB", "abAB"},
        {"aabAAB", "aabAAB"},
        {"dabAcCaCBAcCcaDA", "dabCBAcaDA"},
    }

    for _, suite := range suites {
        output := SolvePartOne(suite.input[:])

        if output != suite.expected {
            t.Errorf("TestSolvePartOne('%s') got '%s' but expect '%s'", suite.input, output, suite.expected)
        } else {
            fmt.Printf("✓ TestSolvePartOne('%s')\n", suite.input)
        }
    }
}

func TestSolvePartTwo(t *testing.T) {
    suites := []struct {
        input string; expected int
    }{
        {"dabAcCaCBAcCcaDA", 4},
    }

    for _, suite := range suites {
        output := SolvePartTwo(suite.input[:])

        if output != suite.expected {
            t.Errorf("TestSolvePartTwo('%s') got '%d' but expect '%d'", suite.input, output, suite.expected)
        } else {
            fmt.Printf("✓ TestSolvePartTwo('%s')\n", suite.input)
        }
    }
}