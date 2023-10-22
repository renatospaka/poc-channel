package main

import (
	"log"

	"github.com/renatospaka/poc-channel/usecase"
)

func main() {
	log.Println("iniciando...")

	// array := repository.NewTransactions(300)
	// for l := 0; l<10; l++ {
	// 	log.Printf("Descrição: %s | ID: %d | Valor: %f | Parcelas: %d", array.ToProcess[l].Descricao, array.ToProcess[l].ID, array.ToProcess[l].Valor, array.ToProcess[l].Parcelas)
	// }
	usecase.ProcessInLine()
	log.Println()
	log.Println()
	usecase.ProcessChannel()
	log.Println()
	log.Println()
}
