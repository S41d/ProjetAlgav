package experimentation

import (
	crypto "crypto/md5"
	"fmt"
	"log"
	"os"
	abr "projet/abr"
	"projet/cle"
	"projet/md5"
	"reflect"
	"strings"
)

func ParseBooks() (abr.ArbreRecherche, []string) {
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

func CollisionMd5(tree abr.ArbreRecherche, words []string) [][]string {
	var motsEnCollision [][]string

	for i := 0; i < len(words); i++ {
		currCollisions := []string{words[i]}

		currHash := crypto.Sum([]byte(words[i]))
		fmt.Println(i+1, "/", len(words), currHash)
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			targetHash := crypto.Sum([]byte(words[j]))
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
