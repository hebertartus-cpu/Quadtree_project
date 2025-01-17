package quadtree

import (
	"testing"
)


// fichier qui permet de tester la méthode valeurQuad
//a compléter

/*
var quad1 Quadtree = Quadtree{width: 2, height: 2,
	root: &node{topLeftX: 0, topLeftY: 0, width: 2, height: 2,
		topLeftNode:     &node{topLeftX: 0, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 1},
		topRightNode:    &node{topLeftX: 1, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 2},
		bottomLeftNode:  &node{topLeftX: 0, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 3},
		bottomRightNode: &node{topLeftX: 1, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 4},
	}}
*/

//teste la méthode ValeurQuad
func TestValQaudSimple(t *testing.T) {

	if quad1.ValeurQuad(0, 0) != 1 {
	
		t.Error("faux")
	}
		
	if quad1.ValeurQuad(1, 0) != 2 {
	
		t.Error("faux")
	}
		
	if quad1.ValeurQuad(0, 1) != 3 {
	
		t.Error("faux")
	}
		
	if quad1.ValeurQuad(1, 1) != 4 {
	
		t.Error("faux")
	}
}


//teste la méthode Modif_case
func TestModifCase(t *testing.T) {

	(quad1).Modif_Case(1, 0, 5)

	if  quad1.ValeurQuad(1, 0) != 5 {

		t.Error("faux")
	}
}

