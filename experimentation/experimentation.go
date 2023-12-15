package experimentation

import (
	"fmt"
	"log"
	"os"
	abr "projet/abr"
	"projet/cle"
	"projet/md5"
	"reflect"
	"slices"
	"strings"
)

// ParseBooksABR renvoi un arbre de recherche et une liste de tous les mots uniques parsés depuis les fichiers du répértoire ./Shakespeare
func ParseBooksABR() (abr.ArbreRecherche, []string) {
	fileEntries, err := os.ReadDir("./Shakespeare")
	if err != nil {
		log.Fatal(err)
	}

	var md5Tree abr.ArbreRecherche
	var words []string

	for _, file := range fileEntries {
		fileContent, err := os.ReadFile("Shakespeare" + string(os.PathSeparator) + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		lines := strings.Split(string(fileContent), "\n")

		for _, line := range lines {
			currWord := strings.TrimSpace(line)
			md5Hash := md5.Md5New([]byte(currWord))
			c := cle.BytesToCle(md5Hash)
			if !md5Tree.Contient(c) {
				md5Tree.Ajout(c)
				words = append(words, currWord)
			}
		}
	}

	return md5Tree, words
}

// ParseBooks renvoi une liste des cles parsés de chaque fichier dans le répértoire ./Shakespeare
func ParseBooks() [][]cle.Cle {
	fileEntries, err := os.ReadDir("./Shakespeare")
	if err != nil {
		log.Fatal(err)
	}

	var cles [][]cle.Cle

	for _, file := range fileEntries {
		fileContent, err := os.ReadFile("Shakespeare" + string(os.PathSeparator) + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		words := strings.Split(string(fileContent), "\n")

		var currCles []cle.Cle
		for _, word := range words {
			md5Hash := md5.Md5New([]byte(word))
			c := cle.BytesToCle(md5Hash)
			if !slices.Contains(currCles, c) {
				currCles = append(currCles, c)
			}
		}

		cles = append(cles, currCles)
	}

	return cles
}

// CollisionMd5 renvoie la liste de tous les mots dont le hash md5 sont en collision
func CollisionMd5(words []string) [][]string {
	var motsEnCollision [][]string

	for i := 0; i < len(words); i++ {
		currCollisions := []string{words[i]}

		currHash := md5.Md5New([]byte(words[i]))
		fmt.Println(i+1, "/", len(words), currHash)
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			targetHash := md5.Md5New([]byte(words[j]))
			if reflect.DeepEqual(currHash, targetHash) {
				currCollisions = append(currCollisions, words[j])
			}
		}
		if len(currCollisions) > 1 {
			motsEnCollision = append(motsEnCollision, currCollisions)
		}
	}

	return motsEnCollision
}
