package solver

import (
    "fmt"
    "testing"
)

func TestSolvePartOne(t *testing.T) {
    suites := []struct {
        input []string; expected int
    }{
        {[]string{
            "1, 1",
            "1, 6",
            "8, 3",
            "3, 4",
            "5, 5",
            "8, 9"}, 17},
    }

    for _, suite := range suites {
        output := SolvePartOne(suite.input[:])

        if output != suite.expected {
            t.Errorf("TestSolvePartOne('%s') got '%d' but expect '%d'", suite.input, output, suite.expected)
        } else {
            fmt.Printf("✓ TestSolvePartOne('%s')\n", suite.input)
        }
    }
}