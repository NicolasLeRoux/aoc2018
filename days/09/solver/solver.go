package solver

import (
    "strings"
    "strconv"
    "container/list"
)

func SolvePartOne(input string) int {
    nbPlayers, lastMarblePoints := parse(input)
    l := list.New()
    current := l.PushBack(0)
    scores := make([]int, nbPlayers)

    for i := 1; i <= lastMarblePoints; i++ {
        if i % 23 == 0 {
            item := current.Prev()
            if item == nil {
                item = l.Back()
            }
            for j := 0; j < 6; j++ {
                item = item.Prev()
                if item == nil {
                    item = l.Back()
                }

                if j == 4 {
                    current = item
                }
            }

            score := scores[i % nbPlayers]
            marble, _ := item.Value.(int)
            scores[i % nbPlayers] = score + i + marble

            l.Remove(item)
        } else {
            next := current.Next()
            if next == nil {
                next = l.Front()
            }
            current = l.InsertAfter(i, next)
        }
    }

    max := 0
    for i := 0; i < len(scores); i++ {
        if scores[i] > max {
            max = scores[i]
        }
    }

    return max
}

func parse(input string) (int, int) {
    words := strings.Split(input, " ")
    nbPlayers, _ := strconv.Atoi(words[0])
    lastMarblePoints, _ := strconv.Atoi(words[6])

    return nbPlayers, lastMarblePoints
}
