package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"royaleapi/backend"
	"royaleapi/templates"
	"strings"
	"strconv"
	"os"
	
)

var token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6ImNhYmEzNjg2LWYzZmYtNDljNC1hMWZmLTRmMDBmNmM5OTkxZiIsImlhdCI6MTcxMDg2NzgxNiwic3ViIjoiZGV2ZWxvcGVyL2M1OGIxYTY4LWQzYzYtNWE0OS04ZDZjLWNiYTBhMDAwY2Q3YyIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI3Ny4yMDUuMjEuMTgzIl0sInR5cGUiOiJjbGllbnQifV19.WfzfrvxSYp4kJuHNVUHvmOODQBw3NkQysOhJSIzjvEwEMhKBZzejvh3PZH-OpVSejaAWa5JkJHacC3fzOHEupg"

func AccueilPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "accueil", nil)
}

func LoadingPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "loading", nil)
}

func ChestsPage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	tag := query.Get("tag")

	if tag[0] == '#' {
		tag = tag[1:]
	}

	fmt.Println(tag)

	url1 := "https://api.clashroyale.com/v1/players/%23"
	url := url1 + tag + "/upcomingchests"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var prochainCoffres backend.ListeCoffres
	if err := json.Unmarshal(body, &prochainCoffres); err != nil {
		log.Fatal(err)
	}

	templates.Temp.ExecuteTemplate(w, "chests", prochainCoffres)
}

func GamesPage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	tag := query.Get("tag")

	if tag[0] == '#' {
		tag = tag[1:]
	}

	fmt.Println(tag)

	url1 := "https://api.clashroyale.com/v1/players/%23"
	url := url1 + tag + "/battlelog"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var battles []backend.BattleLog
	err := json.Unmarshal(body, &battles)
	if err != nil {
		panic(err)
	}

	templates.Temp.ExecuteTemplate(w, "games", battles)
}

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "error", nil)
}

func TagTreatment(w http.ResponseWriter, r *http.Request) {
	tag := r.FormValue("tag")
	if tag[0] == '#' {
		tag = tag[1:]
	}

	fmt.Println(tag)
	url1 := "https://api.clashroyale.com/v1/players/%23"
	url := url1 + tag

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var playerData backend.PlayerData
	if err := json.Unmarshal(body, &playerData); err != nil {
		log.Fatal(err)
	}

	templates.Temp.ExecuteTemplate(w, "player", playerData)
}

func GetCard(w http.ResponseWriter, r *http.Request) {
	url2 := "https://api.clashroyale.com/v1/cards"

	req, err := http.NewRequest("GET", url2, nil)
	if err != nil {
		log.Fatal("erreur getCard")
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var card backend.AllCard

	

	if err := json.Unmarshal(body, &card); err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}

	fmt.Println(card)

	templates.Temp.ExecuteTemplate(w, "cards", card)
}


func SendCard(w http.ResponseWriter, r *http.Request) {
    fmt.Println("sendCard")
    cardName := r.FormValue("cardName")
    cardImage := r.FormValue("cardImage")

    // Vérifier si la carte existe déjà en favoris
    if backend.CardExistsInJSONFile(cardName, "./assets/favoris.json") {
		http.Redirect(w, r, "/cards", http.StatusSeeOther)
        return
    }

    err := backend.AddCardToJSONFile(cardName, cardImage, "./assets/favoris.json")
    if err != nil {
        http.Error(w, fmt.Sprintf("Erreur lors de l'ajout de la carte en favoris : %v", err), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/cards", http.StatusSeeOther)
}

func RemoveCardHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
        return
    }

    cardName := r.FormValue("cardName")
    if cardName == "" {
        http.Error(w, "Nom de carte manquant", http.StatusBadRequest)
        return
    }

    err := backend.RemoveCardFromJSONFile(cardName, "./assets/favoris.json")
    if err != nil {
        http.Error(w, fmt.Sprintf("Erreur lors de la suppression de la carte : %v", err), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/favoris", http.StatusSeeOther)
}





func ClanTreatment(w http.ResponseWriter, r *http.Request) {
	tag := r.FormValue("tag")
	if tag[0] == '#' {
		tag = tag[1:]
	}

	url1 := "https://api.clashroyale.com/v1/clans/%23"
	url := url1 + tag

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var clan backend.ClanData

	if err := json.Unmarshal([]byte(body), &clan); err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}

	templates.Temp.ExecuteTemplate(w, "clan", clan)
}

func FavorisCard(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("./assets/favoris.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	var cards []backend.CardData
	err = json.Unmarshal(data, &cards)
	if err != nil {
		log.Fatal(err)
	}
	templates.Temp.ExecuteTemplate(w, "favoris", cards)
}


func SearchClan(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        log.Fatal("Erreur lors du parsing du formulaire:", err)
        return
    }

    nom := r.FormValue("nom")
    minMembersStr := r.FormValue("min-members")
    maxMembersStr := r.FormValue("max-members")
    minTrophiesStr := r.FormValue("min-trophies")

    nom = strings.Replace(nom, " ", "%20", -1)
    url := fmt.Sprintf("https://api.clashroyale.com/v1/clans?name=%s", nom)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal("Erreur lors de la création de la requête:", err)
        return
    }

    req.Header.Set("Authorization", "Bearer " + token) 

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Erreur lors de l'exécution de la requête:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("Erreur lors de la lecture de la réponse:", err)
        return
    }

    var clans backend.Clans
    if err := json.Unmarshal(body, &clans); err != nil {
        fmt.Println("Erreur lors du décodage JSON:", err)
        return
    }

    minMembers, err := strconv.Atoi(minMembersStr)
    if err != nil {
        minMembers = 0 
    }
    maxMembers, err := strconv.Atoi(maxMembersStr)
    if err != nil {
        maxMembers = 50 
    }
    minTrophies, err := strconv.Atoi(minTrophiesStr)
    if err != nil {
        minTrophies = 0 
    }

    var filteredClans []backend.Clan
    for _, clan := range clans.Clans {
        if clan.Members >= minMembers && clan.Members <= maxMembers && clan.ClanWarTrophies >= minTrophies {
            filteredClans = append(filteredClans, clan)
        }
    }
	fmt.Println(filteredClans)

	
	var currentPage int
page := r.FormValue("page")
if page == "" {
    currentPage = 1
} else {
    pageInt, err := strconv.Atoi(page)
    if err != nil {
        currentPage = 1
    }
    currentPage = pageInt
}

fmt.Println(len(filteredClans))

startIndex := (currentPage - 1) * 20
endIndex := currentPage * 20
if endIndex > len(filteredClans) {
    endIndex = len(filteredClans)
}

currentClans := filteredClans[startIndex:endIndex]

fmt.Println(page)
fmt.Println(currentClans)

var prevPage int
if currentPage > 1 {
    prevPage = currentPage - 1
} else {
    prevPage = currentPage
}

var nextPage int
if endIndex < len(filteredClans) {
    nextPage = currentPage + 1
} else {
    nextPage = currentPage
}

templates.Temp.ExecuteTemplate(w, "searchclan", backend.TemplateData{Clans: currentClans, Nom: nom, MinMembers: minMembers, MaxMembers: maxMembers, MinTrophies: minTrophies, PagePrev: prevPage, PageNext: nextPage})
}