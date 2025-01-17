package quadtree


//ce fichier est dédié à des fonctions servant à la modification 
//et l'utilisation d'un quadtree ou de la racine d'un quadtree


//renvoie la width d'un quadtree
func (q Quadtree) Width() int {

	return q.width
}

//renvoie la height d'un quadtree
func (q Quadtree) Height() int {

	return q.height
}


//permet de trouver la valeur enregistré dans un quadtree 
//aux coordonnées (x, y)
func (q *Quadtree) ValeurQuad(x, y int) (valeur int) {

	valeur = q.root.trouve_valeur(x, y)

	return valeur
}

//permet de modifier la valeur enregistré dans un quadtree 
//aux coordonnées (x, y) en "sol"
func (q Quadtree) Modif_Case(x, y, sol int) {

	q.root.noeud_modif(x, y, sol)
}

//permet de modifier la valeur enregistré dans un la racine 
//d'un quadtree aux coordonnées (x, y) en "sol"
func (n *node) noeud_modif(x, y, sol int) {

	//gère les cas où la case cherchée n'existe pas
	if x < 0 || y < 0 || x >= n.width || y >= n.height {

		return 
	}

	//cas d'arrret
	if n.isLeaf {

		if n.width == 1 && n.height == 1 {

			n.content = sol
			return
		} else {

			n.isLeaf = false
			n.topLeftNode = &(node{topLeftX: n.topLeftX, topLeftY: n.topLeftY, width: n.width/2, height: n.height/2, content: n.content, isLeaf: true})
			n.topRightNode = &(node{topLeftX: n.topLeftX + n.width/2, topLeftY: n.topLeftY, width: n.width - n.width/2, height: n.height/2, content: n.content, isLeaf: true})
			n.bottomLeftNode = &(node{topLeftX: n.topLeftX, topLeftY: n.topLeftY + n.height/2, width: n.width/2, height: n.height - n.height/2, content: n.content, isLeaf: true})
			n.bottomRightNode = &(node{topLeftX: n.topLeftX + n.width/2, topLeftY: n.topLeftY + n.height/2, width: n.width - n.width/2, height: n.height - n.height/2, content: n.content, isLeaf: true})
		}
	}


	if x < n.width/2 {

		if y < n.height/2 {

			n.topLeftNode.noeud_modif(x, y, sol)
		} else {

			n.bottomLeftNode.noeud_modif(x, y-n.height/2, sol)
		}
	} else {

		if y < n.height/2 {

			n.topRightNode.noeud_modif(x-n.width/2, y, sol)
		} else {

			n.bottomRightNode.noeud_modif(x-n.width/2, y-n.height/2, sol)
		}
	}
}