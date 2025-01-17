package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY
	switch configuration.Global.FloorKind {
	case GridFloor:
		f.updateGridFloor(topLeftX, topLeftY)
	case FromFileFloor:
		f.updateFromFileFloor(topLeftX, topLeftY)
	case QuadTreeFloor:
		f.updateQuadtreeFloor(topLeftX, topLeftY)
	}

	//met à jour le nombre de frame écoulée depuis le lancement du jeu
	f.Frame_count ++

	//met à jour l'image de chaque animation
	f.update_anim()
}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(topLeftX, topLeftY int) {
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			absX := topLeftX
			if absX < 0 {
				absX = -absX
			}
			absY := topLeftY
			if absY < 0 {
				absY = -absY
			}
			f.content[y][x] = ((x + absX%2) + (y + absY%2)) % 2
		}
	}
}

// le sol est récupéré depuis un tableau, qui a été lu dans un fichier
//
// la version actuelle recopie fullContent dans content, ce qui n'est pas
// le comportement attendu dans le rendu du projet
func (f *Floor) updateFromFileFloor(topLeftX, topLeftY int) {
	
	var tableau [][]int	= readFloorFromFile(configuration.Global.FloorFile)
	
	for y:=0; y < len(f.content); y++ {

		for x:=0; x < len(f.content[y]); x++ {

			if y+ topLeftY >= 0 && x+ topLeftX >= 0 && (y + topLeftY) < len(tableau) {
				
				if (x + topLeftX) < len(tableau[y+topLeftY]) {
					f.content[y][x] = tableau[y+topLeftY][x+topLeftX]
				
					} else {
					f.content[y][x] = -1
				}

			} else {
				f.content[y][x] = -1
			}

		}
	}
}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(topLeftX, topLeftY int) {
	f.quadtreeContent.GetContent(topLeftX, topLeftY, f.content)
}
