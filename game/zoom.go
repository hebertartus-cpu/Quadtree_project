package game

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

//dans le cadre de l'extention de Zoom
//met à jour le cadre
func modif_cadre() {
	configuration.Global.ScreenWidth = configuration.Global.NumTileX * configuration.Global.TileSize
	configuration.Global.ScreenHeight = configuration.Global.NumTileY * configuration.Global.TileSize
	configuration.Global.ScreenCenterTileX = configuration.Global.NumTileX / 2
	configuration.Global.ScreenCenterTileY = configuration.Global.NumTileY / 2

}

//dans le cadre de l'extention de Zoom
//modifie la taille d'un tableau pour lui enlever une ligne et une colonne
func rezoom(tab *[][]int) {
	
	//suprime la denirère colonne de tab 
	for i:=0; i < len(*tab); i++ {

		(*tab)[i] = (*tab)[i][:len((*tab)[i])-1]
	}

	//suprime la dernière ligne de tab
	*tab = (*tab)[:len(*tab)-1]
}

//dans le cadre de l'extention de Zoom
//modifie la taille d'un tableau pour lui ajouter une ligne et une colonne
func dezoom(tab *[][]int) {
	
	//ajoute une collone à la fin
	for i:=0; i < len(*tab); i++ {

		(*tab)[i] = append((*tab)[i], -1)
	}
	
	//ajoute une ligne à la fin
	var ligne []int
	for i:=0; i < len((*tab)[0]); i++ {

		ligne = append(ligne, -1)
	}
	*tab = append((*tab), ligne)
}
