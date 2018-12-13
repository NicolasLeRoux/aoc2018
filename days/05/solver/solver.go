package solver

import (
    "strings"
)

func SolvePartOne(polymer string) string {
    splitted := strings.Split(polymer, "")

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

func SolvePartTwo(polymer string) int {
    mapResult := make(map[string]int)
    aphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
        "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y",
        "Z"}

    for i := 0; i < len(aphabet); i++ {
        shortenedPolymer := removeLetter(polymer, aphabet[i])
        resultingPolymer := SolvePartOne(shortenedPolymer)

        mapResult[aphabet[i]] = len(resultingPolymer)
    }

    var smallest int = len(polymer)
    for _, val := range mapResult {
        if smallest > val {
            smallest = val
        }
    }

    return smallest
}

func removeLetter(polymers string, letter string) string {
    splitted := strings.Split(polymers, "")

    for i := 0; i < len(splitted); i++ {
        if letter == splitted[i] || strings.ToLower(letter) == splitted[i] {
            splitted = append(splitted[:i], splitted[i + 1:]...)
            i = max(0, i - 1)
        }
    }

    return strings.Join(splitted, "")
}