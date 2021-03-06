package solver

import (
    "testing"
    "reflect"
)

func TestSolvePartOne(t *testing.T) {
    suites := []struct {
        description string; inputs []string; expectedA string; expectedB int
    }{
        {
            "Should display 'HI' after 3 seconds with the given points.",
            []string{
                "position=< 9,  1> velocity=< 0,  2>",
                "position=< 7,  0> velocity=<-1,  0>",
                "position=< 3, -2> velocity=<-1,  1>",
                "position=< 6, 10> velocity=<-2, -1>",
                "position=< 2, -4> velocity=< 2,  2>",
                "position=<-6, 10> velocity=< 2, -2>",
                "position=< 1,  8> velocity=< 1, -1>",
                "position=< 1,  7> velocity=< 1,  0>",
                "position=<-3, 11> velocity=< 1, -2>",
                "position=< 7,  6> velocity=<-1, -1>",
                "position=<-2,  3> velocity=< 1,  0>",
                "position=<-4,  3> velocity=< 2,  0>",
                "position=<10, -3> velocity=<-1,  1>",
                "position=< 5, 11> velocity=< 1, -2>",
                "position=< 4,  7> velocity=< 0, -1>",
                "position=< 8, -2> velocity=< 0,  1>",
                "position=<15,  0> velocity=<-2,  0>",
                "position=< 1,  6> velocity=< 1,  0>",
                "position=< 8,  9> velocity=< 0, -1>",
                "position=< 3,  3> velocity=<-1,  1>",
                "position=< 0,  5> velocity=< 0, -1>",
                "position=<-2,  2> velocity=< 2,  0>",
                "position=< 5, -2> velocity=< 1,  2>",
                "position=< 1,  4> velocity=< 2,  1>",
                "position=<-2,  7> velocity=< 2, -2>",
                "position=< 3,  6> velocity=<-1, -1>",
                "position=< 5,  0> velocity=< 1,  0>",
                "position=<-6,  0> velocity=< 2,  0>",
                "position=< 5,  9> velocity=< 1, -2>",
                "position=<14,  7> velocity=<-2,  0>",
                "position=<-3,  6> velocity=< 2, -1>",
            },
            "#...#..###\n#...#...#.\n#...#...#.\n#####...#.\n#...#...#.\n#...#...#.\n#...#...#.\n#...#..###\n",
            3,
        },
    }

    for _, suite := range suites {
        t.Run(suite.description, func(t *testing.T) {
            outputA, outputB := SolvePartOne(suite.inputs)

            if outputA != suite.expectedA || outputB != suite.expectedB {
                t.Errorf("TestSolvePartOne('[...]') got %s and %d but expect %s and %d",
                         outputA,
                         outputB,
                         suite.expectedA,
                         suite.expectedB)
            }
        })
    }
}

func TestParse(t *testing.T) {
    suites := []struct {
        description string; input string; expected Point
    }{
        {
            "Should return (-3, 6) as position and (2, -1) as velocity.",
            "position=<-3,  6> velocity=< 2, -1>",
            Point{-3, 6, 2, -1},
        },
        {
            "Should return (10, -3) as position and (-1, 1) as velocity.",
            "position=<10, -3> velocity=<-1,  1>",
            Point{10, -3, -1, 1},
        },
        {
            "Should return (20597, -20297) as position and (-2, 2) as velocity.",
            "position=< 20597, -20297> velocity=<-2,  2>",
            Point{20597, -20297, -2, 2},
        },
    }

    for _, suite := range suites {
        t.Run(suite.description, func(t *testing.T) {
            output := parse(suite.input)

            if !reflect.DeepEqual(output, suite.expected) {
                t.Errorf("TestSolvePartOne('%s') got %+v but expect %+v",
                         suite.input,
                         output,
                         suite.expected)
            }
        })
    }
}
