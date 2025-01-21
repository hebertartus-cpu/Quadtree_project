package floor

import (
	"math/rand"
)


//Dans le cadre de l'extention Génértion terrain aléatoire
//génère un tableu de tableau de taille (x, y) remplie de cases aléatoires
func Random_gene(x, y int) [][]int {
	if x <= 0 || y <= 0 {
		return [][]int{}

	}
	L := make([][]int, y)
	for i := 0; i < y; i++ {
		L[i] = make([]int, x)
		for j := 0; j < x; j++ {
			L[i][j] = rand.Intn(5)
		}
	}
	return L

}
