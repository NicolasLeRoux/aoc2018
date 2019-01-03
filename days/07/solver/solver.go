package solver

import (
    "strings"
    "sort"
)

type ByAlphabetical []string

func (a ByAlphabetical) Len() int {
    return len(a)
}
func (a ByAlphabetical) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}
func (a ByAlphabetical) Less(i, j int) bool {
    return a[i] < a[j]
}

func SolvePartOne(instructions []string) string {
    deps := make(map[string][]string)
    parents := make(map[string][]string)

    for i := 0; i < len(instructions); i++ {
        item, next := parseInstruction(instructions[i])

        valA, okA := deps[next]
        if okA {
            deps[next] = append(valA, item)
        } else {
            deps[next] = []string{item}
        }

        valB, okB := parents[item]
        if okB {
            parents[item] = append(valB, next)
        } else {
            parents[item] = []string{next}
        }
    }

    for key, _ := range parents {
        _, ok := deps[key]

        if !ok {
            deps[key] = []string{}
        }
    }

    var result []string
    for len(deps) != 0 {
        var ready []string

        // Extract item wth deps
        for item, childs := range deps {
            if len(childs) == 0 {
                ready = append(ready, item)
            }
        }

        sort.Sort(ByAlphabetical(ready))
        first := ready[0]

        result = append(result, first)

        delete(deps, first)

        for item, childs := range deps {
            for i := 0; i < len(childs); i++ {
                if childs[i] == first {
                    deps[item] = append(childs[:i], childs[i+1:]...)
                }
            }
        }
    }

    return strings.Join(result, "")
}

func parseInstruction(instruction string) (string, string) {
    splitted := strings.Split(string(instruction), " ")

    return splitted[1], splitted[7]
}

type Load struct {
    Duration int
    Name string
}

func SolvePartTwo(instructions []string, nbWorkers, stepDuration int) int {
    deps := make(map[string][]string)
    parents := make(map[string][]string)
    letters := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7,
        "H": 8, "I": 9, "J": 10, "K": 11, "L": 12, "M": 13, "N": 14, "O": 15, "P": 16,
        "Q": 17, "R": 18, "S": 19, "T": 20, "U": 21, "V": 22, "W": 23, "X": 24, "Y": 25,
        "Z": 26}

    for i := 0; i < len(instructions); i++ {
        item, next := parseInstruction(instructions[i])

        valA, okA := deps[next]
        if okA {
            deps[next] = append(valA, item)
        } else {
            deps[next] = []string{item}
        }

        valB, okB := parents[item]
        if okB {
            parents[item] = append(valB, next)
        } else {
            parents[item] = []string{next}
        }
    }

    for key, _ := range parents {
        _, ok := deps[key]

        if !ok {
            deps[key] = []string{}
        }
    }

    // Init workers
    workers := make(map[int]Load)
    for i := 0; i < nbWorkers; i++ {
        workers[i] = Load{0, "."}
    }

    var duration int = 0
    for len(deps) != 0 {
        var ready []string

        // Extract item wth deps
        for item, childs := range deps {
            if len(childs) == 0 {
                ready = append(ready, item)
            }
        }

        sort.Sort(ByAlphabetical(ready))

        for id, load := range workers {
            if load.Name == "." && len(ready) > 0 {
                first := ready[0]
                ready = ready[1:]
                delete(deps, first)

                workers[id] = Load{stepDuration + letters[first] - 1, first}
            } else if load.Name != "." {
                if load.Duration != 0 {
                    workers[id] = Load{load.Duration - 1, load.Name}
                } else {
                    workers[id] = Load{0, "."}

                    for item, childs := range deps {
                        for i := 0; i < len(childs); i++ {
                            if childs[i] == load.Name {
                                deps[item] = append(childs[:i], childs[i+1:]...)
                            }
                        }
                    }
                }
            }
        }

        duration++
    }

    return duration
}
