package solver

import (
    "testing"
)

func TestSolvePartOne(t *testing.T) {
    suites := []struct {
        description, input string; expected int
    }{
        {
            "Should return 8317 given 10 players and a last marble at 1618 points.",
            "10 players; last marble is worth 1618 points",
            8317,
        },
        {
            "Should return 146373 given 13 players and a last marble at 7999 points.",
            "13 players; last marble is worth 7999 points",
            146373,
        },
        {
            "Should return 2764 given 17 players and a last marble at 1104 points.",
            "17 players; last marble is worth 1104 points",
            2764,
        },
        {
            "Should return 54718 given 21 players and a last marble at 6111 points.",
            "21 players; last marble is worth 6111 points",
            54718,
        },
        {
            "Should return 37305 given 30 players and a last marble at 5807 points.",
            "30 players; last marble is worth 5807 points",
            37305,
        },
    }

    for _, suite := range suites {
        t.Run(suite.description, func(t *testing.T) {
            output := SolvePartOne(suite.input)

            if output != suite.expected {
                t.Errorf("TestSolvePartOne('%s') got %d but expect %d",
                         suite.input,
                         output,
                         suite.expected)
            }
        })
    }
}

func TestParse(t *testing.T) {
    suites := []struct {
        description, input string; expectedA, expectedB int
    }{
        {
            "Should return 10 and 1618 for the given sentence.",
            "10 players; last marble is worth 1618 points",
            10,
            1618,
        },
    }

    for _, suite := range suites {
        t.Run(suite.description, func(t *testing.T) {
            outputA, outputB := parse(suite.input)

            if outputA != suite.expectedA || outputB != suite.expectedB {
                t.Errorf("TestParse('%s') got %d and %d but expect %d and %d",
                         suite.input,
                         outputA,
                         outputB,
                         suite.expectedA,
                         suite.expectedB)
            }
        })
    }
}
