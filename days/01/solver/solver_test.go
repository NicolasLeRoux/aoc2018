package solver

import (
    "fmt"
    "testing"
)

func TestSolve(t *testing.T) {
    suites := []struct {
        input []string; expected int
    }{
        {[]string{"+1", "-2", "+3", "+1"}, 3},
        {[]string{"+1", "+1", "+1"}, 3},
        {[]string{"+1", "+1", "-2"}, 0},
        {[]string{"-1", "-2", "-3"}, -6},
    }

    for _, suite := range suites {
        output := Solve(suite.input[:])

        if output != suite.expected {
            t.Errorf("Solve(%s) got %d but expect %d", suite.input, output, suite.expected)
        } else {
            fmt.Printf("✓ Solve(%s)\n", suite.input)
        }
    }
}
