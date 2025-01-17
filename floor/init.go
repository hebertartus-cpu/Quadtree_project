package floor

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.content = make([][]int, configuration.Global.NumTileY)
	//créé un pointeur vers content pour qu'il soit utilisable dans d'autres fichiers
	f.Content = &f.content
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	//frame_count permet de compter le nombre de frame écoulée depuis le lancement du jeu
	f.Frame_count = 0

	//dans le cadre de l'extention Animation du décor
	//initialise le tableau de tableau utiliser pour les animations (floor.Anim_info)
	//En cas de rajout de cases, il faut ajouter un sous tableau dans Anim_info, même s'il n'y a pas d'animation
	if configuration.Global.Animation == 1 {

		//Le premier champ correspond au nombre de frame de l'annimation
		//S'il n'y a pas d'animation, ce nombre est à 1
		//Le troisième champ  corespond au nombre de frame de jeu entre chaque frame d'animation
		f.Anim_info = [][]int{{1, 0, 1}, {1, 0, 1}, {1, 0, 1}, {1, 0, 1}, {6, 0, 30}, {4, 0, 30}, {4, 0, 30}}
	}

	switch configuration.Global.FloorKind {
	case FromFileFloor:
		f.fullContent = readFloorFromFile(configuration.Global.FloorFile)

	case QuadTreeFloor:

		//cas où il faut générer aléatoirement el terrain
		if len(configuration.Global.RandomFixe) == 2 {

			//enregistre le tableau généré pour éventuellement le sauvgarder dans un fichier
			temp := random_gene(configuration.Global.RandomFixe[0], configuration.Global.RandomFixe[1])

			//si l'extention de sauvegarde est acitve
			if len(configuration.Global.SaveName) != 0 {

				//sauvegarde le terrain généré aléatoirement
				savefile(temp)
			}

			//met le térrain aléatoire dans le quadtree pour qu'il soit utilisé
			f.quadtreeContent = quadtree.MakeFromArray(temp)
		} else {

			//cas initial, sans génération aléatoire et cas où 1 ou plus de 2 valeur
			//sont placé dans dans RandomFixe
			f.quadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
		}

		//créé un pointeur vers le quadtree pour qu'il soit utilisable dans d'autre fichier
		f.QuadtreeContent = &f.quadtreeContent
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
		for i := 0; i < len(ligne_str); i++ {

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

				if len(ligne_tab) != len(floorContent[0]) {

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
