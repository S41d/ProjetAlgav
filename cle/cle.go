package cle

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"strings"
)

type Cle struct {
	p1 uint64
	p2 uint64
}

func (c Cle) Inf(o Cle) bool {
	p1inf := c.p1 < o.p1
	p2inf := c.p2 < o.p2
	return p1inf || (p1inf && p2inf)
}

func (c Cle) Eg(o Cle) bool {
	return c.p1 == o.p1 && c.p2 == o.p2
}

func (c Cle) HexaString() string {
	if c.p2 == 0 {
		return fmt.Sprintf("%x", c.p1)
	}
	if c.p1 == 0 {
		return fmt.Sprintf("%x", c.p2)
	}
	return fmt.Sprintf("%x%x", c.p1, c.p2)
}
func (c Cle) BinaryString() string {
	return fmt.Sprintf("%b%064b\n", c.p1, c.p2)
}

func (c Cle) DecimalString() string {
	buf := new(bytes.Buffer)
	if c.p1 != 0 {
		err := binary.Write(buf, binary.BigEndian, c.p1)
		if err != nil {
			fmt.Println("binary.Write failed on p1:", err)
		}
	}
	if c.p2 != 0 {
		err := binary.Write(buf, binary.BigEndian, c.p2)
		if err != nil {
			fmt.Println("binary.Write failed on p2:", err)
		}
	}
	num := big.NewInt(0)
	num.SetBytes(buf.Bytes())
	return fmt.Sprintf("%v", num)

}

func HexToCle(hex string) Cle {
	c := Cle{}

	k := 0
	if strings.HasPrefix(hex, "0x") {
		hex = strings.TrimLeft(hex, "0x")
	}
	for i := 0; i < len(hex); i++ {
		char := hex[i]
		if char >= '0' && char <= '9' {
			char = char - '0'
		} else if char >= 'a' && char <= 'f' {
			char = char - 'a' + 10
		} else if char >= 'A' && char <= 'F' {
			char = char - 'A' + 10
		}

		// chaque lettre/chiffre hexa = 4 bit
		// 64/4 = 16
		if k < 16 {
			c.p1 = (c.p1 << 4) | uint64((char & 0xF))
		} else {
			c.p2 = (c.p2 << 4) | uint64((char & 0xF))
		}
		k++
	}
	return c
}
