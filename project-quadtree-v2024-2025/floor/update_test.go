package floor

import (
	"reflect"
	"testing"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

func TestUpdateFromFileExterne(t *testing.T) {
	///if comparaison entre ce que modifie la fonction update et un tableau content donné
	var f Floor
	configuration.Global.FloorFile = "../floor-files/TESTUPDATE"
	f.content = [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	vide := [][]int{{-1, -1, -1}, {-1, -1, -1}, {-1, -1, -1}}
	f.updateFromFileFloor(-5, -5)
	if !reflect.DeepEqual(f.content, vide) {
		t.Error("faux")
	}

}

func TestUpdateFromFileHautGauche(t *testing.T) {
	var f Floor
	configuration.Global.FloorFile = "../floor-files/TESTUPDATE"
	f.content = [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	hautgauche := [][]int{{-1, -1, -1}, {-1, 0, 0}, {-1, 0, 0}}
	f.updateFromFileFloor(-1, -1)
	if !reflect.DeepEqual(f.content, hautgauche) {
		t.Error("faux")
	}

}

func TestUpdateFromFileHautGauche2(t *testing.T) {
	var f Floor
	configuration.Global.FloorFile = "../floor-files/testupdate2"
	f.content = [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	hautgauche := [][]int{{-1, -1, -1, -1, -1}, {-1, 1, 2, 3, 4}, {-1, 1, 2, 3, 4}, {-1, 1, 2, 3, 4}, {-1, 1, 2, 3, 4}}
	f.updateFromFileFloor(-1, -1)
	if !reflect.DeepEqual(f.content, hautgauche) {
		t.Error("faux")
	}

}

func TestUpdateFromFileHaut(t *testing.T) {
	var f Floor
	configuration.Global.FloorFile = "../floor-files/TESTUPDATE"
	f.content = [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	haut := [][]int{{-1, -1, -1}, {0, 0, 0}, {0, 0, 0}}
	f.updateFromFileFloor(1, -1)
	if !reflect.DeepEqual(f.content, haut) {
		t.Error("faux")
	}

}

func TestUpdateFromFileBas(t *testing.T) {
	var f Floor
	configuration.Global.FloorFile = "../floor-files/TESTUPDATE"
	f.content = [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	bas := [][]int{{0, 0, 0}, {0, 0, 0}, {-1, -1, -1}}
	f.updateFromFileFloor(1, 4)
	if !reflect.DeepEqual(f.content, bas) {
		t.Error("faux")
	}
}

func TestUpdateFromFileGauche(t *testing.T) {
	var f Floor
	configuration.Global.FloorFile = "../floor-files/TESTUPDATE"
	f.content = [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	gauche := [][]int{{-1, 0, 0}, {-1, 0, 0}, {-1, 0, 0}}
	f.updateFromFileFloor(-1, 2)
	if !reflect.DeepEqual(f.content, gauche) {
		t.Error("faux")
	}

}

func TestUpdateFromFileDroite(t *testing.T) {
	var f Floor
	configuration.Global.FloorFile = "../floor-files/TESTUPDATE"
	f.content = [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	droite := [][]int{{0, 0, -1}, {0, 0, -1}, {0, 0, -1}}
	f.updateFromFileFloor(18, 1)
	if !reflect.DeepEqual(f.content, droite) {
		t.Error("faux")
	}
}

func TestUpdateFromFileMilieu(t *testing.T) {
	var f Floor
	configuration.Global.FloorFile = "../floor-files/testupdate2"
	f.content = [][]int{{0, 0, 5}, {0, 2, 0}, {0, 0, 0}}
	milieu := [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	f.updateFromFileFloor(0, 0)
	if !reflect.DeepEqual(f.content, milieu) {
		t.Error("faux")
	}

}

func TestUpdateFromFileVide(t *testing.T) {
	var f Floor
	configuration.Global.FloorFile = "../floor-files/vide"
	f.content = [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	result := [][]int{{-1, -1, -1}, {-1, -1, -1}, {-1, -1, -1}}
	f.updateFromFileFloor(0, 0)
	if !reflect.DeepEqual(f.content, result) {
		t.Error("faux")
	}
}

func TestUpdateFromFileCamDepasse(t *testing.T) {
	var f Floor
	configuration.Global.FloorFile = "../floor-files/exemple"
	f.content = [][]int{{0, 8, 0, 3, 5}, {0, 0, 0, 34, 1234}, {0, 276, 0, 5, 0}, {3, 52, 1, 1, 9}, {3, 0, 6, 1, 7}}
	result := [][]int{{1, 1, 3, 4, -1}, {1, 1, 4, 3, -1}, {0, 0, 2, 2, -1}, {0, 0, 2, 2, -1}, {-1, -1, -1, -1, -1}}
	f.updateFromFileFloor(0, 0)
	if !reflect.DeepEqual(f.content, result) {
		t.Error("faux")
	}
}

func TestUpdateFromFileCamRectangle1(t *testing.T) {
	var f Floor
	configuration.Global.FloorFile = "../floor-files/beaupasbeau"
	f.content = [][]int{{0, 8, 0, 3}, {0, 0, 0, 34}}
	result := [][]int{{1, 0, 0, 0}, {1, 1, 1, 0}}
	f.updateFromFileFloor(3, 3)
	if !reflect.DeepEqual(f.content, result) {
		t.Error("faux")
	}
}

func TestUpdateFromFileCamRectangle2(t *testing.T) {
	var f Floor
	configuration.Global.FloorFile = "../floor-files/beaupasbeau"
	f.content = [][]int{{0, 8}, {0, 0}, {0, 2}, {56, 0}, {4, 8}, {1, 1}, {1, -1}}
	result := [][]int{{2, 2}, {0, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}, {-1, -1}}
	f.updateFromFileFloor(5, 2)
	if !reflect.DeepEqual(f.content, result) {
		t.Error("faux")
	}
}
