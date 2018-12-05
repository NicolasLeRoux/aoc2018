package solver

import (
    "time"
)

type Record struct {
    timestamp time.Time
    action string
}

func SolvePartOne(words []string) int {
    return 0
}

func parseRecord(record string) Record {
    return Record{time.Now(), ""}
}