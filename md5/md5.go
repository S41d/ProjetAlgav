package md5

import (
	"bytes"
	"encoding/binary"
	"math"
)

// r est une table de constantes utilisée dans l'algorithme MD5.
// Elle représente les rotations spécifiques pour chaque étape du calcul.
var r = []int{
	7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
	5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
	4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
	6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21,
}

// k est un tableau qui stockera les constantes dérivées des sinus d'entiers pour MD5.
var k []uint32

// initializeK initialise le tableau k avec les valeurs dérivées des sinus d'entiers.
func initializeK() {
	k = make([]uint32, 64)
	for i := 0; i < 64; i++ {
		// La formule utilise la fonction sinus et l'arrondi pour obtenir des valeurs entières.
		// Ces valeurs sont ensuite multipliées par 2^32 pour créer des constantes de 32 bits.
		k[i] = uint32(math.Floor(math.Abs(math.Sin(float64(i+1))) * math.Pow(2, 32)))
	}
}

// leftRotate effectue une rotation vers la gauche sur un entier de 32 bits.
func leftRotate(x uint32, c int) uint32 {
	return (x << uint(c)) | (x >> (32 - uint(c)))
}

// F est l'une des fonctions de compression de l'algorithme MD5.
// Elle effectue une opération de type "majorité" : si la majorité des bits
// aux positions correspondantes dans x, y et z sont à 1, le résultat est 1,
// sinon le résultat est 0. Les opérations bitwise (&, ^) sont utilisées pour
// réaliser cette opération logique spécifique.
func F(x, y, z uint32) uint32 {
	return (x & y) | ((^x) & z)
}

// G est l'une des fonctions de compression de l'algorithme MD5.
// Elle effectue une opération de type "parité" : si un nombre impair des bits
// aux positions correspondantes dans x, y et z sont à 1, le résultat est 1,
// sinon le résultat est 0. Les opérations bitwise (&, ^) sont utilisées pour
// réaliser cette opération logique spécifique.
func G(x, y, z uint32) uint32 {
	return (x & z) | (y & (^z))
}

// H est l'une des fonctions de compression de l'algorithme MD5.
// Elle effectue une opération de type "XOR" (ou exclusif) entre les bits
// aux positions correspondantes dans x, y et z. Cette opération logique
// spécifique est utilisée dans l'algorithme MD5.
func H(x, y, z uint32) uint32 {
	return x ^ y ^ z
}

// I est l'une des fonctions de compression de l'algorithme MD5.
// Elle effectue une opération de type "majorité inversée" : si la majorité
// des bits aux positions correspondantes dans x, y et z sont à 0, le résultat
// est 0, sinon le résultat est 1. Les opérations bitwise (&, ^) sont utilisées
// pour réaliser cette opération logique spécifique.
func I(x, y, z uint32) uint32 {
	return y ^ (x | (^z))
}

// md5New est une fonction qui calcule le hachage MD5 d'un message donné.
// Le résultat est renvoyé sous la forme d'un tableau de 16 octets.
func md5New(message []byte) [16]byte {
	// Initialiser le tableau k avec les valeurs dérivées des sinus d'entiers.
	initializeK()

	// Initialiser les variables de hachage MD5.
	h0 := uint32(0x67452301)
	h1 := uint32(0xEFCDAB89)
	h2 := uint32(0x98BADCFE)
	h3 := uint32(0x10325476)

	// Calculer la longueur originale du message en bits.
	originalLength := len(message)
	totalLength := uint64(originalLength * 8)

	// Créer un tampon pour le message, ajouter le bit "1" et le padding.
	buffer := new(bytes.Buffer)
	buffer.Write(message)
	buffer.WriteByte(0x80)
	for buffer.Len()%64 != 56 {
		buffer.WriteByte(0x00)
	}
	err := binary.Write(buffer, binary.LittleEndian, totalLength)
	if err != nil {
		return [16]byte{}
	}

	// Obtenir le message final après le padding.
	paddedMessage := buffer.Bytes()

	// Traiter le message par blocs de 512 bits.
	for offset := 0; offset < len(paddedMessage); offset += 64 {
		block := paddedMessage[offset : offset+64]

		// Diviser le bloc en 16 mots de 32 bits en little-endian.
		w := make([]uint32, 16)
		for i := 0; i < 16; i++ {
			w[i] = binary.LittleEndian.Uint32(block[i*4 : (i+1)*4])
		}

		// Initialiser les valeurs de hachage.
		a := h0
		b := h1
		c := h2
		d := h3

		// Boucle principale de traitement du bloc.
		for i := 0; i < 64; i++ {
			var f, g uint32
			switch {
			case i < 16:
				f = F(b, c, d)
				g = uint32(i)
			case i < 32:
				f = G(b, c, d)
				g = uint32((5*i + 1) % 16)
			case i < 48:
				f = H(b, c, d)
				g = uint32((3*i + 5) % 16)
			default:
				f = I(b, c, d)
				g = uint32((7 * i) % 16)
			}

			temp := d
			d = c
			c = b
			b = b + leftRotate((a+f+k[i]+w[g]), r[i])
			a = temp
		}

		// Ajouter le résultat au bloc précédent.
		h0 += a
		h1 += b
		h2 += c
		h3 += d
	}

	// Construire le résultat final en little-endian.
	result := [16]byte{}
	binary.LittleEndian.PutUint32(result[0:4], h0)
	binary.LittleEndian.PutUint32(result[4:8], h1)
	binary.LittleEndian.PutUint32(result[8:12], h2)
	binary.LittleEndian.PutUint32(result[12:16], h3)

	return result
}
