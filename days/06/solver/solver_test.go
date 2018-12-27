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
        current Coord; coords map[int]Coord; expected int; multi bool
    }{
        {Coord{1, 1}, map[int]Coord{2: Coord{10, 10}}, 2, true},
        {Coord{1, 1}, map[int]Coord{1: Coord{2, 2}, 2: Coord{3, 3}}, 1, true},
        {Coord{2, 2}, map[int]Coord{1: Coord{1, 1}, 2: Coord{3, 3}}, 1, false},
    }

    for _, suite := range suites {
        outputA, outputB := closestCoord(suite.current, suite.coords)

        if outputA != suite.expected && outputB != suite.multi {
            t.Errorf("TestclosestCoord('%+v', '%+v') got '%d' and '%t' but expect '%d' and '%t'",
                suite.current,
                suite.coords,
                outputA,
                outputB,
                suite.expected,
                suite.multi)
        } else {
            fmt.Printf("✓ TestclosestCoord('%+v', '%+v')\n", suite.current, suite.coords)
        }
    }
}

func TestSolvePartTwo(t *testing.T) {
    suites := []struct {
        input []string; limit, expected int
    }{
        {[]string{
            "1, 1",
            "1, 6",
            "8, 3",
            "3, 4",
            "5, 5",
            "8, 9"}, 32, 16},
    }

    for _, suite := range suites {
        output := SolvePartTwo(suite.input[:], suite.limit)

        if output != suite.expected {
            t.Errorf("TestSolvePartTwo('%s', '%d') got '%d' but expect '%d'",
                suite.input,
                suite.limit,
                output,
                suite.expected)
        } else {
            fmt.Printf("✓ TestSolvePartTwo('%s', '%d')\n", suite.input, suite.limit)
        }
    }
}

func TestGetAbsoluteValue(t *testing.T) {
    suites := []struct {
        current Coord; coords map[int]Coord; expected int
    }{
        {Coord{1, 1}, map[int]Coord{2: Coord{10, 10}}, 18},
        {Coord{1, 1}, map[int]Coord{1: Coord{2, 2}, 2: Coord{3, 3}}, 6},
        {Coord{2, 2}, map[int]Coord{1: Coord{1, 1}, 2: Coord{3, 3}}, 4},
        {Coord{4, 3}, map[int]Coord{1: Coord{1, 1}, 2: Coord{1, 6},
            3: Coord{8, 3}, 4: Coord{3, 4}, 5: Coord{5, 5}, 6: Coord{8, 9}}, 30},
        {Coord{2, 4}, map[int]Coord{1: Coord{1, 1}, 2: Coord{1, 6},
            3: Coord{8, 3}, 4: Coord{3, 4}, 5: Coord{5, 5}, 6: Coord{8, 9}}, 30},
        {Coord{4, 4}, map[int]Coord{1: Coord{1, 1}, 2: Coord{1, 6},
            3: Coord{8, 3}, 4: Coord{3, 4}, 5: Coord{5, 5}, 6: Coord{8, 9}}, 28},
    }

    for _, suite := range suites {
        output := getAbsoluteValue(suite.current, suite.coords)

        if output != suite.expected {
            t.Errorf("TestGetAbsoluteValue('%+v', '%+v') got '%d' but expect '%d'",
                suite.current,
                suite.coords,
                output,
                suite.expected)
        } else {
            fmt.Printf("✓ TestGetAbsoluteValue('%+v', '%+v')\n", suite.current, suite.coords)
        }
    }
}