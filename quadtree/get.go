package quadtree

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	// TODO

	for y := 0; y < len(contentHolder); y++ {

		for x := 0; x < len(contentHolder[y]); x++ {

			if y+topLeftY >= 0 && x+topLeftX >= 0 && (y+topLeftY) < q.root.height && (x+topLeftX) < q.root.width {

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
	if x < 0 || y < 0 || x >= noeud.width || y >= noeud.height {

		return -2
	}

	//cas d'arret
	if noeud.isLeaf {

		return noeud.content
	}

	if x < noeud.width/2 {

		if y < noeud.height/2 {

			return noeud.topLeftNode.trouve_valeur(x, y)
		} else {

			return noeud.bottomLeftNode.trouve_valeur(x, y-noeud.height/2)
		}
	} else {

		if y < noeud.height/2 {

			return noeud.topRightNode.trouve_valeur(x-noeud.width/2, y)
		} else {

			return noeud.bottomRightNode.trouve_valeur(x-noeud.width/2, y-noeud.height/2)
		}
	}
}
