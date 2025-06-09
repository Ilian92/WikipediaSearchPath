package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
    graphe := [][]int{
        {1, 2}, 
        {0, 2}, 
        {0, 1, 3, 4}, 
        {2}, 
        {2}}
    
    // chemin := bfs(graphe, 0, 4)
    // if chemin != nil {
    //     fmt.Printf("Chemin le plus court de 0 à 4: %v\n", chemin)
    //     fmt.Printf("Distance: %d\n", len(chemin)-1)
    // } else {
    //     fmt.Println("Aucun chemin trouvé")
    // }

    test, err := http.Get("https://fr.wikipedia.org/wiki/Red_Rising")
    
    fmt.Println("CA MARCHE",test)

    fmt.Println("pas march :,(", err)

}

func bfs(graphe [][]int, start, end int) []int {
    if start == end {
        return []int{start}
    }

    queue := []struct {
        node int
        path []int
    }{{start, []int{start}}}

    visited := make([]bool, len(graphe))
    visited[start] = true

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        for _, neighbor := range graphe[current.node] {
            if neighbor == end {
                return append(current.path, neighbor)
            }

            if !visited[neighbor] {
                visited[neighbor] = true
                newPath := make([]int, len(current.path))
                copy(newPath, current.path)
                newPath = append(newPath, neighbor)
                queue = append(queue, struct{node int; path []int}{neighbor, newPath})
            }
        }
    }

    return nil
}

// urls exemple: https://fr.wikipedia.org/wiki/Red_Rising
// func getWikipediaLinks(url string) []string {

// }

func isValidWikipediaLink(link string) bool {
    excludePatterns := []string{
        "Category:", "Template:", "Help:", "File:", "Special:",
        "Talk:", "User:", "Wikipedia:", "Portal:",
    }

    for _, pattern := range excludePatterns {
        if strings.Contains(link, pattern) {
            return false
        }
    }

    return true
}