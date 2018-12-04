package solver

import (
    "strings"
)

func SolvePartOne(words []string) int {
    var sumDouble, sumTiple int = 0, 0

    for i := 0; i < len(words); i++ {
        isDouble, isTriple := hasDoubleOrTriple(words[i])

        if isDouble {
            sumDouble++
        }

        if isTriple {
            sumTiple++
        }
    }

    return sumDouble * sumTiple
}

func hasDoubleOrTriple(word string) (bool, bool) {
    var sumDouble, sumTiple int = 0, 0
    letters := strings.Split(word, "")
    m := make(map[string]int)

    for j := 0; j < len(letters); j++ {
        letter := letters[j]
        val, ok := m[letter]

        if ok {
            m[letter] = val + 1
        } else {
            m[letter] = 1
        }

        val = m[letter]
        if (val == 2 ) {
            sumDouble++;
        } else if (val == 3) {
            sumDouble--;
            sumTiple++;
        }
    }

    return sumDouble > 0, sumTiple > 0
}

func SolvePartTwo(words []string) string {
    length := len(words[0])
    var match string

    Loop:
        for i := 0; i < len(words) - 1; i++ {
            for j := i + 1; j < len(words); j++ {
                match = matchingString(words[i], words[j])

                if len(match) + 1 == length {
                    break Loop
                }
            }
        }

    return match
}

func matchingString(wordA, wordB string) string {
    lettersA := strings.Split(wordA, "")
    lettersB := strings.Split(wordB, "")
    var matchings []string

    for i := 0; i < len(lettersA); i++ {
        if lettersA[i] == lettersB[i] {
            matchings = append(matchings, lettersA[i])
        }
    }

    return strings.Join(matchings[:], "")
}
