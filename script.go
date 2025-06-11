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
    links := getWikipediaLinks("https://fr.wikipedia.org/wiki/Red_Rising")
    fmt.Println(links)
}

// func wikipediaBFS(startLink string, endLink string) (string){

// }

// urls exemple: https://fr.wikipedia.org/wiki/Red_Rising
func getWikipediaLinks(url string) []string{
    var pageContent string = getPageMainContent(url)
    re := regexp.MustCompile(`/wiki/[^"#:]*`)
    var allLinksBytes [][]byte = re.FindAll([]byte(pageContent), -1)
    var allLinks []string
    for i := 0; i < len(allLinksBytes); i++ {
        link := string(allLinksBytes[i])
        if isValidWikipediaLink(link) {
            allLinks = append(allLinks, link)
        }
    }
    return allLinks
}

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
