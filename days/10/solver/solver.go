package solver

import (
    "strings"
    "strconv"
)

type Point struct {
    px int
    py int
    vx int
    vy int
}

func (p Point) next() Point {
    return Point{
        px: p.px + p.vx,
        py: p.py + p.vy,
        vx: p.vx,
        vy: p.vy,
    }
}

func SolvePartOne(inputs []string) string {
    return ""
}

func parse(lign string) Point {
    slice := strings.Split(lign, "<")
    pos := strings.Split(strings.Split(slice[1], ">")[0], ",")
    vel := strings.Split(strings.Split(slice[2], ">")[0], ",")

    px, _ := strconv.Atoi(strings.Trim(pos[0], " "))
    py, _ := strconv.Atoi(strings.Trim(pos[1], " "))
    vx, _ := strconv.Atoi(strings.Trim(vel[0], " "))
    vy, _ := strconv.Atoi(strings.Trim(vel[1], " "))

    return Point{px, py, vx, vy}
}
