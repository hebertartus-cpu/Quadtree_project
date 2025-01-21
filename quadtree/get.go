package quadtree

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	// TODO

	for y := 0; y < len(contentHolder); y++ {

		for x := 0; x < len(contentHolder[y]); x++ {

			if y+topLeftY >= q.MinY && x+topLeftX >= q.MinX && (y+topLeftY) < q.root.height && (x+topLeftX) < q.root.width {

				contentHolder[y][x] = q.root.trouve_valeur(x+topLeftX, y+topLeftY)

			} else {
				contentHolder[y][x] = -1
			}
		}
	}
}

// fonction qui, à partir des coordonnées x, y, trouve la valeur corespondante dans un noeud de quadtree
func (noeud node) trouve_valeur(x, y int) (valeur int) {

	//s'occupe des cas où x ou y ne sont pas dontenue dans le noeud.
	//renvoie -2, valeur pour l'erreur ou l'inexistant
	//si utiliser par GetContent, ne s'active normalement jamais
	//
	//cas spéciale pour l'extention Monde infini, cette condition ne soit pas s'activer
	if x < 0 || y < 0 || x >= noeud.width || y >= noeud.height {

		//si l'extention Monde infini n'est pas active
		if len(configuration.Global.InfinitRandom) !=2 {
			
			return -2
		}
	}

	//cas d'arret
	if noeud.isLeaf {

		return noeud.content
	}

	//permet de gérer les cas où les coordonnées sont négatives.
	//Dans le cadre de l'extention Monde infini
	negatifX := 0
	negatifY := 0
	//annule le décalage des coordonnées si nécessaire
	compenseX := 0
	compenseY := 0

	if noeud.topLeftX < 0 {

		negatifX = noeud.topLeftX
		compenseX = noeud.width/2

	}
	if noeud.topLeftY < 0 {

		negatifY = noeud.topLeftY
		compenseY = noeud.height/2
	}

	//gauche
	if x < noeud.width/2 + negatifX {

		//haut
		if y < noeud.height/2 + negatifY {

			return noeud.topLeftNode.trouve_valeur(x, y)
		//bas
		} else {

			return noeud.bottomLeftNode.trouve_valeur(x, y-noeud.height/2 + compenseY)
		}
	//droite
	} else {

		//haut
		if y < noeud.height/2 + negatifY {

			return noeud.topRightNode.trouve_valeur(x-noeud.width/2 + compenseX, y)
		//bas
		} else {
			
			return noeud.bottomRightNode.trouve_valeur(x-noeud.width/2 + compenseX, y-noeud.height/2 + compenseY)
		}
	}
}
