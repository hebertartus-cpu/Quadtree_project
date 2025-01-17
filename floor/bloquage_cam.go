package floor


//Extention bloquage caméra au bord de la map
//met à jour DecalX et DecalY qui permetteront de décaler corectement la camméra si cela est nessssaire
func (f *Floor) Bloquage(x, y int) {

	//Gauche
	if x-len(f.content[0])/2 < 0 {
		f.DecalX = x - len(f.content[0])/2
	}


	//Droite
	//si paire ou impaire...
	if len(f.content[0])%2 == 0 {
		
		if x+len(f.content[0])/2-1 >= f.quadtreeContent.Width() {
			f.DecalX = x - f.quadtreeContent.Width() + 1 + len(f.content[0])/2 - 1
		}

	} else if len(f.content[0])%2 == 1 {

		if x+len(f.content[0])/2 >= f.quadtreeContent.Width() {
			f.DecalX = x - f.quadtreeContent.Width() + 1 + len(f.content[0])/2
		}
	}


	//Haut
	if y-len(f.content)/2 < 0 {
		f.DecalY = y - len(f.content)/2
	}


	//Bas
	//si paire ou impaire...
	if len(f.content)%2 == 0 {
		if y+len(f.content)/2-1 >= f.quadtreeContent.Height() {
			f.DecalY = y - f.quadtreeContent.Height() + 1 + len(f.content)/2 - 1
		}

	} else if len(f.content)%2 == 1 {
		if y+len(f.content)/2 >= f.quadtreeContent.Height() {
			f.DecalY = y - f.quadtreeContent.Height() + 1 + len(f.content) / 2
		}
	}

	//met DecalX à 0 si on est pas sur un bord de map
	//Droite - Gauche
	if len(f.content[0])%2 == 0 {
	
		if x-len(f.content[0])/2 >= 0 && x+len(f.content[0])/2-1 < f.quadtreeContent.Width() {

			f.DecalX = 0
		}
	} else {

		if x-len(f.content[0])/2 >= 0 && x+len(f.content[0])/2 < f.quadtreeContent.Width() {

			f.DecalX = 0
		}
	}

	//met DecalY à 0 si on est pas sur un bord de map
	//Haut - Bas
	if len(f.content)%2 == 0 {

		if y-len(f.content)/2 >= 0 && y+len(f.content)/2-1 < f.quadtreeContent.Height() {

			f.DecalY = 0
		}
	} else {

		if y-len(f.content)/2 >= 0 && y+len(f.content)/2 < f.quadtreeContent.Height() {

			f.DecalY = 0
		}
	}
}
