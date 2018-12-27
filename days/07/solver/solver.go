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