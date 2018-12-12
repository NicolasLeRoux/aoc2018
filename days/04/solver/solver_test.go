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
            }, 240},
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

func TestBuildSchedule(t *testing.T) {
    timestamp01, _ := time.Parse(time.RFC3339, "1518-11-01T00:00:00Z")
    timestamp02, _ := time.Parse(time.RFC3339, "1518-11-01T00:05:00Z")
    timestamp03, _ := time.Parse(time.RFC3339, "1518-11-01T00:25:00Z")
    timestamp04, _ := time.Parse(time.RFC3339, "1518-11-01T00:30:00Z")
    timestamp05, _ := time.Parse(time.RFC3339, "1518-11-01T00:55:00Z")
    timestamp06, _ := time.Parse(time.RFC3339, "1518-11-02T00:03:00Z")
    timestamp07, _ := time.Parse(time.RFC3339, "1518-11-02T00:05:00Z")
    timestamp08, _ := time.Parse(time.RFC3339, "1518-11-02T00:30:00Z")

    suites := []struct {
        input []Record; expected []Schedule
    }{
        {[]Record {
                {timestamp01, "Guard #10 begins shift"},
                {timestamp02, "falls asleep"},
                {timestamp03, "wakes up"},
            }, []Schedule{
                {10, [60]bool {false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
            },
        },
        {[]Record {
                {timestamp01, "Guard #10 begins shift"},
                {timestamp02, "falls asleep"},
                {timestamp03, "wakes up"},
                {timestamp04, "falls asleep"},
                {timestamp05, "wakes up"},
            }, []Schedule{
                {10, [60]bool {false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false}},
            },
        },
        {[]Record {
                {timestamp01, "Guard #10 begins shift"},
                {timestamp02, "falls asleep"},
                {timestamp03, "wakes up"},
                {timestamp06, "Guard #11 begins shift"},
                {timestamp07, "falls asleep"},
                {timestamp08, "wakes up"},
            }, []Schedule{
                {10, [60]bool {false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
                {11, [60]bool {false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
            },
        },
    }

    for _, suite := range suites {
        output := buildSchedule(suite.input)

        if !reflect.DeepEqual(output, suite.expected) {
            t.Errorf("buildSchedule(%s) got %+v but expect %+v", suite.input, output, suite.expected)
        } else {
            fmt.Printf("✓ buildSchedule(%s)\n", suite.input)
        }
    }
}

func TestGetBiggestSleeper(t *testing.T) {
    suites := []struct {
        input []Schedule; expected int
    }{
        {
            []Schedule{
                {10, [60]bool {false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
                {21, [60]bool {false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
                {42, [60]bool {false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
            },
            21,
        },
        {
            []Schedule{
                {10, [60]bool {false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
                {21, [60]bool {false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
                {10, [60]bool {false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
            },
            10,
        },
    }

    for _, suite := range suites {
        output := getBiggestSleeper(suite.input)

        if !reflect.DeepEqual(output, suite.expected) {
            t.Errorf("getBiggestSleeper([...]) got %d but expect %d", output, suite.expected)
        } else {
            fmt.Printf("✓ getBiggestSleeper([...])\n")
        }
    }
}

func TestGetMostAsleepMinute(t *testing.T) {
    suites := []struct {
        schedule []Schedule; id, expected int
    }{
        {
            []Schedule{
                {10, [60]bool {true, true, true, true, true, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
                {10, [60]bool {false, false, false, false, false, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
            },
            10,
            10,
        },
    }

    for _, suite := range suites {
        output := getMostAsleepMinute(suite.schedule, suite.id)

        if !reflect.DeepEqual(output, suite.expected) {
            t.Errorf("getBiggestSleeper([...]) got %d but expect %d", output, suite.expected)
        } else {
            fmt.Printf("✓ getBiggestSleeper([...])\n")
        }
    }
}

func TestSolvePartTwo(t *testing.T) {
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
            }, 4455},
    }

    for _, suite := range suites {
        output := SolvePartTwo(suite.input[:])

        if output != suite.expected {
            t.Errorf("SolvePartTwo([...]) got %d but expect %d", output, suite.expected)
        } else {
            fmt.Printf("✓ SolvePartTwo(%s)\n", suite.input)
        }
    }
}

