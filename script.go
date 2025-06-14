package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func main() {
    // var startUrl string = "https://fr.wikipedia.org/wiki/France"
    // var endUrl string = "https://fr.wikipedia.org/wiki/Rose_Bertin"

    var startUrl string
    var endUrl string

    fmt.Println("Vous allez devoir rentrer 2 url correspondant à la page de départ et la page d'arrivé (https://fr.wikipedia.org/wiki/[La page])")
    fmt.Print("Veuillez Choisir la page de début:")
    fmt.Scan(&startUrl)

    fmt.Print("Veuillez Choisir la page de d'arrivé:")
    fmt.Scan(&endUrl)

    fmt.Println("Recherche du chemin...")
    chemin := wikipediaBFS(startUrl, endUrl, 4)
    
    if chemin != nil {
        fmt.Printf("Chemin trouvé (%d étapes):\n", len(chemin))
        for i, page := range chemin {
            fmt.Printf("%d. %s\n", i+1, page)
        }
    } else {
        fmt.Println("Aucun chemin trouvé")
    }
    
    // Résultat retourné:

    // Recherche du chemin...
    // Exploration: https://fr.wikipedia.org/wiki/France (profondeur 0)
    // Trouvé x liens valides...
    // Chemin trouvé (3 étapes):
    // 1. https://fr.wikipedia.org/wiki/France
    // 2. https://fr.wikipedia.org/wiki/Paris
    // 3. https://fr.wikipedia.org/wiki/Rose_Bertin
}

func wikipediaBFS(startLink string, endLink string, maxDepth int) ([]string){
    if startLink == endLink {
        return []string{startLink}
    }
    type Node struct {
        url   string
        depth int
    }
    
    var queue []Node
	visited := make(map[string]bool)
	parent := make(map[string]string)

	queue = append(queue, Node{startLink, 0})
	visited[startLink] = true

    parent[startLink] = ""

    for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

        if current.depth >= maxDepth {
            continue
        }
        
        fmt.Printf("Exploration: %s (profondeur %d)\n", current.url, current.depth)
        
		if current.url == endLink {
			return reconstructPathLink(parent, endLink)
		}

        links := getWikipediaLinks(current.url)
        fmt.Printf("Trouvé %d liens valides\n", len(links))

		for _, link := range links {
            var fullUrl string = "https://fr.wikipedia.org" + link

            // Ajout d'une vérification du la présence de l'url dans la liste pour optimiser le code
            if fullUrl == endLink {
                parent[fullUrl] = current.url
                return reconstructPathLink(parent, endLink)
            }

			if !visited[fullUrl] {
				visited[fullUrl] = true
				parent[fullUrl] = current.url
				queue = append(queue, Node{fullUrl, current.depth + 1})
			}
		}
        // On peu rajouter un délai ici si on ne veut pas harceler les serveurs de Wikipedia
	}
	return nil
}

func reconstructPathLink(parent map[string]string, endLink string) []string {
    chemin := []string{}
    current := endLink
    
    // Remonter depuis la fin jusqu'au début
    for current != "" {
		// Ajouter au début
        chemin = append([]string{current}, chemin...)
        current = parent[current]
    }
    
    return chemin
}

// urls exemple: https://fr.wikipedia.org/wiki/Red_Rising
func getWikipediaLinks(url string) []string{
    var pageContent string = getPageMainContent(url)
    re := regexp.MustCompile(`/wiki/[^"#\s]*`)
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
    // Décoder l'URL pour gérer les caractères encodés
    decodedLink, err := url.QueryUnescape(link)
    if err != nil {
        decodedLink = link
    }
    
    excludePatterns := []string{
        "/wiki/Category", "/wiki/Template", "/wiki/Help", 
        "/wiki/File", "/wiki/Special", "/wiki/Talk", 
        "/wiki/User", "/wiki/Wikipedia", "/wiki/Portal", 
        "/wiki/Aide", "/wiki/Spécial", "/wiki/Catégorie",
        "/wiki/Modèle", "/wiki/Fichier", "/wiki/Discussion",
        "/wiki/Utilisateur", "/wiki/Wikipédia", "/wiki/Portail",
    }

    for _, pattern := range excludePatterns {
        if strings.HasPrefix(decodedLink, pattern) {
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
