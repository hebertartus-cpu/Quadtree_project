package quadtree

import "testing"


// fichier qui permet de tester la méthode trouve_valeur

var quad1 Quadtree = Quadtree{width: 2, height: 2,
	root: &node{topLeftX: 0, topLeftY: 0, width: 2, height: 2,
		topLeftNode:     &node{topLeftX: 0, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 1},
		topRightNode:    &node{topLeftX: 1, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 2},
		bottomLeftNode:  &node{topLeftX: 0, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 3},
		bottomRightNode: &node{topLeftX: 1, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 4},
	}}

var quad2 Quadtree = Quadtree{width: 6, height: 5,
	root: &node{topLeftX: 0, topLeftY: 0, width: 6, height: 5,
		topLeftNode: &node{topLeftX: 0, topLeftY: 0, width: 3, height: 2,
			topLeftNode: &node{topLeftX: 0, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 1},
			topRightNode: &node{topLeftX: 1, topLeftY: 0, width: 2, height: 1,
				topLeftNode:     &node{topLeftX: 1, topLeftY: 0, width: 1, height: 0, isLeaf: true, content: 0},
				topRightNode:    &node{topLeftX: 2, topLeftY: 0, width: 1, height: 0, isLeaf: true, content: 0},
				bottomLeftNode:  &node{topLeftX: 1, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 2},
				bottomRightNode: &node{topLeftX: 2, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 3},
			},
			bottomLeftNode: &node{topLeftX: 0, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 2},
			bottomRightNode: &node{topLeftX: 1, topLeftY: 1, width: 2, height: 1,
				topLeftNode:     &node{topLeftX: 1, topLeftY: 1, width: 1, height: 0, isLeaf: true, content: 0},
				topRightNode:    &node{topLeftX: 2, topLeftY: 1, width: 1, height: 0, isLeaf: true, content: 0},
				bottomLeftNode:  &node{topLeftX: 1, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 3},
				bottomRightNode: &node{topLeftX: 2, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 4},
			},
		},
		topRightNode: &node{topLeftX: 3, topLeftY: 0, width: 3, height: 2,
			topLeftNode: &node{topLeftX: 3, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 4},
			topRightNode: &node{topLeftX: 4, topLeftY: 0, width: 2, height: 1,
				topLeftNode:     &node{topLeftX: 4, topLeftY: 0, width: 1, height: 0, isLeaf: true, content: 0},
				topRightNode:    &node{topLeftX: 5, topLeftY: 0, width: 1, height: 0, isLeaf: true, content: 0},
				bottomLeftNode:  &node{topLeftX: 4, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 1},
				bottomRightNode: &node{topLeftX: 5, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 2},
			},
			bottomLeftNode: &node{topLeftX: 3, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 1},
			bottomRightNode: &node{topLeftX: 4, topLeftY: 1, width: 2, height: 1,
				topLeftNode:     &node{topLeftX: 4, topLeftY: 1, width: 1, height: 0, isLeaf: true, content: 0},
				topRightNode:    &node{topLeftX: 5, topLeftY: 1, width: 1, height: 0, isLeaf: true, content: 0},
				bottomLeftNode:  &node{topLeftX: 4, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 2},
				bottomRightNode: &node{topLeftX: 5, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 3},
			},
		},
		bottomLeftNode: &node{topLeftX: 0, topLeftY: 2, width: 3, height: 3,
			topLeftNode: &node{topLeftX: 0, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 3},
			topRightNode: &node{topLeftX: 1, topLeftY: 2, width: 2, height: 1,
				topLeftNode:     &node{topLeftX: 1, topLeftY: 2, width: 1, height: 0, isLeaf: true, content: 0},
				topRightNode:    &node{topLeftX: 2, topLeftY: 2, width: 1, height: 0, isLeaf: true, content: 0},
				bottomLeftNode:  &node{topLeftX: 1, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 4},
				bottomRightNode: &node{topLeftX: 2, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 1},
			},
			bottomLeftNode: &node{topLeftX: 0, topLeftY: 3, width: 1, height: 2,
				topLeftNode:     &node{topLeftX: 0, topLeftY: 3, width: 0, height: 1, isLeaf: true, content: 0},
				topRightNode:    &node{topLeftX: 0, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 4},
				bottomLeftNode:  &node{topLeftX: 0, topLeftY: 4, width: 0, height: 1, isLeaf: true, content: 0},
				bottomRightNode: &node{topLeftX: 0, topLeftY: 4, width: 1, height: 1, isLeaf: true, content: 1},
			},
			bottomRightNode: &node{topLeftX: 1, topLeftY: 3, width: 2, height: 2,
				topLeftNode:     &node{topLeftX: 1, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 1},
				topRightNode:    &node{topLeftX: 2, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 2},
				bottomLeftNode:  &node{topLeftX: 1, topLeftY: 4, width: 1, height: 1, isLeaf: true, content: 2},
				bottomRightNode: &node{topLeftX: 2, topLeftY: 4, width: 1, height: 1, isLeaf: true, content: 3},
			},
		},
		bottomRightNode: &node{topLeftX: 3, topLeftY: 2, width: 3, height: 3,
			topLeftNode: &node{topLeftX: 3, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 2},
			topRightNode: &node{topLeftX: 4, topLeftY: 2, width: 2, height: 1,
				topLeftNode:     &node{topLeftX: 4, topLeftY: 2, width: 1, height: 0, isLeaf: true, content: 0},
				topRightNode:    &node{topLeftX: 5, topLeftY: 2, width: 1, height: 0, isLeaf: true, content: 0},
				bottomLeftNode:  &node{topLeftX: 4, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 3},
				bottomRightNode: &node{topLeftX: 5, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 4},
			},
			bottomLeftNode: &node{topLeftX: 3, topLeftY: 3, width: 1, height: 2,
				topLeftNode:     &node{topLeftX: 3, topLeftY: 3, width: 0, height: 1, isLeaf: true, content: 0},
				topRightNode:    &node{topLeftX: 3, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 3},
				bottomLeftNode:  &node{topLeftX: 3, topLeftY: 4, width: 0, height: 1, isLeaf: true, content: 0},
				bottomRightNode: &node{topLeftX: 3, topLeftY: 4, width: 1, height: 1, isLeaf: true, content: 4},
			},
			bottomRightNode: &node{topLeftX: 4, topLeftY: 3, width: 2, height: 2,
				topLeftNode:     &node{topLeftX: 4, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 4},
				topRightNode:    &node{topLeftX: 5, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 1},
				bottomLeftNode:  &node{topLeftX: 4, topLeftY: 4, width: 1, height: 1, isLeaf: true, content: 1},
				bottomRightNode: &node{topLeftX: 5, topLeftY: 4, width: 1, height: 1, isLeaf: true, content: 2},
			},
		},
	},
}

var quad3 Quadtree = Quadtree{width: 0, height: 0, 
	root: &node{topLeftX: 0, topLeftY: 0, width: 0, height: 0}}


//teste la méthode trouve_valeur, qui à partir d'un x et d'un y, 
//renvoie la valeur contenue dans le noeud d'un quadtree aux cohordonnées sus dites
func TestTtrouveValSimple(t *testing.T) {

	if quad1.root.trouve_valeur(0, 0) != 1 {

		t.Error("faux")
	}
	
	if quad1.root.trouve_valeur(1, 0) != 2 {

		t.Error("faux")
	}
	
	if quad1.root.trouve_valeur(0, 1) != 3 {

		t.Error("faux")
	}
	
	if quad1.root.trouve_valeur(1, 1) != 4 {

		t.Error("faux")
	}
}

func TestTtrouveValRect(t *testing.T) {

	if quad2.root.trouve_valeur(3, 4) != 4 {

		t.Error("faux")
	}

	if quad2.root.trouve_valeur(1, 0) != 2 {

		t.Error("faux")
	}

	if quad2.root.trouve_valeur(5, 4) != 2 {

		t.Error("faux")
	}

	if quad2.root.trouve_valeur(2, 3) != 2 {

		t.Error("faux")
	}

	if quad2.root.trouve_valeur(3, 2) != 2 {

		t.Error("faux")
	}

	if quad2.root.trouve_valeur(0, 3) != 4 {

		t.Error("faux")
	}
}

func TestTtrouveValVide(t *testing.T) {

	if quad3.root.trouve_valeur(0, 0) != -2 {

		t.Error("faux")
	}

	if quad1.root.trouve_valeur(-1, 0) != -2 {

		t.Error("faux")
	}

	if quad1.root.trouve_valeur(0, 5) != -2 {

		t.Error("faux")
	}

	if quad2.root.trouve_valeur(2, -3) != -2 {

		t.Error("faux")
	}

	if quad2.root.trouve_valeur(40, -345) != -2 {

		t.Error("faux")
	}
}