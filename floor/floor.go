package floor

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"

// Floor représente les données du terrain. Pour le moment
// aucun champs n'est exporté.
//
//   - content : partie du terrain qui doit être affichée à l'écran
//   - fullContent : totalité du terrain (utilisé seulement avec le type
//     d'affichage du terrain "fromFileFloor")
//   - quadTreeContent : totalité du terrain sous forme de quadtree (utilisé
//     avec le type d'affichage du terrain "quadtreeFloor")
//
//	 - Content et QuadtreeContent sont des pointeurs vers les chanps de même nom (dans init)
//	   Peremt d'accéder à ces champs à l'exterieur du package floor
//	 - Fram_count : un entier qui compte le nombre de frame depuis le lancemant du jeu
//
//dans le cadre de l'extention Bloquage caméra au bord
//	 - DecalX et DecalY permetent de décaler la caméra corectement pour qu'elle ne déborde pas du terrain
//
//dans le cadre de l'extention Animation du décor
//	 - Anim_info: tableau de tableau à deux champs qui permet l'animation du décor
//	   Les sous tableaux sont constitués en premier champ du nombre de frame d'animation d'un type de sol (diviseur de 60, si possible),
//	   en deuxième champ la frame actuel de l'animation
//	   et en troisième champ le nombre de frame de jeu entre chaque frame d'animation
//	   Le type de sol corespondant est la position du sous tableau dans Anim_info ([[4, 2], [3, 1]] -> [3, 1] est le sol 1)
//	   En cas de rajout de cases, il faut ajouter un sous tableau dans Anim_info, même s'il n'y a pas d'animation

type Floor struct {
	content         [][]int
	fullContent     [][]int
	quadtreeContent quadtree.Quadtree
	
	
	Content 		*[][]int
	QuadtreeContent *quadtree.Quadtree
	Frame_count		int

	DecalX			int
	DecalY			int

	Anim_info		[][]int
	
}

// types d'affichage du terrain disponibles
const (
	GridFloor int = iota
	FromFileFloor
	QuadTreeFloor
)

// GetHeight retourne la hauteur (en cases) du terrain
// à partir du tableau fullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetHeight() (height int) {
	return len(f.fullContent)
}

// GetWidth retourne la largeur (en cases) du terrain
// à partir du tableau fullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetWidth() (width int) {
	if len(f.fullContent) > 0 {
		width = len(f.fullContent[0])
	}
	return
}
