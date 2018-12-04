package solver

import (
    "fmt"
    "testing"
)

func TestSolvePartOne(t *testing.T) {
    suites := []struct {
        input []string; expected int
    }{
        {[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, 12},
        {[]string{"aaa", "bbb", "ccc", "dd", "ee"}, 6},
        {[]string{"aaa", "bbb", "cc", "dd", "ee"}, 6},
        {[]string{"abc", "def", "ghi"}, 0},
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

func TestHasDoubleOrTriple(t *testing.T) {
    suites := []struct {
        input string; expectedDouble, expectedTriple bool
    }{
        {"abcdef", false, false},
        {"aa", true, false},
        {"aabb", true, false},
        {"aaabb", true, true},
        {"aaabbb", false, true},
    }

    for _, suite := range suites {
        outputDouble, outputTriple := hasDoubleOrTriple(suite.input[:])

        if outputDouble != suite.expectedDouble && outputTriple != suite.expectedTriple {
            t.Errorf("hasDoubleOrTriple(%s) got %t and %t but expect %t and %t",
                suite.input,
                outputDouble,
                outputTriple,
                suite.expectedDouble,
                suite.expectedTriple)
        } else {
            fmt.Printf("✓ hasDoubleOrTriple(%s)\n", suite.input)
        }
    }
}
