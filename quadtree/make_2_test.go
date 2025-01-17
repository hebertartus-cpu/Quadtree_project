package quadtree

import (
	"reflect"
	"testing"
)

//teste les fonctions auxilières du fichier Make.go, nécessaire à l'exécution de MakeFromArray

// teste la fonction faire_portion
func TestFairePortion00(t *testing.T) {

	var tableau [][]int = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}

	var soluce [][]int = [][]int{
		{1, 2},
		{5, 6}}

	result := faire_portion(tableau, 0, 0)

	if !reflect.DeepEqual(result, soluce) {

		t.Error("faux")
	}
}

func TestFairePortion10(t *testing.T) {

	var tableau [][]int = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}

	var soluce [][]int = [][]int{
		{9, 10},
		{13, 14}}

	result := faire_portion(tableau, 1, 0)

	if !reflect.DeepEqual(result, soluce) {

		t.Error("faux")
	}
}

func TestFairePortion01(t *testing.T) {

	var tableau [][]int = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}

	var soluce [][]int = [][]int{
		{3, 4},
		{7, 8}}

	result := faire_portion(tableau, 0, 1)

	if !reflect.DeepEqual(result, soluce) {

		t.Error("faux")
	}
}

func TestFairePortion11(t *testing.T) {

	var tableau [][]int = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}

	var soluce [][]int = [][]int{
		{11, 12},
		{15, 16}}

	result := faire_portion(tableau, 1, 1)

	if !reflect.DeepEqual(result, soluce) {

		t.Error("faux")
	}
}

func TestFairePortionVecteurColonne(t *testing.T) {

	//teste si le tableau initiale ne fait que 1 de large

	var tableau [][]int = [][]int{
		{1},
		{2},
		{3},
		{4}}

	var soluce2 [][]int = [][]int{
		{1},
		{2}}

	var soluce1 [][]int = nil

	result := faire_portion(tableau, 0, 0)
	if !reflect.DeepEqual(result, soluce1) {

		t.Error("faux1")
	}

	result = faire_portion(tableau, 0, 1)
	if !reflect.DeepEqual(result, soluce2) {

		t.Error("faux2")
	}
}

func TestFairePortionVecteurLigne(t *testing.T) {

	//teste si le tableau initiale ne fait que 1 de long

	var tableau [][]int = [][]int{{1, 2, 3, 4}}

	var soluce2 [][]int = [][]int{{1, 2}}

	var soluce1 [][]int = nil

	result := faire_portion(tableau, 0, 0)
	if !reflect.DeepEqual(result, soluce1) {

		t.Error("faux1")
	}

	result = faire_portion(tableau, 1, 0)
	if !reflect.DeepEqual(result, soluce2) {

		t.Error("faux2")
	}
}

func TestFairePortionVecteurPoint(t *testing.T) {

	//teste si le tableau initial ne fait que 1 de long et de large

	var tableau [][]int = [][]int{{1}}

	var soluce1 [][]int = [][]int{{1}}

	var soluce2 [][]int = nil

	result := faire_portion(tableau, 0, 0)
	if !reflect.DeepEqual(result, soluce2) {

		t.Error("faux1")
	}

	result = faire_portion(tableau, 1, 0)
	if !reflect.DeepEqual(result, soluce2) {

		t.Error("faux2")
	}

	result = faire_portion(tableau, 0, 1)
	if !reflect.DeepEqual(result, soluce2) {

		t.Error("faux3")
	}

	result = faire_portion(tableau, 1, 1)
	if !reflect.DeepEqual(result, soluce1) {

		t.Error("faux4")
	}
}

func TestFairePortionImpairLarge(t *testing.T) {

	var tableau [][]int = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12}}

	var soluce [][]int = [][]int{
		{7},
		{10}}

	result := faire_portion(tableau, 1, 0)
	if !reflect.DeepEqual(result, soluce) {

		t.Error("faux")
	}
}

func TestFairePortionImpairLong(t *testing.T) {

	var tableau [][]int = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12}}

	var soluce [][]int = [][]int{
		{3, 4}}

	result := faire_portion(tableau, 0, 1)
	if !reflect.DeepEqual(result, soluce) {

		t.Error("faux")
	}
}

func TestFairePortionImpairDouble(t *testing.T) {

	var tableau [][]int = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}

	var soluce1 [][]int = [][]int{
		{4},
		{7}}

	var soluce2 [][]int = [][]int{
		{2, 3}}

	var soluce3 [][]int = [][]int{
		{1}}

	var soluce4 [][]int = [][]int{
		{5, 6},
		{8, 9}}

	result := faire_portion(tableau, 1, 0)
	if !reflect.DeepEqual(result, soluce1) {
		t.Error("faux1")
	}

	result = faire_portion(tableau, 0, 1)
	if !reflect.DeepEqual(result, soluce2) {

		t.Error("faux2")
	}

	result = faire_portion(tableau, 0, 0)
	if !reflect.DeepEqual(result, soluce3) {

		t.Error("faux3")
	}

	result = faire_portion(tableau, 1, 1)
	if !reflect.DeepEqual(result, soluce4) {

		t.Error("faux4")
	}
}

//teste la fonction esthomogene
// un retour -2 signifie faulse. sinon, renvoie ce que devrait devenir le content (le type de sol)

func TestEstHomogneFaux(t *testing.T) {

	var quad Quadtree = Quadtree{width: 2, height: 2,
		root: &node{topLeftX: 0, topLeftY: 0, width: 2, height: 2,
			topLeftNode:     &node{topLeftX: 0, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 0},
			topRightNode:    &node{topLeftX: 1, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 1},
			bottomLeftNode:  &node{topLeftX: 0, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 2},
			bottomRightNode: &node{topLeftX: 1, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 3},
		},
	}

	if est_homogene(quad.root) != -2 {

		t.Error("faux")
	}
}

func TestEstHomogneVrai(t *testing.T) {

	var quad Quadtree = Quadtree{width: 2, height: 2,
		root: &node{topLeftX: 0, topLeftY: 0, width: 2, height: 2,
			topLeftNode:     &node{topLeftX: 0, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 1},
			topRightNode:    &node{topLeftX: 1, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 1},
			bottomLeftNode:  &node{topLeftX: 0, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 1},
			bottomRightNode: &node{topLeftX: 1, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 1},
		},
	}

	if est_homogene(quad.root) != 1 {

		t.Error("faux")
	}
}

func TestEstHomogneNooeudvide1(t *testing.T) {

	var quad Quadtree = Quadtree{width: 2, height: 1,
		root: &node{topLeftX: 0, topLeftY: 0, width: 2, height: 1,
			topLeftNode:     &node{topLeftX: 0, topLeftY: 1, width: 0, height: 1, isLeaf: true, content: 0},
			topRightNode:    &node{topLeftX: 1, topLeftY: 1, width: 0, height: 1, isLeaf: true, content: 0},
			bottomLeftNode:  &node{topLeftX: 0, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 2},
			bottomRightNode: &node{topLeftX: 1, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 2},
		},
	}

	if est_homogene(quad.root) != 2 {

		t.Error("faux")
	}
}

func TestEstHomogneNooeudvide2(t *testing.T) {

	var quad Quadtree = Quadtree{width: 1, height: 2,
		root: &node{topLeftX: 0, topLeftY: 0, width: 1, height: 2,
			topLeftNode:     &node{topLeftX: 0, topLeftY: 1, width: 0, height: 1, isLeaf: true, content: 0},
			topRightNode:    &node{topLeftX: 0, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 2},
			bottomLeftNode:  &node{topLeftX: 1, topLeftY: 1, width: 0, height: 1, isLeaf: true, content: 0},
			bottomRightNode: &node{topLeftX: 1, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 2},
		},
	}

	if est_homogene(quad.root) != 2 {

		t.Error("faux")
	}
}

func TestEstHomogneNooeudvide3(t *testing.T) {

	var quad Quadtree = Quadtree{width: 1, height: 2,
		root: &node{topLeftX: 0, topLeftY: 0, width: 1, height: 2,
			topLeftNode:     &node{topLeftX: 0, topLeftY: 1, width: 0, height: 1, isLeaf: true, content: 0},
			topRightNode:    &node{topLeftX: 0, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 3},
			bottomLeftNode:  &node{topLeftX: 1, topLeftY: 1, width: 0, height: 1, isLeaf: true, content: 0},
			bottomRightNode: &node{topLeftX: 1, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 2},
		},
	}

	if est_homogene(quad.root) != -2 {

		t.Error("faux")
	}
}

//cas où tous les noeus sont vides
func TestEstHomogneNooeudvideAll(t *testing.T) {

	var quad Quadtree = Quadtree{width: 0, height: 0,
		root: &node{topLeftX: 0, topLeftY: 0, width: 0, height: 0,
			topLeftNode:     &node{topLeftX: 0, topLeftY: 1, width: 0, height: 1, isLeaf: true, content: 0},
			topRightNode:    &node{topLeftX: 0, topLeftY: 1, width: 1, height: 0, isLeaf: true, content: 0},
			bottomLeftNode:  &node{topLeftX: 1, topLeftY: 1, width: 0, height: 1, isLeaf: true, content: 0},
			bottomRightNode: &node{topLeftX: 1, topLeftY: 1, width: 0, height: 0, isLeaf: true, content: 0},
		},
	}

	if est_homogene(quad.root) != 0 {

		t.Error("faux")
	}
}

func TestEstHomogneQuadvide1(t *testing.T) {

	var quad Quadtree = Quadtree{width: 0, height: 0,
		root: &node{topLeftX: 0, topLeftY: 0, width: 0, height: 0, isLeaf: true},
	}

	if est_homogene(quad.root) != -2 {

		t.Error("faux")
	}
}

func TestEstHomogneQuadvide2(t *testing.T) {

	var quad Quadtree = Quadtree{width: 1, height: 1,
		root: &node{topLeftX: 0, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 3},
	}

	if est_homogene(quad.root) != -2 {

		t.Error("faux")
	}
}

//teste la fonction faire_noeud_vide
//

func TestFaireNoeudVideLigne(t *testing.T) {

	var tableau [][]int = [][]int{{1, 2, 3}}

	var noeud_result1 node = node{topLeftX: 0, topLeftY: 0, width: 1, height: 0, content: -2, isLeaf: true}
	var noeud_result2 node = node{topLeftX: 0, topLeftY: 0, width: 2, height: 0, content: -2, isLeaf: true}

	if !equalNode(faire_noeud_vide(tableau, 0, 0), &noeud_result1, true) {

		t.Error("faux")
	}

	if !equalNode(faire_noeud_vide(tableau, 0, 1), &noeud_result2, true) {

		t.Error("faux")
	}
}

func TestFaireNoeudVideColonne(t *testing.T) {

	var tableau [][]int = [][]int{
		{1},
		{2},
		{3}}

	var noeud_result1 node = node{topLeftX: 0, topLeftY: 0, width: 0, height: 1, content: -2, isLeaf: true}
	var noeud_result2 node = node{topLeftX: 0, topLeftY: 0, width: 0, height: 2, content: -2, isLeaf: true}

	if !equalNode(faire_noeud_vide(tableau, 0, 0), &noeud_result1, true) {

		t.Error("faux")
	}

	if !equalNode(faire_noeud_vide(tableau, 1, 0), &noeud_result2, true) {

		t.Error("faux")
	}
}

func TestFaireNoeudVidePoint(t *testing.T) {

	var tableau [][]int = [][]int{{1}}

	var noeud_result1 node = node{topLeftX: 0, topLeftY: 0, width: 0, height: 0, content: 0, isLeaf: true}

	if !equalNode(faire_noeud_vide(tableau, 0, 0), &noeud_result1, true) {

		t.Error("faux")
	}
}

func TestFaireNoeudVideTabvide(t *testing.T) {

	var tableau [][]int = [][]int{{}}

	var noeud_result1 node = node{topLeftX: 0, topLeftY: 0, width: 0, height: 0, content: 0, isLeaf: true}

	if !equalNode(faire_noeud_vide(tableau, 0, 0), &noeud_result1, true) {

		t.Error("faux")
	}
}

// teste la méthode actualise_co
//

func TestActualiseCoSimple1(t *testing.T) {

	var quad Quadtree = Quadtree{width: 1, height: 1,
		root: &node{topLeftX: 2, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 2}}

	var quad_result Quadtree = Quadtree{width: 1, height: 1,
		root: &node{topLeftX: 3, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 2}}

	quad.root.actualise_co(1, 0)

	if !equal(quad, quad_result, true) {

		t.Error("faux")
	}
}

func TestActualiseCoSimple2(t *testing.T) {

	var quad Quadtree = Quadtree{width: 1, height: 1,
		root: &node{topLeftX: 4, topLeftY: 7, width: 1, height: 1, isLeaf: true, content: 2}}

	var quad_result Quadtree = Quadtree{width: 1, height: 1,
		root: &node{topLeftX: 6, topLeftY: 13, width: 1, height: 1, isLeaf: true, content: 2}}

	quad.root.actualise_co(2, 6)

	if !equal(quad, quad_result, true) {

		t.Error("faux")
	}
}

func TestActualiseCoPlainX(t *testing.T) {

	var quad Quadtree = Quadtree{width: 2, height: 2,
		root: &node{topLeftX: 0, topLeftY: 0, width: 2, height: 2,
			topLeftNode:     &node{topLeftX: 0, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 0},
			topRightNode:    &node{topLeftX: 0, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 3},
			bottomLeftNode:  &node{topLeftX: 1, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 0},
			bottomRightNode: &node{topLeftX: 1, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 2},
		},
	}

	var quad_result Quadtree = Quadtree{width: 2, height: 2,
		root: &node{topLeftX: 2, topLeftY: 0, width: 2, height: 2,
			topLeftNode:     &node{topLeftX: 2, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 0},
			topRightNode:    &node{topLeftX: 2, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 3},
			bottomLeftNode:  &node{topLeftX: 3, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 0},
			bottomRightNode: &node{topLeftX: 3, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 2},
		},
	}

	quad.root.actualise_co(2, 0)

	if !equal(quad, quad_result, true) {

		t.Error("faux")
	}
}

func TestActualiseCoPlainY(t *testing.T) {

	var quad Quadtree = Quadtree{width: 2, height: 2,
		root: &node{topLeftX: 0, topLeftY: 0, width: 2, height: 2,
			topLeftNode:     &node{topLeftX: 0, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 0},
			topRightNode:    &node{topLeftX: 0, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 3},
			bottomLeftNode:  &node{topLeftX: 1, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 0},
			bottomRightNode: &node{topLeftX: 1, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 2},
		},
	}

	var quad_result Quadtree = Quadtree{width: 2, height: 2,
		root: &node{topLeftX: 0, topLeftY: 1, width: 2, height: 2,
			topLeftNode:     &node{topLeftX: 0, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 0},
			topRightNode:    &node{topLeftX: 0, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 3},
			bottomLeftNode:  &node{topLeftX: 1, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 0},
			bottomRightNode: &node{topLeftX: 1, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 2},
		},
	}

	quad.root.actualise_co(0, 1)

	if !equal(quad, quad_result, true) {

		t.Error("faux")
	}
}

func TestActualiseCoSimpleNoeudvide(t *testing.T) {

	var quad Quadtree = Quadtree{width: 0, height: 0,
		root: &node{topLeftX: 4, topLeftY: 7, width: 0, height: 1, isLeaf: true, content: 0}}

	var quad_result Quadtree = Quadtree{width: 0, height: 0,
		root: &node{topLeftX: 21, topLeftY: 9, width: 0, height: 1, isLeaf: true, content: 0}}
	
	quad.root.actualise_co(17, 2)

	if !equal(quad, quad_result, true) {

		t.Error("faux")
	}
}
