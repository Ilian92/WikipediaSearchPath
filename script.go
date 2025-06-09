package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
    // graphe := [][]int{
    //     {1, 2}, 
    //     {0, 2}, 
    //     {0, 1, 3, 4}, 
    //     {2}, 
    //     {2}}
    
    // chemin := bfs(graphe, 0, 4)
    // if chemin != nil {
    //     fmt.Printf("Chemin le plus court de 0 à 4: %v\n", chemin)
    //     fmt.Printf("Distance: %d\n", len(chemin)-1)
    // } else {
    //     fmt.Println("Aucun chemin trouvé")
    // }

    test := getPageMainContent("https://fr.wikipedia.org/wiki/Red_Rising")
    
    fmt.Println(test, "salut mon gaté")

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

func getPageMainContent(link string) string {
    res, err := http.Get(link)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    content, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    re := regexp.MustCompile(`(?s)<main id="content" class="mw-body">(.*?)</main>`)
    match := re.FindStringSubmatch(string(content))
    if len(match) > 1 {
        return match[1]
    }

    log.Println("Pas de balise <main> trouvée.")
    return ""
}

// func wikipediaBFS(startLink string, endLink string) (string){

// }