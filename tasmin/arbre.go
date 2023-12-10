package tasmin

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"projet/cle"
	"strings"
	"syscall"
)

// Arbre est une structure de données représentant un nœud dans un arbre.
type Arbre struct {
	// Size indique le nombre total de nœuds dans le sous-arbre enraciné à ce nœud.
	Size int
	// Height représente la hauteur du sous-arbre enraciné à ce nœud.
	height int
	// Cle est la valeur de la clé stockée dans ce nœud.
	cle cle.Cle
	// Parent pointe vers le nœud parent dans l'arbre.
	parent *Arbre
	// LeftChild pointe vers le nœud enfant gauche dans l'arbre.
	leftChild *Arbre
	// RightChild pointe vers le nœud enfant droit dans l'arbre.
	rightChild *Arbre
}

var index = 0

// Ajout ajoute un élément à l'arbre.
func (t *Arbre) Ajout(c cle.Cle) {
	if t.Size == 0 {
		*t = NewArbre(c)
		return
	}
	// Initialisation du pointeur courant à la racine de l'arbre
	curr := t

	// Vérifier si l'arbre est plein (2^h - 1 nœuds), où h est la hauteur actuelle de l'arbre
	if curr.Size == pow(2, curr.height)-1 {
		// Si l'arbre est plein, descendre à gauche jusqu'au dernier niveau
		for curr.leftChild != nil {
			curr.Size += 1
			curr.height += 1
			curr = curr.leftChild
		}

		// Mettre à jour la taille et la hauteur du nœud actuel
		curr.Size += 1
		curr.height += 1

		// Ajouter un nouveau nœud à gauche avec la clé donnée
		curr.leftChild = &Arbre{cle: c, Size: 1, height: 1, parent: curr}
		curr = curr.leftChild
	} else {
		// Si l'arbre n'est pas plein, parcourir l'arbre pour trouver l'emplacement d'insertion
		for {
			curr.Size += 1

			// Si le nœud courant n'a pas d'enfant gauche, ajouter le nouveau nœud à gauche
			if curr.leftChild == nil {
				curr.height += 1
				curr.leftChild = &Arbre{cle: c, Size: 1, height: 1, parent: curr}

				// Mettre à jour la hauteur des parents jusqu'à la racine
				currHeight := curr
				for currHeight.parent != nil {
					if currHeight.parent.height == currHeight.height+1 {
						break
					}
					currHeight.parent.height += 1
					currHeight = currHeight.parent
				}

				// Aller au nouveau nœud ajouté
				curr = curr.leftChild
				break
			} else if curr.rightChild == nil {
				// Si le nœud courant n'a pas d'enfant droit, ajouter le nouveau nœud à droite
				curr.rightChild = &Arbre{cle: c, Size: 1, height: 1, parent: curr}
				curr = curr.rightChild
				break
			}

			// Choisir le chemin en fonction de la taille des sous-arbres
			if curr.leftChild.Size <= curr.rightChild.Size || curr.leftChild.Size != pow(2, curr.leftChild.height)-1 {
				curr = curr.leftChild
			} else {
				curr = curr.rightChild
			}
		} // for
	} // else

	// Rétablir la propriété de l'arbre binaire équilibré en effectuant des échanges
	for curr.parent != nil {
		if !curr.parent.cle.Inf(curr.cle) {
			// Échanger les clés si la clé du parent est supérieure à la clé actuelle
			temp := *curr
			curr.cle = curr.parent.cle
			curr.parent.cle = temp.cle
			curr = curr.parent
		} else {
			break
		}
	} // for
}

func ConstructionArbre(cles []cle.Cle) Arbre {
	if len(cles) == 1 {
		return NewArbre(cles[0])
	}
	hauteur := int(math.Floor(math.Log2(float64(len(cles))))) + 1
	numNodes := 0
	if len(cles) == pow(2, hauteur)-1 {
		numNodes = pow(2, hauteur-1)
	} else {
		numNodes = len(cles) - (pow(2, hauteur-1) - 1)
	}
	bound := len(cles) - numNodes
	nodesClesEnf := cles[bound:]
	nodesEnf := listCleToListArbre(nodesClesEnf)

	for i := hauteur - 1; i > 0; i-- {
		cles = cles[:len(cles)-numNodes]
		numNodes = pow(2, i-1)
		nodesCles := cles[len(cles)-numNodes:]
		nodes := listCleToListArbre(nodesCles)
		for j := 0; j < len(nodes); j++ {
			k := j * 2
			if k >= len(nodesEnf) {
				break
			}
			if k != len(nodesEnf)-1 {
				nodes[j].leftChild = &nodesEnf[k]
				nodes[j].rightChild = &nodesEnf[k+1]

				nodes[j].leftChild.parent = &nodes[j]
				nodes[j].rightChild.parent = &nodes[j]

				nodes[j].Size = 1 + nodes[j].leftChild.Size + nodes[j].rightChild.Size
				nodes[j].height = 1 + max(nodes[j].leftChild.height, nodes[j].rightChild.height)
			} else {
				nodes[j].leftChild = &nodesEnf[k]
				nodes[j].leftChild.parent = &nodes[j]

				nodes[j].Size = 1 + nodesEnf[k].Size
				nodes[j].height = 1 + nodesEnf[k].height
			}
			nodes[j].reOrganiser()
		}
		nodesEnf = nodes
		nodesClesEnf = nodesCles
	}
	return nodesEnf[0]
}

func (t *Arbre) SupprMin() cle.Cle {
	if t.Size == 1 {
		return t.cle
	}
	deleted := t.cle

	curr := t
	for curr.Size != 1 {
		curr.Size -= 1
		if curr.rightChild != nil {
			if curr.leftChild.height == curr.rightChild.height {
				curr = curr.rightChild
			} else {
				// hauteur de leftChild >= hauteur de rightChild
				curr = curr.leftChild
			}
		} else if curr.leftChild != nil {
			curr = curr.leftChild
		}
	}
	m := curr.cle
	if curr.parent.leftChild == curr {
		curr.parent.leftChild = nil
		for {
			curr.height -= 1
			p := curr.parent
			maxH := p.height
			if p.leftChild != nil && p.leftChild.height > maxH {
				// left est toujours plus grand que right
				maxH = p.leftChild.height
			}
			if p.height == maxH {
				break
			}
			curr = p
		}
	} else {
		curr.parent.rightChild = nil
	}
	for curr.parent != nil {
		curr = curr.parent
	}
	curr.cle = m
	*t = *curr
	t.reOrganiser()
	return deleted
}

func (t *Arbre) Union(o *Arbre) Arbre {
	cles := append(ClesOfArbre(t), ClesOfArbre(o)...)
	return ConstructionArbre(cles)
}

func ClesOfArbre(a *Arbre) []cle.Cle {
	cles := make([]cle.Cle, a.Size)

	var aux func(curr *Arbre, idx int)
	aux = func(curr *Arbre, idx int) {
		cles[idx] = curr.cle
		if curr.leftChild != nil {
			aux(curr.leftChild, (2*idx)+1)
		}
		if curr.rightChild != nil {
			aux(curr.rightChild, (2*idx)+2)
		}
	}
	aux(a, 0)
	return cles
}

func (t *Arbre) reOrganiser() {
	smallest := &Arbre{}
	if t.leftChild != nil && t.rightChild != nil {
		if t.leftChild.cle.Inf(t.rightChild.cle) {
			smallest = t.leftChild
		} else {
			smallest = t.rightChild
		}
	} else if t.leftChild != nil {
		smallest = t.leftChild
	} else {
		return
	}
	if smallest.cle.Inf(t.cle) {
		t.cle, smallest.cle = smallest.cle, t.cle
		smallest.reOrganiser()
	}
}

// AjoutIteratif ajoute de manière itérative une liste de clés à l'arbre.
func (t *Arbre) AjoutIteratif(cles []cle.Cle) {
	// Appeler la fonction pour gérer les signaux (à implémenter ailleurs dans votre code)
	setSigHandler()

	// Parcourir toutes les clés dans la liste
	for index := 0; index < len(cles); index++ {
		// Appeler la fonction Ajout pour ajouter la clé actuelle à l'arbre
		t.Ajout(cles[index])
	}
}

// _string génère une représentation en chaîne de caractères de l'arbre, avec un décalage pour l'indentation.
func (t *Arbre) _string(indent int) string {
	// Initialiser la chaîne avec l'en-tête de l'arbre
	str := "Tas{\n"
	indent += 1 // Augmenter l'indentation pour les sous-arbres

	// Ajouter la représentation en chaîne de la clé du nœud courant
	str += strings.Repeat("  ", indent) + "cle: " + t.cle.DecimalString() + ",\n"

	// Ajouter d'autres propriétés du nœud (commentées pour le moment)
	// str += strings.Repeat("  ", indent) + "size: " + fmt.Sprint(t.size) + ",\n"
	// str += strings.Repeat("  ", indent) + "height: " + fmt.Sprint(t.height) + ",\n"

	// Ajouter la représentation en chaîne du sous-arbre gauche
	if t.leftChild != nil {
		str += strings.Repeat("  ", indent) + "left: " + t.leftChild._string(indent) + ",\n"
	} else {
		str += strings.Repeat("  ", indent) + "left: nil,\n"
	}

	// Ajouter la représentation en chaîne du sous-arbre droit
	if t.rightChild != nil {
		str += strings.Repeat("  ", indent) + "right: " + t.rightChild._string(indent) + "\n"
	} else {
		str += strings.Repeat("  ", indent) + "right: nil\n"
	}

	// Ajouter la fermeture de l'arbre avec le bon décalage
	str += strings.Repeat("  ", indent-1) + "}"

	return str
}

// setSigHandler configure un gestionnaire de signaux pour intercepter le signal SIGQUIT.
func setSigHandler() {
	// Créer un canal (channel) pour recevoir les signaux
	sigChan := make(chan os.Signal, 1)

	// Notifier le canal lorsqu'un signal SIGQUIT est reçu
	signal.Notify(sigChan, syscall.SIGQUIT)

	// Lancer une goroutine pour traiter les signaux en arrière-plan
	go func() {
		for {
			// Attendre la réception d'un signal sur le canal
			s := <-sigChan

			// Vérifier si le signal reçu est SIGQUIT
			if s == syscall.SIGQUIT {
				// affiche la valeur de l'index
				fmt.Println(index)
			}
		}
	}()
}

// NewArbre crée un nouvel arbre avec la clé spécifiée, initialisant la hauteur à 1 et la taille à 1.
func NewArbre(c cle.Cle) Arbre {
	return Arbre{cle: c, height: 1, Size: 1}
}

// pow retourne x élevé à la puissance y.
func pow(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// String retourne une représentation en chaîne de caractères de l'arbre.
func (t *Arbre) String() string {
	return t._string(0) + "\n"
}

func listCleToListArbre(cles []cle.Cle) []Arbre {
	arbres := []Arbre{}
	for i := 0; i < len(cles); i++ {
		arbres = append(arbres, NewArbre(cles[i]))
	}
	return arbres
}
