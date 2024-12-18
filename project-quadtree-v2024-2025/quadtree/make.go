package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	// TODO

	//gère le cas où floorContent est vide
	if len(floorContent) == 0 {

		return Quadtree{width: 0, height: 0, 
			root: &node{topLeftX: 0, topLeftY: 0, width: 0, height: 0, isLeaf: true, content: 0}}
	}

	//Cas d'arret
	if len(floorContent) == 1 && len(floorContent[0]) == 1 {

		//renvoie un arbre qui est une feuille
		var tree Quadtree = Quadtree{width: 1, height: 1,
			root: &node{topLeftX: 0, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: floorContent[0][0]},
		}

		return tree
	}

	//initialisation du quadtree à renvoyer
	var tree Quadtree = Quadtree{width: len(floorContent[0]), height: len(floorContent),

		root: &node{topLeftX: 0, topLeftY: 0, width: len(floorContent[0]), height: len(floorContent), content: -2, isLeaf: false},
	}

	//creation récursive des noeuds
	var portion [][]int

	for i := 0; i < 2; i++ {

		for j := 0; j < 2; j++ {

			//coupe le tableau floorContent en 4 tableau plus petit
			portion = faire_portion(floorContent, i, j)

			if i == 0 {

				if j == 0 {

					if portion != nil {

						tree.root.topLeftNode = MakeFromArray(portion).root
					} else {

						//gère les cas où le noeud est "vide"/n'existe pas
						tree.root.topLeftNode = faire_noeud_vide(floorContent, i, j)
					}

				} else {

					if portion != nil {

						tree.root.topRightNode = MakeFromArray(portion).root
					} else {

						//gère les cas où le noeud est "vide"/n'existe pas
						tree.root.topRightNode = faire_noeud_vide(floorContent, i, j)
					}

					tree.root.topRightNode.actualise_co(tree.width/2, 0)
				}
			} else {

				if j == 0 {

					if portion != nil {

						tree.root.bottomLeftNode = MakeFromArray(portion).root
					} else {

						//gère les cas où le noeud est "vide"/n'existe pas
						tree.root.bottomLeftNode = faire_noeud_vide(floorContent, i, j)
					}

					tree.root.bottomLeftNode.actualise_co(0, tree.height/2)

				} else {

					if portion != nil {

						tree.root.bottomRightNode = MakeFromArray(portion).root
					} else {

						//gère les cas où le noeud est "vide"/n'existe pas
						//en principe, n'est jamais activé
						tree.root.bottomRightNode = faire_noeud_vide(floorContent, i, j)
					}

					tree.root.bottomRightNode.actualise_co(tree.width/2, tree.height/2)
				}
			}
		}
	}

	//vérifie si tous les noeuds enfant sont ou non identiques et non vide
	//si est_homogene() renvoi -2, le noed n'est pas homogene
	var type_sol int = est_homogene(tree.root)
	if type_sol != -2 {

		//le noeud de l'arbre devient une feuille
		tree.root.isLeaf = true
		tree.root.content = type_sol

		//suprétion des noeud
		tree.root.topLeftNode = nil
		tree.root.topRightNode = nil
		tree.root.bottomLeftNode = nil
		tree.root.bottomRightNode = nil
	}

	return tree
}

// fonction qui dit si un noeud est homogène, càd si ses noeuds enfants sont tous des feuilles,
// et que les types de sol des noeuds existants (ceux dont ni la largeur ni la hauteur ne sont nul) sont tous les même.
// si le noeud n'est pas homogène, renvoie -2, sinon, le type de sol qui doit lui être attribué
func est_homogene(noeud *node) (contient int) {

	//permet de gérer le cas où le noeud est une feuille, ou un noeud qui n'existe pas (width ou height = 0)
	//si noeud est une feuille, renvoie -2, car le noeud n'a dans tous les cas pas besoins d'être reduit
	//en principe, ne s'active jamais quand apellé par MakeFromArray
	if noeud.isLeaf {

		return -2
	}
	
	contient = -2
	
	if noeud.topLeftNode.isLeaf && noeud.topRightNode.isLeaf && noeud.bottomLeftNode.isLeaf && noeud.bottomRightNode.isLeaf {

		var tab []*node
		var not_noeud int

		if !(noeud.topLeftNode.height == 0 || noeud.topLeftNode.width == 0) {

			tab = append(tab, noeud.topLeftNode)
		} else {

			not_noeud ++
		}

		if !(noeud.topRightNode.height == 0 || noeud.topRightNode.width == 0) {

			tab = append(tab, noeud.topRightNode)
		} else {

			not_noeud ++
		}

		if !(noeud.bottomLeftNode.height == 0 || noeud.bottomLeftNode.width == 0) {

			tab = append(tab, noeud.bottomLeftNode)
		} else {

			not_noeud ++
		}

		if !(noeud.bottomRightNode.height == 0 || noeud.bottomRightNode.width == 0) {

			tab = append(tab, noeud.bottomRightNode)
		} else {

			not_noeud ++
		}

		//verifie que tout les noeud valides ont le même content
		if len(tab) > 0 {

			contient = tab[0].content

			for i := 1; i < len(tab) && contient != -2; i++ {

				if tab[i].content != contient {

					contient = -2
				}
			}
		}

		//s'occupe du cas où tous les sous noeuds sont inexistant
		if not_noeud == 4 {

			contient = 0
		}
	}

	return contient
}

// créé un noeud inexisgtant (avec width ou height = 0) pour compléter le quadtree dans le cas d'un flooreContent rectangulaire
// il était initialment prévu que les noeuds inexistants soit égaux à nil
func faire_noeud_vide(floorContent [][]int, i, j int) (*node) {

	var largeure int
	var hauteur int

	//permet de gérer le cas ou floorContent est vide. En principe, n'arrive jamais
	if len(floorContent) != 0 {
		if i == 0 {

			if j == 0 {

				largeure = len(floorContent[0][:len(floorContent[0])/2])
				hauteur = len(floorContent[:len(floorContent)/2])
			} else {

				largeure = len(floorContent[0][len(floorContent[0])/2:])
				hauteur = len(floorContent[:len(floorContent)/2])
			}
		} else {

			if j == 0 {

				largeure = len(floorContent[0][:len(floorContent[0])/2])
				hauteur = len(floorContent[len(floorContent)/2:])
			//jamais utilisé au vu de comment sont construits les noeuds
			} else {

				largeure = len(floorContent[0][len(floorContent[0])/2:])
				hauteur = len(floorContent[len(floorContent)/2:])
			}
		}
	}

	var noeud node = node{width: largeure, height: hauteur, content: 0, isLeaf: true}

	return &noeud
}

// permet de mettre à jour topleftX et topleftY pour un noeud et tous ses enfants
func (noeud *node) actualise_co(large, long int) {

	//cas d'arrêt
	if noeud.isLeaf {

		noeud.topLeftX += large
		noeud.topLeftY += long
	
	} else {

		if noeud.topLeftNode != nil {

			noeud.topLeftNode.actualise_co(large, long)
		}

		if noeud.topRightNode != nil {

			noeud.topRightNode.actualise_co(large, long)
		}

		if noeud.bottomLeftNode != nil {

			noeud.bottomLeftNode.actualise_co(large, long)
		}

		if noeud.bottomRightNode != nil {

			noeud.bottomRightNode.actualise_co(large, long)
		}

		noeud.topLeftX += large
		noeud.topLeftY += long
	}
}

// permet de couper floorContent en plusieur parties en fonction du i et j donné
func faire_portion(floorContent [][]int, i, j int) (portion [][]int) {

	//sera utile pour ne pas modifier floorContent
	var p_portion [][]int

	//prend la partie haute ou la partie basse de floorContent
	if i == 0 {

		p_portion = floorContent[:len(floorContent)/2]
	} else {

		p_portion = floorContent[len(floorContent)/2:]
	}

	//permet de na pas modifier floorContent
	portion = make([][]int, len(p_portion))
	copy(portion, p_portion)

	//prend la partie gauche ou la partie droite de la portion
	if len(portion) != 0 {

		if j == 0 {

			for k := 0; k < len(portion); k++ {

				portion[k] = portion[k][:len(portion[k])/2]
			}
		} else {

			for k := 0; k < len(portion); k++ {

				portion[k] = portion[k][len(portion[k])/2:]

			}
		}
	}

	//suprime les lignes vides
	k := 0
	for k < len(portion) {

		if len(portion[k]) == 0 {

			portion = append(portion[:k], portion[k+1:]...)
			k--
		}
		k++
	}

	//met portion à nil si portion est vide
	if len(portion) == 0 {

		portion = nil
	}

	return portion
}
