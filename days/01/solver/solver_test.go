package solver

import (
    "testing"
)

func TestSolvePartOne(t *testing.T) {
    suites := []struct {
        description string; input []string; expected int
    }{
        {"Run 1 - 2 + 3 + 1", []string{"+1", "-2", "+3", "+1"}, 3},
        {"Run 1 + 1 + 1", []string{"+1", "+1", "+1"}, 3},
        {"Run 1 + 1 - 2", []string{"+1", "+1", "-2"}, 0},
        {"Run -1 - 2 - 3", []string{"-1", "-2", "-3"}, -6},
    }

    for _, suite := range suites {
        t.Run(suite.description, func(t *testing.T) {
            output := SolvePartOne(suite.input[:])

            if output != suite.expected {
                t.Errorf("TestSolvePartOne(%s) got %d but expect %d",
                         suite.input,
                         output,
                         suite.expected)
            }
        })
    }
}

func TestSolvePartTwo(t *testing.T) {
    suites := []struct {
        description string; input []string; expected int
    }{
        {"Run 1 - 2 + 3 + 1", []string{"+1", "-2", "+3", "+1"}, 2},
        {"Run 1 - 1", []string{"+1", "-1"}, 0},
        {"Run 3 + 3 + 4 - 2- 4", []string{"+3", "+3", "+4", "-2", "-4"}, 10},
        {"Run -6 + 3 + 8 + 5 - 6", []string{"-6", "+3", "+8", "+5", "-6"}, 5},
        {"Run 7 + 7 - 2 - 7 - 4", []string{"+7", "+7", "-2", "-7", "-4"}, 14},
    }

    for _, suite := range suites {
        t.Run(suite.description, func(t *testing.T) {
            output := SolvePartTwo(suite.input[:])

            if output != suite.expected {
                t.Errorf("TestSolvePartTwo(%s) got %d but expect %d",
                         suite.input,
                         output,
                         suite.expected)
            }
        })
    }
}

func BenchmarkSolvePartOne(t *testing.B) {
    for i := 0; i < t.N; i++ {
        SolvePartOne([]string{"+1", "-2", "+3", "+1"})
    }
}

func BenchmarkSolvePartTwo(t *testing.B) {
    for i := 0; i < t.N; i++ {
        SolvePartTwo([]string{"+1", "-2", "+3", "+1"})
    }
}
