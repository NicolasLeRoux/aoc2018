package solver

import (
    "fmt"
    "testing"
    "reflect"
    "time"
)

func TestSolvePartOne(t *testing.T) {
    suites := []struct {
        input []string; expected int
    }{
        {[]string{
            "[1518-11-01 00:00] Guard #10 begins shift",
            "[1518-11-01 00:05] falls asleep",
            "[1518-11-01 00:25] wakes up",
            "[1518-11-01 00:30] falls asleep",
            "[1518-11-01 00:55] wakes up",
            "[1518-11-01 23:58] Guard #99 begins shift",
            "[1518-11-02 00:40] falls asleep",
            "[1518-11-02 00:50] wakes up",
            "[1518-11-03 00:05] Guard #10 begins shift",
            "[1518-11-03 00:24] falls asleep",
            "[1518-11-03 00:29] wakes up",
            "[1518-11-04 00:02] Guard #99 begins shift",
            "[1518-11-04 00:36] falls asleep",
            "[1518-11-04 00:46] wakes up",
            "[1518-11-05 00:03] Guard #99 begins shift",
            "[1518-11-05 00:45] falls asleep",
            "[1518-11-05 00:55] wakes up",
            }, 24},
    }

    for _, suite := range suites {
        output := SolvePartOne(suite.input[:])

        if output != suite.expected {
            t.Errorf("TestSolvePartOne([...]) got %d but expect %d", output, suite.expected)
        } else {
            fmt.Printf("✓ TestSolvePartOne(%s)\n", suite.input)
        }
    }
}

func TestParseRecord(t *testing.T) {
    timestamp, _ := time.Parse(time.RFC3339, "1518-11-01T00:00:00Z")

    suites := []struct {
        input string; expected Record
    }{
        {"[1518-11-01 00:00] Guard #10 begins shift", Record{timestamp, "Guard #10 begins shift"}},
    }

    for _, suite := range suites {
        output := parseRecord(suite.input)

        if !reflect.DeepEqual(output, suite.expected) {
            t.Errorf("parseRecord(%s) got %+v but expect %+v", suite.input, output, suite.expected)
        } else {
            fmt.Printf("✓ parseRecord(%s)\n", suite.input)
        }
    }
}