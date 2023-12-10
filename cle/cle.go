package cle

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"strings"
)

// Cle est une structure représentant une clé codés sur 128 bits composée de deux parties uint64.
type Cle struct {
	// P1 est la première partie de la clé.
	P1 uint64

	// P2 est la deuxième partie de la clé.
	P2 uint64
}

// Inf compare deux clés et retourne vrai si la première clé est inférieure à la seconde.
// La comparaison se fait d'abord sur la première partie (P1). Si les premières parties sont égales,
// la comparaison se fait ensuite sur la deuxième partie (P2).
// La fonction retourne true si c.P1 < o.P1 ou (c.P1 == o.P1 et c.P2 < o.P2), sinon elle retourne false.
func (c Cle) Inf(o Cle) bool {
	return (c.P1 < o.P1) || ((c.P1 == o.P1) && (c.P2 < o.P2))
}

// Eg compare la clé actuelle (c) avec une autre clé (o) et retourne true si
// les deux clés sont égales (ont les mêmes valeurs pour les deux parties),
// sinon elle retourne false.
//
// La fonction retourne true si la première partie de la clé actuelle est égale
// à la première partie de la clé donnée en paramètre, et si la deuxième partie
// de la clé actuelle est égale à la deuxième partie de la clé donnée en paramètre.
// Sinon, elle retourne false.
func (c Cle) Eg(o Cle) bool {
	// Retourne true si les deux parties de la clé sont égales, sinon retourne false.
	return c.P1 == o.P1 && c.P2 == o.P2
}

// HexaString retourne une représentation hexadécimale de la clé sous la forme d'une chaîne de caractères.
// Si la deuxième partie de la clé est zéro, seule la première partie est utilisée pour la représentation hexadécimale.
// Si la première partie de la clé est zéro, seule la deuxième partie est utilisée pour la représentation hexadécimale.
// Si les deux parties de la clé ne sont pas zéro, les deux parties sont utilisées pour la représentation hexadécimale.
//
// La fonction utilise le format %x de la fonction Sprintf de la bibliothèque fmt pour générer la représentation hexadécimale.
func (c Cle) HexaString() string {
	// Si la deuxième partie de la clé est zéro, utiliser seulement la première partie pour la représentation hexadécimale.
	if c.P2 == 0 {
		return fmt.Sprintf("%x", c.P1)
	}

	// Si la première partie de la clé est zéro, utiliser seulement la deuxième partie pour la représentation hexadécimale.
	if c.P1 == 0 {
		return fmt.Sprintf("%x", c.P2)
	}

	// Si les deux parties de la clé ne sont pas zéro, utiliser les deux parties pour la représentation hexadécimale.
	return fmt.Sprintf("%x%x", c.P1, c.P2)
}

// BinaryString retourne une représentation binaire de la clé sous la forme d'une chaîne de caractères.
// La première partie de la clé est représentée en binaire à l'aide du format %b de la fonction Sprintf de la bibliothèque fmt.
// La deuxième partie de la clé est représentée en binaire avec un remplissage de zéros à gauche pour atteindre une longueur de 64 bits.
//
// La fonction utilise le format %b de la fonction Sprintf de la bibliothèque fmt pour générer la représentation binaire.
func (c Cle) BinaryString() string {
	// Utiliser le format %b pour la première partie de la clé.
	// Utiliser %064b pour la deuxième partie avec un remplissage de zéros à gauche pour atteindre une longueur de 64 bits.
	return fmt.Sprintf("%064b%064b\n", c.P1, c.P2)
}

// DecimalString retourne une représentation décimale de la clé sous la forme d'une chaîne de caractères.
// Les deux parties de la clé sont converties en un seul entier décimal en utilisant le codage binaire BigEndian,
// puis le résultat est formatté en utilisant le type big.Int et la fonction Sprintf de la bibliothèque fmt.
//
// La fonction utilise le paquet bytes pour créer un tampon (buffer) et le paquet binary pour écrire les parties de la clé
// dans le tampon en utilisant le codage binaire BigEndian. Ensuite, elle utilise le paquet math/big pour créer un entier
// décimal à partir des données du tampon et formate le résultat en utilisant la fonction Sprintf de la bibliothèque fmt.
func (c Cle) DecimalString() string {
	// Créer un tampon (buffer) pour stocker les données binaires.
	buf := new(bytes.Buffer)

	// Écrire la première partie de la clé dans le tampon si elle n'est pas zéro.
	if c.P1 != 0 {
		err := binary.Write(buf, binary.BigEndian, c.P1)
		if err != nil {
			fmt.Println("binary.Write failed on P1:", err)
		}
	}

	// Écrire la deuxième partie de la clé dans le tampon si elle n'est pas zéro.
	if c.P2 != 0 {
		err := binary.Write(buf, binary.BigEndian, c.P2)
		if err != nil {
			fmt.Println("binary.Write failed on P2:", err)
		}
	}

	// Créer un entier décimal à partir des données du tampon.
	num := big.NewInt(0)
	num.SetBytes(buf.Bytes())

	// Formater le résultat en utilisant la fonction Sprintf de la bibliothèque fmt.
	return fmt.Sprintf("%v", num)
}

// HexToCle convertit une chaîne hexadécimale en une structure Cle.
// La fonction prend en compte les préfixes "0x" éventuels dans la chaîne hexadécimale et les ignore.
// La chaîne hexadécimale est parcourue caractère par caractère, chaque caractère est converti en valeur hexadécimale,
// et les valeurs hexadécimales sont assignées aux parties correspondantes de la structure Cle (P1 et P2).
// Chaque lettre/chiffre hexadécimal correspond à 4 bits, et la fonction utilise cette information pour remplir les parties
// de la structure Cle en conséquence.
func HexToCle(hex string) Cle {
	c := Cle{}

	// Vérifier si la chaîne hexadécimale commence par "0x" et la tronquer si c'est le cas.
	if strings.HasPrefix(hex, "0x") {
		hex = strings.TrimLeft(hex, "0x")
	}

	k := 0
	for i := 0; i < len(hex); i++ {
		char := hex[i]

		// Convertir le caractère hexadécimal en valeur hexadécimale.
		if char >= '0' && char <= '9' {
			char -= '0'
		} else if char >= 'a' && char <= 'f' {
			char = char - 'a' + 10
		} else if char >= 'A' && char <= 'F' {
			char = char - 'A' + 10
		}

		// Chaque lettre/chiffre hexadécimal correspond à 4 bits.
		// Remplir les parties de la structure Cle en conséquence.
		if k < 16 {
			c.P1 = (c.P1 << 4) | uint64(char&0xF)
		} else {
			c.P2 = (c.P2 << 4) | uint64(char&0xF)
		}
		k++
	}
	return c
}

func BytesToCle(bytes [16]byte) Cle {
	c := Cle{}
	for i := 0; i < 8; i++ {
		c.P1 = (c.P1 << 8) | uint64(bytes[i])
	}
	for i := 8; i < 16; i++ {
		c.P2 = (c.P1 << 8) | uint64(bytes[i])
	}
	return c
}
