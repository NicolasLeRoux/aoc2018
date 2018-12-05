package solver

import (
    "time"
    "strconv"
    "sort"
)

type Record struct {
    Timestamp time.Time
    Action string
}

type ByTimestamp []Record

func (a ByTimestamp) Len() int {
    return len(a)
}
func (a ByTimestamp) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}
func (a ByTimestamp) Less(i, j int) bool {
    return a[i].Timestamp.Before(a[j].Timestamp)
}

func SolvePartOne(words []string) int {
    sort.Sort(ByTimestamp(people))

    return 0
}

func parseRecord(record string) Record {
    year, _ := strconv.Atoi(record[1:5])
    month, _ := strconv.Atoi(record[6:8])
    day, _ := strconv.Atoi(record[9:11])
    hour, _ := strconv.Atoi(record[12:14])
    minute, _ := strconv.Atoi(record[15:17])
    timestamp := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
    action := record[19:len(record)]

    return Record{timestamp, action}
}