# project-quadtree

Dans le cadre de la SEA S1.02, Implémentation Besoins Client, année 2024-2025, 
rendu des élèves du groupe 2 MONNIN Damien et HEBERT Artus.

Travail rendu:

-Fonctions readFloorFromFile, updateFromFileFloor, MakeFromArray et GetContent complétées.
Elève(s) en charge: Artus et Damien.

-Implémentation de divers extentions détaillées ci-dessous.
Pour que ces extentions fonctionnent, il faut utiliser les quadtree et avoir une caméra mouvante. Les champs CameraMode et FloorKind doivent donc avoir respectivement les valeurs 1 et 2. L'utiliation de certains autres champs est expliqué ci-dessous. Si ce n'est pas le cas, leur utilisation est expliquer dans le sujet. Si ce n'est pas le cas, merci de laisser les champs dans leur état initial.

n°0: Coordonnées initiales
Dans le fichier json, remplir le champ CoInit avec un tableau de 2 entier corespondant au tuple (x, y) des coordonnées qui seront atribués au joueur comme point de départ. Ces entiers doivent être strictements positifs et infèrieurs à la taille du terrain utilisé. Si ces condition ne sont pas respecté, ou que le tableau est vide ou ne contient pas exactement 2 entiers, le joueur apparaitra aux coordonnées (0, 0).
Elève(s) en charge: Damien

n°1: Génération aléatoire de terrain.
Dans le fichier json, remplir le champ RandomFixe avec un tableau de 2 entiers corespondant à la largeur et à la hauteur du terrain à générer. Si le tableau est vide ou ne contient pas exactement 2 entiers strictements positifs, l'extention ne sera pas active.
Elève(s) en charge: Artus

n°2: Enregistrement d’un terrain
Dans le fichier json, remplir le champ SaveName avec une chaine de caractère qui sera le nom à donner au fichier qui sauvegardera le terrain. Le fichier sera enregistré dans le dossier floor-files. Si le nom rempli est identique à celui d'un fichier déjà présent dans le dossier floor-files, l'ancien fichier sera détruit. Si la chaine de caractère est vide, l'extention ne sera pas active.
Elève(s) en charge: Artus

n°3: Animation des décors
Dans le fichier json, remplir le champ Animation avec un 1. Les annimations enregistrées seront visible en jeu. Si le champ est rempli avec un 0, l'extention ne sera pas active.
Elève(s) en charge: Damien

n°4: Téléporteurs
Dans le fichier json, remplir le champ TPactive avec un 1. En jeu, en appuyant sur T, un portail sera placé sous le joueur. Si un deuxième portail est placé, le joueur pourra se téléporter de l'un à l'autre en marchant dessus. Il faut qu'il y ai 2 portails pour que le joueur puisse être téléporté. Après une téléportation ou après avoir placé un portail, le joueur doit quitter la case du téléporteur pour pouvoir à nouveau se téléporter. En jeu, en appuyant sur R, le ou les portails seront suprimés. Il ne peut pas y avoir plus de 2 portails en même temps. Si le champ est rempli avec un 0, l'extention ne sera pas active.
Elève(s) en charge: Artus

n°5: Interdiction de marcher sur l’eau
Dans le fichier json, remplir le champ ForbiddenFloor avec un entier représentant le type de sol qui devient bloquant. Pour connaitre l'entier associé à un type de sol, se référer au code ou à l'image floor.png. Si le champ est rempli avec un -1, l'extention ne sera pas active.
Elève(s) en charge: Damien

n°6: Caméra bloquée aux bords du terrain
Dans le fichier json, remplir le champ CamBloque avec un 1. En jeu, quand le joeur s'approche des bords du terrain, la caméra se bloque avant de laiser aparraitre du vide. Ainsi, le joueur ne se trouve plus au centre de la caméra. Il en redevient le centre en s'éloignant du bord. Le joueur ne peut pas sortir du terrain. Atention: la caméra soit être plus petite que la taille du terrain. De même, avec l'extention Zoom, il ne faut pas dézoomer plus que la taille du terrain, au risque de faire buguer cette extention. Si le champ est rempli avec un 0, l'extention ne sera pas active.
Elève(s) en charge: Artus

n°11: La terre est ronde
Dans le fichier json, remplir le champ RoundEarth avec un 1. Active l'extention Caméra bloquée aux bords du terrain. Si le joueur se trouve sur un bord et avance vers le vide, il sera téléporté au bord opposé. Puisque se basant sur l'extention Caméra bloquée aux bords du terrain, les recommandations quand à la taille de la caméra et du dézoom s'appliquent également pour cette extention. Si le champ est rempli avec un 0, l'extention ne sera pas active.
Elève(s) en charge: Artus et Damien

n°12: Zoom
Dans le fichier json, remplir le champ Zomm avec un 1. En jeu, en appuyant sur E, un zoom d'1 case s'effectuera: le terrain sera réduit d'une case de large et de haut. A l'invers, en appuyant sur F, un dézoom d'1 cas s'effectuera: le terrain sera augmenter d'une case de large et de haut. Le terrain ne peut pas être zoomé si un de ses dimentions attient la valeur de 3. Le terrain minimal est donc un carré de 3 par 3. Si le champ est rempli avec un 0, l'extention ne sera pas active.
Elève(s) en charge: Damien

n°13: Génération de terrain à la volée au fur et à mesure de l’exploration
Dans le fichier json, remplir le champ InfinitRandom avec un tableau de 2 entiers corespondant à la largeur et à la hauteur du terrain initiale. Si, lors d'un déplacement, la caméra sort de la zone déjà généré, un nouveau terrain de la même taille que le terrain initiale sera généré dans la direction du déplacement. cette génération est infinie et se poursuit dans les 4 directions, en fonction des déplacements du joueur. Le terrain ainsi généré ne peut pas être sauvegarder à l'aide de l'extention Enregistrement d’un terrain. Au lancement du jeu, le joueur sera placer au centre du terrai initial, sans tenir compte des éventuelles valeurs renseignées dans le champ CoInit. Par ailleurs, la taille de la caméra sera fixé pour être égale à celle du terrain initial, sans tenir compte des valeurs renseignées dans les champs NumTileX et NumTileY. Cette extention ne peut être active que si les extentions Caméra bloquée aux bords du terrain, La terre est ronde et Zoom sont désactivées. Si le tableau est vide ou ne contient pas exactement 2 entiers strictements positifs, l'extention ne sera pas active.
Elève(s) en charge: Damien


Ci-après, pour chaque extention, les fichiers où trouver le code relatif à chaque extention.

n°0:
Game: init.go

n°1:
Floor: init.go, random.go

n°2:
Floor: init.go, save.go

n°3:
Floor: init.go, update.go, animation.go, draw.go

n°4:
Game: portail_teleport.go, update.go
Character: init.go, character.go

n°5:
Floor: blocking.go, 

n°6:
Floor: bloquage_cam.go
Game: update.go

n°11:
Game: round.go, update.go
Character: update.go, init.go

n°12:
Game: zoom.go, update.go

n°13:
Game: update.go, infinit.go, init.go
Configuration: configuration.go
Floor: init.go, draw.go
Quadtree: quadtree.go, get.go

Toutes les extentions ajoutent au moins une ligne dans le fichier cmd.json et dans la structure Configuration du fichier Configuration.configuration.go.
Certaines des extentions utilisent des fonctions ou méthodes déjà codées.
Par ailleurs, le fichier quadtree.modifquad.go est utilisé dans plusieurs extention, et n'est pas attibuer à une seul d'entre elle. Son utilisation n'est donc pas précisée ci-dessus.