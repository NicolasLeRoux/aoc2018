package solver

import (
    "strconv"
)

type Node struct {
    Parent *Node
    NbChild int
    Childs []*Node
    NbMetadata int
    Metadata []int
}

func SolvePartOne(items []string) int {
    nbChild, _ := strconv.Atoi(items[0])
    nbMetadata, _ := strconv.Atoi(items[1])
    rootNode := Node{
        NbChild: nbChild,
        NbMetadata: nbMetadata,
    }
    flatNode := []*Node{&rootNode}
    currentNode := &rootNode

    for i := 2; i < len(items); i++ {
        number, _ := strconv.Atoi(items[i])

        if len(currentNode.Childs) == currentNode.NbChild && currentNode == &rootNode {
            rootNode.Metadata = append(rootNode.Metadata, number)
        } else if len(currentNode.Childs) < currentNode.NbChild {
            next, _ := strconv.Atoi(items[i + 1])
            node := Node{
                Parent: currentNode,
                NbChild: number,
                NbMetadata: next,
            }
            i++

            flatNode = append(flatNode, &node)
            currentNode.Childs = append(currentNode.Childs, &node)
            currentNode = &node
        } else if len(currentNode.Metadata) < currentNode.NbMetadata {
            currentNode.Metadata = append(currentNode.Metadata, number)
        } else {
            currentNode = currentNode.Parent
            i--
        }
    }

    sum := 0
    for i := 0; i < len(flatNode); i++ {
        for j := 0; j < len(flatNode[i].Metadata); j++ {
            sum += flatNode[i].Metadata[j]
        }
    }

    return sum
}
