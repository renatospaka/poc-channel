package main

import (
	"log"

	"github.com/renatospaka/poc-channel/usecase"
)

func main() {
	log.Println("iniciando...")

	qtdeInLine, durationInLine := usecase.ProcessInLine()
	log.Println()
	log.Println()

	qtdeChannel, durationChannel := usecase.ProcessChannel()
	log.Println()
	log.Println()

	log.Println("fim...")
	log.Printf("in-line: %s for %d items\n", durationInLine, qtdeInLine)
	log.Printf("channel: %s for %d items\n", durationChannel, qtdeChannel)
}
