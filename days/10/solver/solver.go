package solver

import (
    "strings"
    "strconv"
)

type Point struct {
    Px int
    Py int
    Vx int
    Vy int
}

func (p Point) next() Point {
    return Point{
        Px: p.Px + p.Vx,
        Py: p.Py + p.Vy,
        Vx: p.Vx,
        Vy: p.Vy,
    }
}

func SolvePartOne(inputs []string) (string, int) {
    points := make([]Point, 0, len(inputs))

    var pxMinRef, pxMaxRef, pyMinRef, pyMaxRef int = 0, 0, 0, 0
    for i := 0; i < len(inputs); i++ {
        p := parse(inputs[i])
        points = append(points, p)

        if p.Px < pxMinRef {
            pxMinRef = p.Px
        }
        if p.Px > pxMaxRef {
            pxMaxRef = p.Px
        }
        if p.Py < pyMinRef {
            pyMinRef = p.Py
        }
        if p.Py > pyMaxRef {
            pyMaxRef = p.Py
        }
    }


    area := (pxMaxRef - pxMinRef) * (pyMaxRef - pyMinRef)
    isSmaller := true

    newPoints := make([]Point, len(points))
    iter := 0
    for isSmaller {
        var pxMin, pxMax, pyMin, pyMax = 0, 0, 0, 0

        for i := 0; i < len(points); i++ {
            next := points[i].next()
            newPoints[i] = next

            if next.Px < pxMin {
                pxMin = next.Px
            }
            if next.Px > pxMax {
                pxMax = next.Px
            }
            if next.Py < pyMin {
                pyMin = next.Py
            }
            if next.Py > pyMax {
                pyMax = next.Py
            }
        }

        newArea := (pxMax - pxMin) * (pyMax - pyMin)
        if newArea > area {
            isSmaller = false
        } else {
            copy(points, newPoints)

            pxMinRef = pxMin
            pxMaxRef = pxMax
            pyMinRef = pyMin
            pyMaxRef = pyMax

            iter++
        }
        area = newArea
    }

    pointMap := make(map[string]string)
    for i := 0; i < len(points); i++ {
        p := points[i]
        key := strings.Join([]string{
            strconv.Itoa(p.Px),
            strconv.Itoa(p.Py)}, ",")
        pointMap[key] = "#"
    }

    letters := make([]string, 0, area)
    for i := pyMinRef; i <= pyMaxRef; i++ {
        for j := pxMinRef; j <= pxMaxRef; j++ {
            key := strings.Join([]string{
                strconv.Itoa(j),
                strconv.Itoa(i)}, ",")
            val, ok := pointMap[key]

            if ok {
                letters = append(letters, val)
            } else {
                letters = append(letters, ".")
            }
        }
        letters = append(letters, "\n")
    }

    text := strings.Join(letters, "")

    return text, iter
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
