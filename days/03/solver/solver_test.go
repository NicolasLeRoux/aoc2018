package solver

import (
    "fmt"
    "testing"
    "reflect"
)

func TestSolvePartOne(t *testing.T) {
    suites := []struct {
        input []string; expected int
    }{
        {[]string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}, 4},
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

func TestParseClaim(t *testing.T) {
    suites := []struct {
        input string; expected Claim
    }{
        {"#1 @ 1,3: 4x4", Claim{1, 1, 3, 4, 4}},
        {"#78 @ 242,958: 22x16", Claim{78, 242, 958, 22, 16}},
        {"#765 @ 559,438: 29x23", Claim{765, 559, 438, 29, 23}},
        {"#1311 @ 407,579: 27x23", Claim{1311, 407, 579, 27, 23}},
    }

    for _, suite := range suites {
        output := parseClaim(suite.input)

        if !reflect.DeepEqual(output, suite.expected) {
            t.Errorf("parseClaim(%s) got %+v but expect %+v", suite.input, output, suite.expected)
        } else {
            fmt.Printf("✓ parseClaim(%s)\n", suite.input)
        }
    }
}
