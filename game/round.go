package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	
)

// Dans le cadre de l'extention Terre Ronde
// Permet de téléporter le personnage au bord opposé du terrain si on en sort
func (g *Game) RoundEarth() {
	//Bord haut
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if g.character.Y == 0 {
			if !*g.character.Moving {
				*g.character.Moving = true
				
				//enregsitre les coordonnées où téléporter le personnage
				g.character.Tp_roundY = g.floor.QuadtreeContent.Height() - 1
			}
		}
	}
	//Bord Bas
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if g.character.Y == g.floor.QuadtreeContent.Height()-1 {
			if !*g.character.Moving {
				*g.character.Moving = true
				
				//enregsitre les coordonnées où téléporter le personnage
				g.character.Tp_roundY = 0
			}
		}
	}
	//Bord Droit
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if g.character.X == g.floor.QuadtreeContent.Width()-1 {
			if !*g.character.Moving {
				*g.character.Moving = true
				
				//enregsitre les coordonnées où téléporter le personnage
				g.character.Tp_roundX = 0
			}
		}
	}
	//Bord Gauche
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if g.character.X == 0 {
			if !*g.character.Moving {
				*g.character.Moving = true
				
				//enregsitre les coordonnées où téléporter le personnage
				g.character.Tp_roundX = g.floor.QuadtreeContent.Width() - 1
			}
		}
	}
}
