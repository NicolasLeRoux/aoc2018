package solver

import (
    "strings"
)

func SolvePartOne(polymers string) string {
    splitted := strings.Split(polymers, "")

    for i := 0; i < len(splitted); i++ {
        letter := splitted[i]

        if isUpperCase(letter) {
            previous, next := getAdjacent(splitted, i)
            lower := strings.ToLower(letter)

            if lower == previous {
                splitted = append(splitted[:i - 1], splitted[i + 1:]...)
                i = max(0, i - 3)
            } else if lower == next {
                splitted = append(splitted[:i], splitted[i + 2:]...)
                i = max(0, i - 2)
            }
        }
    }

    return strings.Join(splitted, "")
}

func isUpperCase(word string) bool {
    return strings.ToUpper(word) == word
}

func getAdjacent(array []string, idx int) (string, string) {
    if idx == 0 {
        return "", array[1]
    } else if idx < len(array) - 1 {
        return array[idx - 1], array[idx + 1]
    } else {
        return array[idx - 1], ""
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}