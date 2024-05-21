package encoder

import (
	"fmt"
	"log"
)

type Encoder struct {
	rs     *ReedSolomon
	size   int
	matrix [][]bool
}

func (e Encoder) PrintQrToAscii() {
	for _, row := range e.matrix {
		for _, col := range row {
			if col {
				fmt.Print("■ ")
			} else {
				fmt.Print("□ ")
			}
		}
		fmt.Println()
	}
}

func (e Encoder) encodeToBinary(data string) string {
	binary_data := ""
	for _, v := range data {
		binary_data += fmt.Sprintf("%08b", v)
	}
	return binary_data
}

func (e Encoder) EncodeWithErrorCorrection(data string) ([]int, error) {
	binaryData := e.encodeToBinary(data)

	msg := make([]int, len(binaryData))
	for i, bit := range binaryData {
		msg[i] = int(bit - '0')
	}

	log.Println("msg:", msg)

	encodedMsg, err := e.rs.encodeMsg(msg)
	if err != nil {
		return nil, err
	}

	return encodedMsg, nil
}

func (e *Encoder) createMatrix() [][]bool {
	version := e.size
	size := 21 + (version-1)*4

	matrix := make([][]bool, size)
	for i := range matrix {
		matrix[i] = make([]bool, size)
	}

	e.addPositionPatterns(matrix)
	// TODO: add other needed patterns
	e.matrix = matrix
	return matrix
}

// addPositionPatterns adiciona padrões de posição à matriz
func (e Encoder) addPositionPatterns(matrix [][]bool) {
	size := len(matrix)
	patterns := [3][2]int{{0, 0}, {size - 7, 0}, {0, size - 7}}
	for _, pos := range patterns {
		for y := 0; y < 7; y++ {
			for x := 0; x < 7; x++ {
				if x == 0 || x == 6 || y == 0 || y == 6 || (x >= 2 && x <= 4 && y >= 2 && y <= 4) {
					matrix[pos[0]+y][pos[1]+x] = true // preto
				}
			}
		}
	}
}

func NewEncoder(size int) Encoder {
	e := Encoder{newReedSolomon(3), size, [][]bool{}}
	e.createMatrix()
	return e
}
