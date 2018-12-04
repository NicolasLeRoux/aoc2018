package solver

import (
    "strings"
    "strconv"
)

type Claim struct {
    Id string
    OffsetX int
    OffsetY int
    Width int
    Height int
}

func SolvePartOne(words []string) int {
    fabric := make(map[string]string)

    for _, word := range words {
        claim := parseClaim(word)

        for i := claim.OffsetX; i < claim.OffsetX + claim.Width; i++ {
            for j := claim.OffsetY; j < claim.OffsetY + claim.Height; j++ {
                id := strings.Join([]string{strconv.Itoa(i), strconv.Itoa(j)}, ",")
                _, ok := fabric[id]

                if ok {
                    fabric[id] = "X"
                } else {
                    fabric[id] = claim.Id
                }
            }
        }
    }

    sum := 0
    for _, val := range fabric {
        if val == "X" {
            sum++
        }
    }

    return sum
}

func parseClaim(claim string) Claim {
    splitted := strings.Fields(claim)
    // Extract ID
    id := splitted[0]
    // Extract offsets
    offsets := strings.Split(splitted[2], ",")
    offsetX, _ := strconv.Atoi(offsets[0])
    offsetY, _ := strconv.Atoi(strings.Trim(offsets[1], ":"))
    // Extract boundary
    boundary := strings.Split(splitted[3], "x")
    width, _ := strconv.Atoi(boundary[0])
    height, _ := strconv.Atoi(boundary[1])

    return Claim{id, offsetX, offsetY, width, height}
}

func SolvePartTwo(words []string) int {
    fabric := make(map[string]string)
    mapOfClaimId := make(map[string]bool)

    for _, word := range words {
        claim := parseClaim(word)
        mapOfClaimId[claim.Id] = true

        for i := claim.OffsetX; i < claim.OffsetX + claim.Width; i++ {
            for j := claim.OffsetY; j < claim.OffsetY + claim.Height; j++ {
                id := strings.Join([]string{strconv.Itoa(i), strconv.Itoa(j)}, ",")
                val, ok := fabric[id]

                if ok {
                    fabric[id] = "X"
                    delete(mapOfClaimId, claim.Id)
                    delete(mapOfClaimId, val)
                } else {
                    fabric[id] = claim.Id
                }
            }
        }
    }

    var idWthOverlap string
    for key, _ := range mapOfClaimId {
        idWthOverlap = key
    }
    result, _ := strconv.Atoi(strings.Trim(idWthOverlap, "#"))

    return result
}