package solver

import (
    "testing"
)

func TestSolvePartOne(t *testing.T) {
    suites := []struct {
        description string; input []string; expected int
    }{
        {
            "Four nodes A[B[10,11,12],C[D[99],2],1,2,2], so 1+1+2+10+11+12+2+99=138",
            []string{"2", "3", "0", "3", "10", "11", "12", "1", "1", "0", "1", "99", "2", "1", "1", "2"},
            138,
        },
    }

    for _, suite := range suites {
        t.Run(suite.description, func(t *testing.T) {
            output := SolvePartOne(suite.input[:])

            if output != suite.expected {
                t.Errorf("TestSolvePartOne(%d) got %d but expect %d",
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
        {
            "Four nodes A[B[10,11,12],C[D[99],2],1,2,2], so 33+33+0=66",
            []string{"2", "3", "0", "3", "10", "11", "12", "1", "1", "0", "1", "99", "2", "1", "1", "2"},
            66,
        },
    }

    for _, suite := range suites {
        t.Run(suite.description, func(t *testing.T) {
            output := SolvePartTwo(suite.input[:])

            if output != suite.expected {
                t.Errorf("TestSolvePartOne(%d) got %d but expect %d",
                         suite.input,
                         output,
                         suite.expected)
            }
        })
    }
}

func BenchmarkSolvePartOne(t *testing.B) {
    for i := 0; i < t.N; i++ {
        SolvePartOne([]string{"2", "3", "0", "3", "10", "11", "12", "1", "1", "0", "1", "99", "2", "1", "1", "2"})
    }
}

func BenchmarkSolvePartTwo(t *testing.B) {
    for i := 0; i < t.N; i++ {
        SolvePartTwo([]string{"2", "3", "0", "3", "10", "11", "12", "1", "1", "0", "1", "99", "2", "1", "1", "2"})
    }
}
