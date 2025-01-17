package floor

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

//Dans le cadre de l'extention sauvegarde terrain aléatoire
//enregistre chaque valeur d'un tableau de tableau dans un nouveau fichier
func savefile(t [][]int) {

	var myFile *os.File
	var err error

	//créé un fichier de nom SaveName
	myFile, err = os.Create( "../floor-files/" + configuration.Global.SaveName)
	if err != nil {
		log.Fatal(err)
	}

	var ligne string

	//enregistre chaque ligne dans le fichier
	for i := 0; i < len(t); i++ {

		ligne = ""

		//créé la ligne à enregistrer
		for j := 0; j < len(t[i]); j++ {

			ligne = ligne + strconv.Itoa(t[i][j])
		}

		//enregistre la ligne
		_, err = fmt.Fprintln(myFile, ligne)
		if err != nil {
			log.Fatal(err)
		}
	}

	//ferme le fichier
	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
