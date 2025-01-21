package game

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

//ce fichier gère certaines des fonctions associées à l'extention Monde Infinie

//Cette méthode détecte si la caméra sort du cadre déjà généré et donc s'il faut générer un nouveau terrain aléatoire,
//et le cas échéant, s'il faut agrandir le quadtree.
//Renvoie un entier qui dit dans quelle direction généré.
//Entre 1 et 8: le quadtree doit être agrandit
//entre (1 et 4) et (10 et 40) et (91 et 94): il faut générer une seul zone de terrain
//entre (5 et 8) et (50 et 80): il faut générer deux zones de terrain
func (g *Game) detect_generation() (sortie int) {

	//gère le décalage à droite et en bas si la caméra est de taille paire
	decale_camX := ((configuration.Global.NumTileX %2) +1) %2
	decale_camY := ((configuration.Global.NumTileY %2) +1) %2

	//Par la droite, si on sort du quadtree initilaisé
	if g.character.X + configuration.Global.NumTileX /2 - decale_camX >= g.floor.QuadtreeContent.Width() + g.floor.QuadtreeContent.MinX {

		//si la caméra se retrouve dans 2 zones différentes, il faut alors toutes les 2 les générer
		if (g.character.Y + configuration.Global.NumTileY /2 - decale_camY) / configuration.Global.InfinitRandom[1] !=
		   (g.character.Y - configuration.Global.NumTileY /2) / configuration.Global.InfinitRandom[1] {

			return 5
		} else {

			return 1
		}

	//Par le bas, si on sort du quadtree initilaisé
	} else if g.character.Y + configuration.Global.NumTileY /2 - decale_camY >= g.floor.QuadtreeContent.Height() + g.floor.QuadtreeContent.MinY {

		//si la caméra se retrouve dans 2 zones différentes, il faut alors toutes les 2 les générer
		if (g.character.X + configuration.Global.NumTileX /2 - decale_camX) / configuration.Global.InfinitRandom[0] !=
		   (g.character.X - configuration.Global.NumTileX /2) / configuration.Global.InfinitRandom[0] {

			return 6
		} else {

			return 2
		}
	
	//Par la gauche, si on sort du quadtree initilaisé
	} else if g.character.X - configuration.Global.NumTileX /2 < g.floor.QuadtreeContent.MinX {

		//si la caméra se retrouve dans 2 zones différentes, il faut alors toutes les 2 les générer
		if (g.character.Y + configuration.Global.NumTileY /2 - decale_camY) / configuration.Global.InfinitRandom[1] !=
		   (g.character.Y - configuration.Global.NumTileY /2) / configuration.Global.InfinitRandom[1] {

			return 7
		} else {

			return 3
		}

	//Par le haut, si on sort du quadtree initilaisé
	} else if g.character.Y - configuration.Global.NumTileY /2 < g.floor.QuadtreeContent.MinY {

		//si la caméra se retrouve dans 2 zones différentes, il faut alors toutes les 2 les générer
		if (g.character.X + configuration.Global.NumTileX /2 - decale_camX) / configuration.Global.InfinitRandom[0] !=
		   (g.character.X - configuration.Global.NumTileX /2) / configuration.Global.InfinitRandom[0] {

			return 8
		} else {

			return 4
		}
	
	//arrivé ici, on sait qu'on est dans le quadtree. 
	//Il faut savoir s'il faut malgrés tout générer un nouveau terrain aléatoire
	} else {

		tab := []bool{false, false, 
					  false, false}

		//Couin droit bas de la cam
		if g.floor.QuadtreeContent.ValeurQuad(g.character.X + configuration.Global.NumTileX /2 - decale_camX, g.character.Y + configuration.Global.NumTileY /2 - decale_camY) == -3 {
			
			tab[3] = true
		}
		//Couin droit haut de la cam
		if g.floor.QuadtreeContent.ValeurQuad(g.character.X + configuration.Global.NumTileX /2 - decale_camX, g.character.Y - configuration.Global.NumTileY /2) == -3 {
			
			tab[1] = true
		}
		//Couin gauche bas de la cam
		if g.floor.QuadtreeContent.ValeurQuad(g.character.X - configuration.Global.NumTileX /2, g.character.Y + configuration.Global.NumTileY /2 - decale_camY) == -3 {
			
			tab[2] = true
		}
		//Couin gauche haut de la cam
		if g.floor.QuadtreeContent.ValeurQuad(g.character.X - configuration.Global.NumTileX /2, g.character.Y - configuration.Global.NumTileY /2) == -3 {
			
			tab[0] = true
		}

		//droit bas - droit haut
		if tab[3] && tab[1] {

			//si la caméra se retrouve dans 2 zones différentes, il faut alors toutes les 2 les générer
			if (g.character.Y + configuration.Global.NumTileY /2 - decale_camY) / configuration.Global.InfinitRandom[1] !=
			   (g.character.Y - configuration.Global.NumTileY /2) / configuration.Global.InfinitRandom[1] {

				return 50
		    } else {

				return 10
			}
		//bas droit - bas gauche
		} else if tab[3] && tab[2] {

			//si la caméra se retrouve dans 2 zones différentes, il faut alors toutes les 2 les générer
			if (g.character.X + configuration.Global.NumTileX /2 - decale_camX) / configuration.Global.InfinitRandom[0] !=
		   	   (g.character.X - configuration.Global.NumTileX /2) / configuration.Global.InfinitRandom[0] {

				return 60
			} else {

				return 20
			}
		//gauche bas - gauche haut
		} else if tab[2] && tab[0] {

			//si la caméra se retrouve dans 2 zones différentes, il faut alors toutes les 2 les générer
			if (g.character.Y + configuration.Global.NumTileY /2 - decale_camY) / configuration.Global.InfinitRandom[1] !=
			   (g.character.Y - configuration.Global.NumTileY /2) / configuration.Global.InfinitRandom[1] {

		 		return 70
	 		} else {

		 		return 30
	 		}
		//haut droit - haut gauche
		} else if tab[1] && tab[0] {

			//si la caméra se retrouve dans 2 zones différentes, il faut alors toutes les 2 les générer
			if (g.character.X + configuration.Global.NumTileX /2 - decale_camX) / configuration.Global.InfinitRandom[0] !=
			   (g.character.X - configuration.Global.NumTileX /2) / configuration.Global.InfinitRandom[0] {

		 		return 80
	 		} else {

		 		return 40
	 		}
		
		//seleument haut gauche
		} else if tab[0] {

			return 93

		//seleument haut droit
		} else if tab[1] {

			return 94

		//seleument bas gauche
		} else if tab[2] {

			return 92

		//seleument bas droit
		} else if tab[3] {

			return 91		
		}
	}

	//s'il n'y a rien à générer
	return 0
}


func (g *Game) lance_generation(verif int) {

	//gère le décalage à gauche et en heut si la caméra est de taille paire
	decale_camX := ((configuration.Global.NumTileX %2) +1) %2
	decale_camY := ((configuration.Global.NumTileY %2) +1) %2


	//Cas où il faut agrandir le quadtree
	if verif < 9 {

		if verif >=5 {

			g.floor.QuadtreeContent.Agrandir(verif-4)
		} else {

			g.floor.QuadtreeContent.Agrandir(verif)
		}
	}

	//Dans tous les cas, génère un terrain aléatoire
	terrain := quadtree.MakeFromArray(floor.Random_gene(configuration.Global.InfinitRandom[0], configuration.Global.InfinitRandom[1]))

	//cas où l'on génère une seule zone
	//genère vers la droite ou la zone bas-droit
	if verif == 1 || verif == 10 || verif == 91 {

		g.floor.QuadtreeContent.Modif_zone(g.character.X + configuration.Global.NumTileX /2 - decale_camX, g.character.Y + configuration.Global.NumTileY /2 - decale_camY, terrain.Width(), terrain.Height(), terrain.Root())
	
	//genère vers le bas ou la zone bas-gauche
	} else if verif == 2 || verif == 20 || verif == 92 {

		g.floor.QuadtreeContent.Modif_zone(g.character.X - configuration.Global.NumTileX /2, g.character.Y + configuration.Global.NumTileY /2 - decale_camY, terrain.Width(), terrain.Height(), terrain.Root())
	
	//genère vers la gauche ou la zone haut-gauche
	} else if verif == 3 || verif == 30 || verif == 93 {

		g.floor.QuadtreeContent.Modif_zone(g.character.X - configuration.Global.NumTileX /2, g.character.Y - configuration.Global.NumTileY /2, terrain.Width(), terrain.Height(), terrain.Root())
	
	//genère vers le haut ou la zone haut-droit
	} else if verif == 4 || verif == 40 || verif == 94 {

		g.floor.QuadtreeContent.Modif_zone(g.character.X + configuration.Global.NumTileX /2 - decale_camX, g.character.Y - configuration.Global.NumTileY /2, terrain.Width(), terrain.Height(), terrain.Root())
	
	//cas où l'on génère 2 zones
	} else {

		//doit générer un deuxième terrain
		terrain2 := quadtree.MakeFromArray(floor.Random_gene(configuration.Global.InfinitRandom[0], configuration.Global.InfinitRandom[1]))

		//genère bas-droit et haut-droit
		if verif == 5 || verif == 50 {

			//bas-droit / terrain
			g.floor.QuadtreeContent.Modif_zone(g.character.X + configuration.Global.NumTileX /2 - decale_camX, g.character.Y + configuration.Global.NumTileY /2 - decale_camY, terrain.Width(), terrain.Height(), terrain.Root())

			//haut-droit / terrain2
			g.floor.QuadtreeContent.Modif_zone(g.character.X + configuration.Global.NumTileX /2 - decale_camX, g.character.Y - configuration.Global.NumTileY /2, terrain2.Width(), terrain2.Height(), terrain2.Root())
		
		//genère bas-droit et bas gauche
		} else if verif == 6 || verif == 60 {

			//bas-droit / terrain
			g.floor.QuadtreeContent.Modif_zone(g.character.X + configuration.Global.NumTileX /2 - decale_camX, g.character.Y + configuration.Global.NumTileY /2 - decale_camY, terrain.Width(), terrain.Height(), terrain.Root())

			//bas-gauche / terrain2
			g.floor.QuadtreeContent.Modif_zone(g.character.X - configuration.Global.NumTileX /2, g.character.Y + configuration.Global.NumTileY /2 - decale_camY, terrain2.Width(), terrain2.Height(), terrain2.Root())
		
		//genère haut-gauche et bas gauche
		} else if verif == 7 || verif == 70 {
			
			//haut-gauche / terrain
			g.floor.QuadtreeContent.Modif_zone(g.character.X - configuration.Global.NumTileX /2, g.character.Y - configuration.Global.NumTileY /2, terrain.Width(), terrain.Height(), terrain.Root())
			
			//bas-gauche /terrain2
			g.floor.QuadtreeContent.Modif_zone(g.character.X - configuration.Global.NumTileX /2, g.character.Y + configuration.Global.NumTileY /2 - decale_camY, terrain2.Width(), terrain2.Height(), terrain2.Root())
		
		//genère haut-gauche et haut-droit
		} else if verif == 8 || verif == 80 {

			//haut-gauche / terrain
			g.floor.QuadtreeContent.Modif_zone(g.character.X - configuration.Global.NumTileX /2, g.character.Y - configuration.Global.NumTileY /2, terrain.Width(), terrain.Height(), terrain.Root())

			//haut-droit / terrain2
			g.floor.QuadtreeContent.Modif_zone(g.character.X + configuration.Global.NumTileX /2 - decale_camX, g.character.Y - configuration.Global.NumTileY /2, terrain2.Width(), terrain2.Height(), terrain2.Root())
		}
	}
}



