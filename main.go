package main

import (
	"flag"
	"fmt"
	"log"
	"main/src/engine"
	"net/http"
	_ "net/http/pprof"

)


func main() {
	var e engine.Engine
	Speed := flag.Bool("SPEED",false, "vitesse du personnage ultra rapide\n")  //pointeur Speed qui pointe sur SPEED de type bool donc vrai ou faux ( activé ou désactivé)défini par défaut a false car il n'est pas appelé en argument par défaut
	BigJump := flag.Bool("JUMP",false, "mega saut du personnage\n") // idem que pour Speed la variale est défini par défaut a false car il n'est pas appelé en argument par défaut
	help := flag.Bool("h",false, "Afficher de l'aide\n")          //idem que pour Speed défini par défaut a false car non appelé comme argument par defaut
	flag.Parse()                        // la fonction flag.Parse lit les argument entré en ligne de commande et les compare avec ceux defini ci dessus
	if *BigJump {                 // si le contenu du pointeur BigJump est a true donc  si l'utilisateur specifie l'argument -JUMP le flag bool est  a mit a true
		e.BigJump = true          // alors la dans la fonction engine bigJump qui est une variable booléen défini dans la structure est mit a true 
		fmt.Println("BigJump :", *BigJump) // on affiche dans le terminal "BigJump:" le contenu du pointeur  BigJump car on utilise *
	}
	if *help {
		flag.Usage()    // affiche toutes les Usage de main.go
		return            // fait un retour pour ne pas lancer le jeux
	} 
	if *Speed {                      // si le contenu du pointeur Speed est a true donc  si l'utilisateur specifie l'argument -SPEED le flag bool est  a mit a true
		e.SupSpeed = true
		fmt.Println("ultraspeed :", *Speed)
	}


	e.Init()
	e.Load()
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	e.Run()
	e.Unload()
	e.Close()

}