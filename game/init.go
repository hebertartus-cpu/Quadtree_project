package game

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisation car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	g.floor.Init()
	g.character.Init(g.floor.GetWidth(), g.floor.GetHeight())

	//permet de faire apparaitre le joueur à des coordonnées variables
	if len(configuration.Global.CoInit) == 2 {
		// évite de faire spawer le joueur à des valeurs négatives dans CoInit, fait spawn en 0,0 si condition non respectée
		if configuration.Global.CoInit[0] > -1 && configuration.Global.CoInit[1] > -1 {
			//Empêche le joueur de spawner trop loin sur la longueur ou la hauteur de la map, fait spawn en 0,0 sinonssssssss
			if configuration.Global.CoInit[0] < g.floor.QuadtreeContent.Width() && configuration.Global.CoInit[1] < g.floor.QuadtreeContent.Height() {
				g.character.X = configuration.Global.CoInit[0]
				g.character.Y = configuration.Global.CoInit[1]
			}

		}

	}

	//Extention sol interdit
	if configuration.Global.ForbiddenFloor >= 0 {

		//si la casse d'apparition du joueur est une case interdite
		if g.floor.QuadtreeContent.ValeurQuad(g.character.X, g.character.Y) == configuration.Global.ForbiddenFloor {

			//change le type de sol en "0" de la case de départ si le sol intedite n'est pas 0
			if configuration.Global.ForbiddenFloor != 0 {

				g.floor.QuadtreeContent.Modif_Case(g.character.X, g.character.Y, 0)

				//change le type de sol en "1" de la case de départ sinon
			} else {

				g.floor.QuadtreeContent.Modif_Case(g.character.X, g.character.Y, 1)
			}
		}
	}

	g.camera.Init(g.character.X, g.character.Y)
}
