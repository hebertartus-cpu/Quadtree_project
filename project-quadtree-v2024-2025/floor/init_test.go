package floor

import (
    "testing"
    "reflect"
)


func TestReadFile1 (t *testing.T) {

    if !reflect.DeepEqual(readFloorFromFile("../floor-files/exemple"), [][]int{
        []int{1,1,3,4},
        []int{1,1,4,3},
        []int{0,0,2,2},
        []int{0,0,2,2}}) {
        
        t.Error("faux")
    }

}

func TestReadFile2(t *testing.T) {

	if !reflect.DeepEqual(readFloorFromFile("../floor-files/beaupasbeau"), [][]int{
		[]int{0, 0, 1, 0, 2, 2, 2, 2},
		[]int{0, 0, 1, 0, 2, 2, 2, 2},
		[]int{0, 0, 1, 0, 2, 2, 2, 2},
		[]int{0, 0, 1, 1, 0, 0, 0, 0},
		[]int{0, 0, 0, 1, 1, 1, 0, 0},
		[]int{2, 2, 0, 0, 0, 1, 0, 0},
		[]int{2, 2, 2, 0, 0, 1, 0, 0},
		[]int{2, 2, 2, 0, 0, 1, 0, 0}}) {

		t.Error("faux")
	}
}

func TestReadFile3(t *testing.T) {

	if !reflect.DeepEqual(readFloorFromFile("../floor-files/validation"), [][]int{
		[]int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		[]int{1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 2, 2, 1},
		[]int{0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0},
		[]int{1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1},
		[]int{0, 2, 0, 2, 0, 2, 2, 1, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0},
		[]int{1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 0, 2, 0, 1, 2, 2, 2, 1},
		[]int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}}) {

		t.Error("faux")

	}
}

func TestReadFileVide (t *testing.T) {

    var tab [][]int = readFloorFromFile("../floor-files/vide")

    if len(tab) !=0 {
        
        t.Error("faux")
    }

}

func TestReadFileTroue (t *testing.T) {

    if !reflect.DeepEqual(readFloorFromFile("../floor-files/exemple_trou"), [][]int{
        []int{1,1,3,4},
        []int{1,1,4,3},
        []int{0,0,2,2},
        []int{0,0,2,2}}) {
        
        t.Error("faux")
    }
}



func TestReadFileVideTroue (t *testing.T) {

    var tab [][]int = readFloorFromFile("../floor-files/vide_trou")

    if len(tab) != 0 {

        t.Error("faux")
    }
}
