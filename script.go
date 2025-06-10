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
    graphe := [][]int{
        {1, 2}, 
        {0, 2}, 
        {0, 1, 3, 4}, 
        {2}, 
        {2}}


    fmt.Println(bfs(graphe, 0, 4))

}

// urls exemple: https://fr.wikipedia.org/wiki/Red_Rising
// func getWikipediaLinks(url string) []string {

// }

func isValidWikipediaLink(link string) bool {
    excludePatterns := []string{
        "Category:", "Template:", "Help:", "File:", "Special:",
        "Talk:", "User:", "Wikipedia:", "Portal:", "Aide:",
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