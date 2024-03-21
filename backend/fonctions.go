package backend

import "fmt"
import "encoding/json"
import "io/ioutil"


func DisplayPlayerData(data PlayerData) {
	fmt.Println("Tag: ", data.Tag)
	fmt.Println("Name: ", data.Name)
	fmt.Println("Experience Level: ", data.ExpLevel)
	fmt.Println("Trophies: ", data.Trophies)
	fmt.Println("Best Trophies: ", data.BestTrophies)
	fmt.Println("Wins: ", data.Wins)

	fmt.Println("Arena ID: ", data.Arena.ID)
	fmt.Println("Arena Name: ", data.Arena.Name)

	fmt.Println("Clan Tag: ", data.Clan.Tag)
	fmt.Println("Clan Name: ", data.Clan.Name)
	fmt.Println(data.CurrentDeck)
	fmt.Println(data.Badges)
}


func AddCardToJSONFile(name, imageURL, jsonFilePath string) error {
    // Lire le contenu actuel du fichier JSON
    data, err := ioutil.ReadFile(jsonFilePath)
    if err != nil {
        return fmt.Errorf("erreur lors de la lecture du fichier JSON : %v", err)
    }

    // Convertir les données JSON en une slice de structs Card
    var cards []CardData
    err = json.Unmarshal(data, &cards)
    if err != nil {
        return fmt.Errorf("erreur lors de l'analyse du fichier JSON : %v", err)
    }

    // Vérifier si la carte existe déjà dans la liste des favoris
    for _, card := range cards {
        if card.Name == name {
            return fmt.Errorf("la carte '%s' existe déjà dans les favoris", name)
        }
    }

    // Ajouter la nouvelle carte à la slice
    newCard := CardData{
        Image: imageURL,
        Name:  name,
    }
    cards = append(cards, newCard)

    // Convertir la slice de structs Card en données JSON
    newData, err := json.MarshalIndent(cards, "", "    ")
    if err != nil {
        return fmt.Errorf("erreur lors de la conversion en JSON : %v", err)
    }

    // Écrire les nouvelles données JSON dans le fichier
    err = ioutil.WriteFile(jsonFilePath, newData, 0644)
    if err != nil {
        return fmt.Errorf("erreur lors de l'écriture dans le fichier JSON : %v", err)
    }

    return nil
}

func RemoveCardFromJSONFile(name, jsonFilePath string) error {
    // Lire le contenu actuel du fichier JSON
    data, err := ioutil.ReadFile(jsonFilePath)
    if err != nil {
        return fmt.Errorf("erreur lors de la lecture du fichier JSON : %v", err)
    }

    // Convertir les données JSON en une slice de structs Card
    var cards []CardData
    err = json.Unmarshal(data, &cards)
    if err != nil {
        return fmt.Errorf("erreur lors de l'analyse du fichier JSON : %v", err)
    }

    // Parcourir la slice pour trouver et supprimer la carte avec le nom spécifié
    var found bool
    for i, card := range cards {
        if card.Name == name {
            // Supprimer la carte trouvée de la slice
            cards = append(cards[:i], cards[i+1:]...)
            found = true
            break
        }
    }

    // Si la carte n'a pas été trouvée, retourner une erreur
    if !found {
        return fmt.Errorf("la carte '%s' n'a pas été trouvée dans le fichier JSON", name)
    }

    // Convertir la slice de structs Card en données JSON
    newData, err := json.MarshalIndent(cards, "", "    ")
    if err != nil {
        return fmt.Errorf("erreur lors de la conversion en JSON : %v", err)
    }

    // Écrire les nouvelles données JSON dans le fichier
    err = ioutil.WriteFile(jsonFilePath, newData, 0644)
    if err != nil {
        return fmt.Errorf("erreur lors de l'écriture dans le fichier JSON : %v", err)
    }

    return nil
}

func CardExistsInJSONFile(cardName, jsonFilePath string) bool {
    // Lire le contenu actuel du fichier JSON
    data, err := ioutil.ReadFile(jsonFilePath)
    if err != nil {
        fmt.Printf("erreur lors de la lecture du fichier JSON : %v", err)
        return false
    }

    // Convertir les données JSON en une slice de structs Card
    var cards []CardData
    err = json.Unmarshal(data, &cards)
    if err != nil {
        fmt.Printf("erreur lors de l'analyse du fichier JSON : %v", err)
        return false
    }

    // Vérifier si la carte existe déjà dans la liste des favoris
    for _, card := range cards {
        if card.Name == cardName {
            return true
        }
    }

    return false
}

