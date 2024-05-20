package main

import (
	"log"

	"github.com/pvskp/qr2go/encoder"
)

func main() {
	e := encoder.NewEncoder()
	log.Println(e.EncodeWithErrorCorrection("meunomeehpaulo"))
}
