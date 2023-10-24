package main

import (
	"log"
	"runtime"
	"time"

	"github.com/renatospaka/poc-channel/usecase"
)

func main() {
	log.Println("iniciando...")
	log.Printf("runtime NumGoroutine: %d\n", runtime.NumGoroutine())

	qtdeInLine, durationInLine := usecase.ProcessInLine()
	log.Println()
	log.Println()

	qtdeChannel, durationChannel := usecase.ProcessChannel()
	log.Println()
	log.Println()

	log.Println("fim...")
	time.Sleep(700*time.Millisecond)
	log.Printf("in-line: %s for %d items\n", durationInLine, qtdeInLine)
	log.Printf("channel: %s for %d items\n", durationChannel, qtdeChannel)
}
