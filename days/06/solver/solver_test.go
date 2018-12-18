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

func TestParseCoord(t *testing.T) {
    suites := []struct {
        input string; expectedA, expectedB int
    }{
        {"1, 1", 1, 1},
        {"259, 316", 259, 316},
        {"213, 120", 213, 120},
    }

    for _, suite := range suites {
        outputA, outputB := parseCoord(suite.input)

        if outputA != suite.expectedA || outputB != suite.expectedB {
            t.Errorf("TestParseCoord('%s') got '%d' and '%d' but expect '%d' and '%d'",
                suite.input,
                outputA,
                outputB,
                suite.expectedA,
                suite.expectedB)
        } else {
            fmt.Printf("✓ TestParseCoord('%s')\n", suite.input)
        }
    }
}

func TestClosestCoord(t *testing.T) {
    suites := []struct {
        current Coord; coords map[int]Coord; expected int
    }{
        {Coord{1, 1}, map[int]Coord{2: Coord{10, 10}}, 2},
        {Coord{1, 1}, map[int]Coord{1: Coord{2, 2}, 2: Coord{3, 3}}, 1},
    }

    for _, suite := range suites {
        output := closestCoord(suite.current, suite.coords)

        if output != suite.expected {
            t.Errorf("TestclosestCoord('%+v', '%+v') got '%d' but expect '%d'",
                suite.current,
                suite.coords,
                output,
                suite.expected)
        } else {
            fmt.Printf("✓ TestclosestCoord('%+v', '%+v')\n", suite.current, suite.coords)
        }
    }
}