package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
	"bufio"
	"log"
	"os"
	"strconv"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	switch configuration.Global.FloorKind {
	case FromFileFloor:
		f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
	case QuadTreeFloor:
		f.quadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
	}
}

// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func readFloorFromFile(fileName string) (floorContent [][]int) {
	// TODO

	//Ovreture du fichier
	var myFile *os.File
	var err error

	myFile, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	//cree le scan
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(myFile)

	//deffinition de variable
	var ligne_str string
	var ligne_tab []int
	var carac int
	floorContent = [][]int{}

	//lance le scan
	for scanner.Scan() {

		ligne_str = scanner.Text()
		ligne_tab = make([]int, len(ligne_str))

		//convertie chaque caratère de la ligne en int et les met dans ligne_tab
		for i:=0; i < len(ligne_str); i++ {

			carac, err = strconv.Atoi(string(ligne_str[i]))
			if err != nil {
				log.Fatal(err)
			}

			ligne_tab = append(ligne_tab, carac)
		}

		//suprime des carractères indésirables
		ligne_tab = ligne_tab[len(ligne_tab)/2:]

		//suprime les lignes vide
		if len(ligne_str) != 0 {
			//verifie si la ligne fait la même taille que la première
			if len(floorContent) > 1 {

				if len(ligne_tab) != len(floorContent[0]){

					log.Fatal()
				}
			}

			//ajoute la ligne dans le tableau
			floorContent = append(floorContent, ligne_tab)
		}
	}

	//ferme le fichier
	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	return floorContent
}
