package floor

//gère les focntions et méthodes qui sont nécessaire ppour l'extention
//Animation du decor

//met à jour l'image de chaque animation
func (f *Floor) update_anim() {

	for i:=0; i < len(f.Anim_info); i++ {

		//si le bon nombre de frames sont passées
		if f.Frame_count % f.Anim_info[i][2] == 0 {

			//fait pasé l'animation à l'image suivante
			f.Anim_info[i][1] = (f.Anim_info[i][1] + 1) % f.Anim_info[i][0]
		}
	}

}