package main

import (
	"fmt"
	"os"
	"strings"

	"go-reloaded/Reloaded"
)

func args() bool {
	if len(os.Args) != 3 {
		fmt.Println("Il faut donner 2 fichiers")
		return false
	}
	return true
}

func main() {
	ok := args()
	if !ok {
		return
	}

	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Erreur dans la lecture du fichier :", err)
		return
	}

	texte := string(contenu)
	lignes := strings.Split(texte, "\n")
	for i, ligne := range lignes {
		lignes[i] = Reloaded.Words(ligne)
	}
	texte = strings.Join(lignes, "\n")

	err = os.WriteFile(os.Args[2], []byte(texte), 0644)
	if err != nil {
		fmt.Println("Erreur dans la création du fichier :", err)
		return
	}
}