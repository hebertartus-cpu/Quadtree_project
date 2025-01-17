package character


//dans le cadre de l'extention de portail de téléportation
//enregistre les coordonnées des portail à leur placement
func (c *Character) Placementportails() {
	
	//s'il y a déjà un portail
	if c.NbrT == 2 {
		//si le portail n'est pas au même endroit que le premier
		if c.Valeur_posX != c.X || c.Valeur_posY != c.Y{
			
			c.Valeur_posX2 = c.X
			c.Valeur_posY2 = c.Y
			c.NbrT += 1
			c.tp_instant = 1
		}
	
	//s'il n't a pas encore de portail
	} else if c.NbrT == 0 {

		c.Valeur_posX = c.X
		c.Valeur_posY = c.Y
		c.NbrT += 1
		c.tp_instant = 1
		
	}
}

//dans le cadre de l'extention de portail de téléportation
//permet de téléporter le personnage s'il passe sur un portail
func (c *Character) Teleport() {

	//si les co du joueur corespondent à celle du portail 1
	if c.X == c.Valeur_posX && c.Y == c.Valeur_posY {
		//s'il ne vient pas à l'instant de placer ou utiliser un portail
		if c.tp_instant == 0 {
			//si les deuxportail sont posés
			if c.NbrT == 4 {
				c.X = c.Valeur_posX2
				c.Y = c.Valeur_posY2
				c.tp_instant = 1
			}
		}
	//meme chose pour le portail 2
	} else if c.X == c.Valeur_posX2 && c.Y == c.Valeur_posY2 {
		if c.tp_instant == 0 {
			if c.NbrT == 4 {
				c.X = c.Valeur_posX
				c.Y = c.Valeur_posY
				c.tp_instant = 1
			}
		}
	}

}

//dans le cadre de l'extention de portail de téléportation
//permet de suprimer les info relatives aux portail et de les réinitialiser
func (c *Character) Suppr_portails() {
	c.NbrT = 0
	c.Valeur_posX = 0
	c.Valeur_posY = 0
	c.Valeur_posX2 = 0
	c.Valeur_posY2 = 0
	c.Typesol=0
	c.Typesol2=0
}
