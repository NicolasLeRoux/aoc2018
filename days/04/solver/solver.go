package solver

import (
    "time"
    "strconv"
    "sort"
    "strings"
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
    records := make([]Record, len(words))
    for i := 0; i < len(words); i++ {
        records[i] = parseRecord(words[i])
    }

    schedules := buildSchedule(records)
    id := getBiggestSleeper(schedules)

    return id * getMostAsleepMinute(schedules, id)
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

type Schedule struct {
    Id int
    Sleeping [60]bool
}

func buildSchedule(records []Record) []Schedule {
    sort.Sort(ByTimestamp(records))
    var resp []Schedule
    var id int
    var start time.Time
    var sleeping [60]bool

    for i := 0; i < len(records); i++ {
        record := records[i]

        if strings.Contains(record.Action, "begins shift") {
            if i != 0 {
                resp = append(resp, Schedule{id, sleeping})
                copy(sleeping[:], make([]bool, 60))
            }

            spitted := strings.Split(record.Action, " ")
            id, _ = strconv.Atoi(strings.Trim(spitted[1], "#"))
        } else if record.Action == "falls asleep" {
            start = record.Timestamp
        } else if record.Action == "wakes up" {
            var init int
            if start.Hour() != 0 {
                init = 0
            } else {
                init = start.Minute()
            }
            end := record.Timestamp.Minute()

            for j := init; j < end; j++ {
                sleeping[j] = true
            }
        }
    }

    resp = append(resp, Schedule{id, sleeping})

    return resp
}

func getBiggestSleeper(schedules []Schedule) int {
    guards := make(map[int]int)
    for i := 0; i < len(schedules); i++ {
        id := schedules[i].Id
        val, ok := guards[id]

        var sum int
        for j := 0; j < 60; j++ {
            if schedules[i].Sleeping[j] {
                sum++
            }
        }

        if ok {
            guards[id] = val + sum
        } else {
            guards[id] = sum
        }
    }

    var max, id int
    for key, val := range guards {
        if max < val {
            max = val
            id = key
        }
    }

    return id
}

func getMostAsleepMinute(schedules []Schedule, id int) int {
    minutes := make(map[int]int)
    for i := 0; i < len(schedules); i++ {
        if schedules[i].Id == id {
            for j := 0; j < 60; j++ {
                if schedules[i].Sleeping[j] {
                    val, ok := minutes[j]

                    if ok {
                        minutes[j] = val + 1
                    } else {
                        minutes[j] = 1
                    }
                }
            }
        }
    }

    var max, minute int
    for key, val := range minutes {
        if max < val {
            max = val
            minute = key
        }
    }

    return minute
}

func findMostAsleepMinute(schedules []Schedule) (int, int) {
    minutes := make(map[string]int)
    for i := 0; i < len(schedules); i++ {
        for j := 0; j < 60; j++ {
            if schedules[i].Sleeping[j] {
                key := strings.Join([]string{
                    strconv.Itoa(schedules[i].Id),
                    strconv.Itoa(j),
                }, ",")
                val, ok := minutes[key]

                if ok {
                    minutes[key] = val + 1
                } else {
                    minutes[key] = 1
                }
            }
        }
    }

    var max int
    var token string
    for key, val := range minutes {
        if max < val {
            max = val
            token = key
        }
    }
    splitted := strings.Split(string(token), ",")
    id, _ := strconv.Atoi(splitted[0])
    minute, _ := strconv.Atoi(splitted[1])

    return id, minute
}

func SolvePartTwo(words []string) int {
    records := make([]Record, len(words))
    for i := 0; i < len(words); i++ {
        records[i] = parseRecord(words[i])
    }

    schedules := buildSchedule(records)
    id, minute := findMostAsleepMinute(schedules)

    return id * minute
}