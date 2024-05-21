package main

import (
	"github.com/pvskp/qr2go/encoder"
)

func main() {
	e := encoder.NewEncoder(1)
	// e.PrintQrToAscii()
	e.EncodeWithErrorCorrection("Hello, World!")
}
