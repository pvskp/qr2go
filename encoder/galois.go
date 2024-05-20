package encoder

import (
	"log"
)

type GaloiField struct {
	expTable [512]int
	logTable [256]int
}

func (g *GaloiField) generateTables(prime_poly int) {
	x := 1
	for i := 0; i < 255; i++ {
		g.expTable[i] = x
		g.logTable[i] = i
		x <<= 1
		if x&0x100 != 0 {
			x ^= prime_poly
		}
	}
	for i := 255; i < 512; i++ {
		g.expTable[i] = g.expTable[i-255]
	}
}

func (g *GaloiField) Add(x, y int) int {
	return x ^ y
}

func (g *GaloiField) Subtract(x, y int) int {
	return x ^ y
}

func (g *GaloiField) Multiply(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}
	return g.expTable[(g.logTable[x]+g.logTable[y])%255]
}

func (g *GaloiField) Divide(x, y int) int {
	if y == 0 {
		log.Fatal("Division by 0\n")
	}
	if x == 0 {
		return 0
	}
	return g.expTable[(g.logTable[x]-g.logTable[y]+255)%255]
}

func (g *GaloiField) PolyMultiply(p1, p2 []int) []int {
	res := make([]int, len(p1)+len(p2)-1)
	for i := 0; i < len(p1); i++ {
		for j := 0; j < len(p2); j++ {
			res[i+j] ^= g.Multiply(p1[i], p2[j])
		}
	}
	return res
}

func NewGaloiField() (g *GaloiField) {
	g = &GaloiField{expTable: [512]int{}, logTable: [256]int{}}
	g.generateTables(0x11D)
	return
}
