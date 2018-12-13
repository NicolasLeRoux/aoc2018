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