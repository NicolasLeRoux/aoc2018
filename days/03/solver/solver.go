package solver

import (
    "strings"
    "strconv"
)

type Claim struct {
    Id int
    OffsetX int
    OffsetY int
    Width int
    Height int
}

func SolvePartOne(claims []string) int {
    return 0
}

func parseClaim(claim string) Claim {
    splitted := strings.Fields(claim)
    // Extract ID
    id, _ := strconv.Atoi(strings.Trim(splitted[0], "#"))
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