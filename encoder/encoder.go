package encoder

import (
	"fmt"
)

type Encoder struct {
	rs *ReedSolomon
}

func (e Encoder) encodeToBinary(data string) string {
	binary_data := ""
	for _, v := range data {
		binary_data += fmt.Sprintf("%08b", v)
	}
	return binary_data
}

func (e Encoder) EncodeWithErrorCorrection(data string) ([]int, error) {
	// Converte a string em dados binários
	binaryData := e.encodeToBinary(data)

	// Converte a string binária para um slice de inteiros
	msg := make([]int, len(binaryData))
	for i, bit := range binaryData {
		msg[i] = int(bit - '0')
	}

	// Adiciona correção de erros usando Reed-Solomon
	encodedMsg, err := e.rs.encodeMsg(msg)
	if err != nil {
		return nil, err
	}

	return encodedMsg, nil
}

func NewEncoder() Encoder {
	return Encoder{newReedSolomon(3)}
}
