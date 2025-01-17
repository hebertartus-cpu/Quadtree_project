package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).
func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		configuration.Global.DebugMode = !configuration.Global.DebugMode

	}

	g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y))

	//Extention Terre Ronde
	if configuration.Global.RoundEarth == 1 {

		g.RoundEarth()
	}

	//Extention de portail de téléportation
	if configuration.Global.TPactive == 1 {

		//permet de placer un portail si "T" est précé
		if inpututil.IsKeyJustPressed(ebiten.KeyT) {

			//enregistre les co d'un portail
			g.character.Placementportails()
			//modifie la case du portail et sauvegardela valeur initiale
			g.sauv_sol_portail()
		}
		
		//permet de suprimer les portail placé
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {

			//réinitialise le type de sol au niveau des potatails
			g.reprend_val_sol()
			//suprime les valeur enregistré pour les portails.
			g.character.Suppr_portails()
		}
		
		g.character.Teleport()
	}


	//Exetention de zomm
	if configuration.Global.Zoom == 1 {

		//zoom en appuyant sur "E"
		if inpututil.IsKeyJustPressed(ebiten.KeyE) {

			//ne peut pas zoomé plus que 3 case de large ou de long
			if len(*g.floor.Content) > 3 && len((*g.floor.Content)[0]) > 3 {

				//retire une ligne et une colonne à floor.content pour réduire la taille de la caméra
				rezoom(g.floor.Content)

				//met à jour la taille du cadre
				configuration.Global.NumTileX --
				configuration.Global.NumTileY --
				modif_cadre()
			}
		}

		//dezoom en appuyant sur "F"
		if inpututil.IsKeyJustPressed(ebiten.KeyF) {
			
			//ajoute une ligne et une colonne à floor.content pour augmenter la taille de la caméra
			dezoom(g.floor.Content)
			
			//met à jour la taille du cadre
			configuration.Global.NumTileX ++
			configuration.Global.NumTileY ++
			modif_cadre()
		}
	}

	//Extention Bloquage caméra au bord
	//s'active aussi pour l'extention Round Earth
	if configuration.Global.CamBloque == 1 || configuration.Global.RoundEarth == 1{

		//met DecalX et DecalY de floor au bon chiffres pour decaler corectement la cam
		g.floor.Bloquage(g.character.X, g.character.Y)
	}

	g.camera.Update(g.character.X - g.floor.DecalX, g.character.Y - g.floor.DecalY)
	g.floor.Update(g.camera.X, g.camera.Y)
	
	return nil
}




