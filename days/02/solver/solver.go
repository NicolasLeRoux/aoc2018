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
