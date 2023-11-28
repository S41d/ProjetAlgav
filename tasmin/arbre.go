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

type Arbre struct {
	Size       int
	height     int
	cle        cle.Cle
	parent     *Arbre
	leftChild  *Arbre
	rightChild *Arbre
}

var index = 0

func (t *Arbre) Ajout(c cle.Cle) {
	curr := t
	if curr.Size == (2<<curr.height)-1 {
		// si t contient 2^h - 1 neouds,
		// on descend à gauche jusqu'au dernier niveau
		for curr.leftChild != nil {
			curr.Size += 1
			curr.height += 1
			curr = curr.leftChild
		}
		curr.Size += 1
		curr.height += 1
		curr.leftChild = &Arbre{cle: c, Size: 1, height: 1, parent: curr}
		curr = curr.leftChild
	} else {
		for {
			curr.Size += 1
			if curr.leftChild == nil {
				curr.height += 1
				curr.leftChild = &Arbre{cle: c, Size: 1, height: 1, parent: curr}

				currHeight := curr
				for currHeight.parent != nil {
					if currHeight.parent.height == currHeight.height+1 {
						break
					}
					currHeight.parent.height += 1
					currHeight = currHeight.parent
				}
				curr = curr.leftChild

				break
			} else if curr.rightChild == nil {
				curr.rightChild = &Arbre{cle: c, Size: 1, height: 1, parent: curr}
				curr = curr.rightChild
				break
			}
			if curr.leftChild.Size <= curr.rightChild.Size || curr.leftChild.Size != pow(2, curr.leftChild.height)-1 {
				curr = curr.leftChild
			} else {
				curr = curr.rightChild
			}
		} // for
	} // else
	for curr.parent != nil {
		// curr.parent.height = max(curr)
		if !curr.parent.cle.Inf(curr.cle) {
			temp := *curr
			curr.cle = curr.parent.cle
			curr.parent.cle = temp.cle
			curr = curr.parent
		} else {
			break
		}
	} // for
}

func (t *Arbre) AjoutIteratif(cles []cle.Cle) {
	setSigHandler()
	for index = 0; index < len(cles); index++ {
		t.Ajout(cles[index])
	}
}

func (t *Arbre) _string(indent int) string {
	str := "Tas{\n"
	indent += 1
	// str += strings.Repeat("│ ", indent) + "cle: " + t.cle.DecimalString() + ",\n"
	str += strings.Repeat("  ", indent) + "cle: " + t.cle.DecimalString() + ",\n"
	// str += strings.Repeat("  ", indent) + "size: " + fmt.Sprint(t.size) + ",\n"
	// str += strings.Repeat("  ", indent) + "height: " + fmt.Sprint(t.height) + ",\n"
	if t.leftChild != nil {
		str += strings.Repeat("  ", indent) + "left: " + t.leftChild._string(indent) + ",\n"
	} else {
		str += strings.Repeat("  ", indent) + "left: nil,\n"
	}
	if t.rightChild != nil {
		str += strings.Repeat("  ", indent) + "right: " + t.rightChild._string(indent) + "\n"
	} else {
		str += strings.Repeat("  ", indent) + "right: nil\n"
	}
	str += strings.Repeat("  ", indent-1) + "}"
	return str
}

func setSigHandler() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGQUIT)
	go func() {
		for {
			s := <-sigChan
			if s == syscall.SIGQUIT {
				fmt.Println(index)
			}
		}
	}()
}

func NewArbre(c cle.Cle) Arbre {
	return Arbre{cle: c, height: 1, Size: 1}
}

func pow(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func (t *Arbre) String() string {
	return t._string(0) + "\n"
}
