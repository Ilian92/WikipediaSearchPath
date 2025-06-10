package main

// import "fmt"

func bfs(graphe [][]int, start int, end int) []int {
	// var returnMessage string
	if start == end {
		return nil
	}
	var queue []int
	visited := make(map[int]bool)
	parent := make(map[int]int)

	//Noeud de départ
	queue = append(queue, start)
	visited[start] = true
	// Le nœud de départ n'a pas de parent
	parent[start] = -1
	//Tant qu'il y a un élément dans la file d'attente on continue
	for len(queue) > 0 {
		//On prend le premier élément de la file d'attente
		current := queue[0]
		//On le supprime de la file d'attente
		queue = queue[1:]

		
		if current == end {
			chemin := reconstructPath(parent, end)
			return chemin
		}
		//On traverse les éléments de l'élément courant du graphe
		for _, currentBranches := range graphe[current] {
			//Si visited ne contient pas cet éléments on le rajoute à la file d'attente
			if !visited[currentBranches] {
				visited[currentBranches] = true
				parent[currentBranches] = current
				queue = append(queue, currentBranches)
			}
		}
	}
	return nil
}

// Fonction pour reconstruire le chemin depuis la map des parents
func reconstructPath(parent map[int]int, end int) []int {
    chemin := []int{}
    current := end
    
    // Remonter depuis la fin jusqu'au début
    for current != -1 {
		// Ajouter au début
        chemin = append([]int{current}, chemin...)
        current = parent[current]
    }
    
    return chemin
}