package quadtree

import (
	"fmt"
	
)

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

//renvoie la racine d'un quadtree
func (q Quadtree) Root() *node {

	return q.root
}

//permet d'afficher des info utile sur les quadtree
func (q Quadtree) Affiche_info() {

	fmt.Println("quad width", q.width)
	fmt.Println("quad height", q.height)
	fmt.Println("quad MinX", q.MinX)
	fmt.Println("quad MinY", q.MinY)
	fmt.Println("root topX", q.root.topLeftX)
	fmt.Println("root topy", q.root.topLeftY)
	fmt.Println("root width", q.root.width)
	fmt.Println("root height", q.root.height)
	if !q.root.isLeaf {

		fmt.Println("topLeft", q.root.topLeftNode.topLeftX, "/", q.root.topLeftNode.topLeftY)
		fmt.Println("topRight", q.root.topRightNode.topLeftX, "/", q.root.topRightNode.topLeftY)
		fmt.Println("bottomLeft", q.root.bottomLeftNode.topLeftX, "/", q.root.bottomLeftNode.topLeftY)
		fmt.Println("bottomRight", q.root.bottomRightNode.topLeftX, "/", q.root.bottomRightNode.topLeftY)
		fmt.Println("topLeft", q.root.topLeftNode.isLeaf)
		fmt.Println("topRight", q.root.topRightNode.isLeaf)
		fmt.Println("bottomLeft", q.root.bottomLeftNode.isLeaf)
		fmt.Println("bottomRight", q.root.bottomRightNode.isLeaf)
	}
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

	//permet de gérer les cas où les coordonnées sont négatives
	negatifX := 0
	negatifY := 0
	//annule le décalage des coordonnées qi nécessaire
	compenseX := 0
	compenseY := 0

	if n.topLeftX < 0 {

		negatifX = n.topLeftX
		compenseX = n.width/2
	}
	if n.topLeftY < 0 {

		negatifY = n.topLeftY
		compenseY = n.height/2
	}


	if x < n.width/2 + negatifX {

		if y < n.height/2 + negatifY {

			n.topLeftNode.noeud_modif(x, y, sol)
		} else {

			n.bottomLeftNode.noeud_modif(x, y-n.height/2 + compenseY, sol)
		}
	} else {

		if y < n.height/2 + negatifY {

			n.topRightNode.noeud_modif(x-n.width/2 + compenseX, y, sol)
		} else {

			n.bottomRightNode.noeud_modif(x-n.width/2 + compenseX, y-n.height/2 + compenseY, sol)
		}
	}
}


//Méthode qui agrandi le quadtree en le multipliant par 4 (longeur *2 + largeur *2)
//L'agrandit dans des directions différente en fonction de verif:
//1 ou 2 -> agrandit vers en bas à droite
//3 ou 4 -> agrandit vers en heut à gauche
func (q *Quadtree) Agrandir(verif int) {

	//vers la droite ou vers le bas
	if verif == 1 || verif == 2 {

		//créé le noeud de taille supérieure
		noeud := node{topLeftX: q.root.topLeftX, topLeftY: q.root.topLeftY, width: q.width *2, height: q.height *2, content: -2, isLeaf: false,

			//créé les sous noeud associé, avec le noeud actuel en topleftNode
			topLeftNode: q.root,
			topRightNode: &(node{topLeftX: q.root.topLeftX + q.width, topLeftY: q.root.topLeftY, width: q.width, height: q.height, content: -3, isLeaf: true}),
			bottomLeftNode: &(node{topLeftX: q.root.topLeftX, topLeftY: q.root.topLeftY + q.height, width: q.width, height: q.height, content: -3, isLeaf: true}),
			bottomRightNode: &(node{topLeftX: q.root.topLeftX + q.width, topLeftY: q.root.topLeftY + q.height, width: q.width, height: q.height, content: -3, isLeaf: true}),
		}
		
		//Modifie la taille du quadtree pour corespondre
		q.width = q.width *2
		q.height = q.height *2
		
		//Met le nouveau noeud de taille supérieure dans le quadtree
		q.root = &noeud
	
	//vers la gauche ou vers le haut
	} else if verif == 3 || verif == 4 {
		
		//créé le noeud de taille supérieure
		noeud := node{topLeftX: q.root.topLeftX - q.width, topLeftY: q.root.topLeftY - q.height, width: q.width *2, height: q.height *2, content: -2, isLeaf: false,

			//créé les sous noeud associé, avec le noeud actuel en topleftNode
			topLeftNode: &(node{topLeftX: q.root.topLeftX - q.width, topLeftY: q.root.topLeftY - q.height, width: q.width, height: q.height, content: -3, isLeaf: true}),
			topRightNode: &(node{topLeftX: q.root.topLeftX, topLeftY: q.root.topLeftY - q.height, width: q.width, height: q.height, content: -3, isLeaf: true}),
			bottomLeftNode: &(node{topLeftX: q.root.topLeftX - q.width, topLeftY: q.root.topLeftY, width: q.width, height: q.height, content: -3, isLeaf: true}),
			bottomRightNode: q.root,
		}

		//enregistre les coordonnée en haut à droite du quadtree, désormais négatives
		q.MinX -= q.width
		q.MinY -= q.height

		//Modifie la taille du quadtree pour corespondre
		q.width = q.width *2
		q.height = q.height *2

		//Met le nouveau noeud de taille supérieure dans le quadtree
		q.root = &noeud
	}
}

//! A n'utiliser que dans le cadre de la fontion game.lance_generation !
//Fonctions qui modifient un noeud au coordonnée (x, y) et qui fait la taille (w, h)
//par le noeud donné en paramêtre
//
//la première fonction permet l'activation à partir d'un quadtree et non d'un noeud
func (q *Quadtree) Modif_zone(x, y, w, h int, noeud *node) {

	q.root.Modif_zone_noeud(x, y, w, h, noeud)
}


func (n *node) Modif_zone_noeud(x, y, w, h int, noeud *node) {
	
	//s'il faut couper le quadtree
	if n.isLeaf {

		n.content = -2
		n.isLeaf = false

		n.topLeftNode = &(node{topLeftX: n.topLeftX, topLeftY: n.topLeftY, width: n.width /2, height: n.height /2, content: -3, isLeaf: true})
		n.topRightNode = &(node{topLeftX: n.topLeftX + n.width /2, topLeftY: n.topLeftY, width: n.width /2, height: n.height /2, content: -3, isLeaf: true})
		n.bottomLeftNode = &(node{topLeftX: n.topLeftX, topLeftY: n.topLeftY + n.height /2, width: n.width /2, height: n.height /2, content: -3, isLeaf: true})
		n.bottomRightNode = &(node{topLeftX: n.topLeftX + n.width /2, topLeftY: n.topLeftY + n.height /2, width: n.width /2, height: n.height /2, content: -3, isLeaf: true})
	}

	//permet de gérer les cas où les coordonnées sont négatives
	negatifX := 0
	negatifY := 0
	//annule le décalage des coordonnées qi nécessaire
	compenseX := 0
	compenseY := 0

	if n.topLeftX < 0 {

		negatifX = n.topLeftX
		compenseX = n.width/2
	}
	if n.topLeftY < 0 {

		negatifY = n.topLeftY
		compenseY = n.height/2
	}

	//cas d'arret
	if n.width/2 == w && n.height/2 == h {

		//à gauche
		if x < n.width/2 + negatifX {

			//en haut
			if y < n.height/2 + negatifY {
				
				noeud.actualise_co(n.topLeftNode.topLeftX, n.topLeftNode.topLeftY)
				n.topLeftNode = noeud
			
			//en bas
			} else {
	
				noeud.actualise_co(n.bottomLeftNode.topLeftX, n.bottomLeftNode.topLeftY)
				n.bottomLeftNode = noeud
			}
		
		//à droite
		} else {
			
			//en haut
			if y < n.height/2 + negatifY {
	
				noeud.actualise_co(n.topRightNode.topLeftX, n.topRightNode.topLeftY)
				n.topRightNode = noeud
			
			//en bas
			} else {
	
				noeud.actualise_co(n.bottomRightNode.topLeftX, n.bottomRightNode.topLeftY)
				n.bottomRightNode = noeud
			}
		}
	
	//progression dans l'arborecence
	} else {

		if x < n.width/2 + negatifX {

			if y < n.height/2 + negatifY {
	
				n.topLeftNode.Modif_zone_noeud(x, y, w, h, noeud)
			} else {
	
				n.bottomLeftNode.Modif_zone_noeud(x, y-n.height/2 + compenseY, w, h, noeud)
			}
		} else {
	
			if y < n.height/2 + negatifY {
	
				n.topRightNode.Modif_zone_noeud(x-n.width/2 + compenseX, y, w, h, noeud)
			} else {
	
				n.bottomRightNode.Modif_zone_noeud(x-n.width/2 + compenseX, y-n.height/2 + compenseY, w, h, noeud)
			}
		}
	}
}
