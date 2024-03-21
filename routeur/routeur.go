package routeur

import (
	"fmt"
	"log"
	"net/http"
	"royaleapi/controller"
)

func Initserv() {

	css := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	http.HandleFunc("/accueil", controller.AccueilPage)
	http.HandleFunc("/player", controller.TagTreatment)
	http.HandleFunc("/clan", controller.ClanTreatment)
	http.HandleFunc("/search_clan", controller.SearchClan)
	http.HandleFunc("/loading", controller.LoadingPage)
	http.HandleFunc("/chests", controller.ChestsPage)
	http.HandleFunc("/games", controller.GamesPage)
	http.HandleFunc("/cards", controller.GetCard)
	http.HandleFunc("/add-card", controller.SendCard)
	http.HandleFunc("/remove", controller.RemoveCardHandler)
	http.HandleFunc("/favoris", controller.FavorisCard)
	http.HandleFunc("/", controller.ErrorPage)


	// Démarrage du serveur
	log.Println("[✅] Serveur lancé !")
	fmt.Println("[🌐] http://localhost:8080/accueil")
	http.ListenAndServe(":8080", nil)
	log.Fatal()
}
