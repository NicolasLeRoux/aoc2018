package solver

import (
    "strings"
)

func SolvePartOne(polymer string) string {
    splitted := strings.Split(polymer, "")
    var stack []string
    stack = append(stack, splitted[0])

    for i := 1; i < len(splitted); i++ {
        letter := splitted[i]
        previous := stack[len(stack) - 1]

        if strings.ToUpper(letter) == strings.ToUpper(previous) &&
            (isUpperCase(letter) && isLowerCase(previous) ||
            isUpperCase(previous) && isLowerCase(letter)) {
            stack = stack[:len(stack) - 1]
        } else {
            stack = append(stack, letter)
        }
    }

    return strings.Join(stack, "")
}

func isUpperCase(word string) bool {
    return strings.ToUpper(word) == word
}

func isLowerCase(word string) bool {
    return strings.ToLower(word) == word
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
    splitted := strings.Split(polymer, "")
    aphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
        "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y",
        "Z"}

    for i := 0; i < len(aphabet); i++ {
        var stack []string
        stack = append(stack, splitted[0])

        for j := 1; j < len(splitted); j++ {
            letter := splitted[j]
            previous := stack[len(stack) - 1]

            if aphabet[i] != letter && strings.ToLower(aphabet[i]) != letter {
                if strings.ToUpper(letter) == strings.ToUpper(previous) &&
                    (isUpperCase(letter) && isLowerCase(previous) ||
                    isUpperCase(previous) && isLowerCase(letter)) {
                    stack = stack[:len(stack) - 1]
                } else {
                    stack = append(stack, letter)
                }
            }
        }

        mapResult[aphabet[i]] = len(strings.Join(stack, ""))
    }

    var smallest int = len(polymer)
    for _, val := range mapResult {
        if smallest > val {
            smallest = val
        }
    }

    return smallest
}