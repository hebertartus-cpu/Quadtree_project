package game



//dans le cadre de l'extention de portail de téléportation
//sauvegarde le type de sol sous le portail 
//et replace cette case par une case de portail
func (g *Game) sauv_sol_portail() {

	//pour le premier portail
	if g.character.NbrT == 1 {

		g.character.Typesol = g.floor.QuadtreeContent.ValeurQuad(g.character.Valeur_posX, g.character.Valeur_posY) // va chercher une valeur
		g.floor.QuadtreeContent.Modif_Case(g.character.Valeur_posX, g.character.Valeur_posY, 5) // modifie une valeur aux coordonnées
		g.character.NbrT += 1

	//pour le deuxième portail
	} else if g.character.NbrT == 3 {
		
		g.character.Typesol2 = g.floor.QuadtreeContent.ValeurQuad(g.character.Valeur_posX2, g.character.Valeur_posY2)
		g.floor.QuadtreeContent.Modif_Case(g.character.Valeur_posX2, g.character.Valeur_posY2,6)
		g.character.NbrT += 1
	}

	
}

//dans le cadre de l'extention de portail de téléportation
//réinitialise les case de portail pour leur redonner leur apparence initiale
func (g *Game) reprend_val_sol() {

	//S'il y a un au moins un portail
	if g.character.NbrT >= 2 {
		
		//suprime le portail 1
		g.floor.QuadtreeContent.Modif_Case(g.character.Valeur_posX, g.character.Valeur_posY, g.character.Typesol)
		
		//s'il y a un deuxième portail
		if g.character.NbrT >= 4 {
			
			//suprime aussi le portail 2
			g.floor.QuadtreeContent.Modif_Case(g.character.Valeur_posX2, g.character.Valeur_posY2, g.character.Typesol2)
		}
	}
}
