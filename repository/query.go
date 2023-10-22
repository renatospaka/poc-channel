package repository

import (
	"log"
	"time"
)

var data *Transactions

func ProcessInLine() (*Transactions, error) {
	processed := &Transactions{}
	now := time.Now()
	data = NewTransactions(20)

	// processed.ToProcess = append(processed.ToProcess, data.ToProcess...)
	for x := 0; x<len(data.ToProcess);x++ {
		processed.ToProcess = append(processed.ToProcess, data.ToProcess[x])
		log.Printf("in-line at Repository ==> ID: %d | Parcela: %d | Valor: %f | Descrição: %s\n", processed.ToProcess[x].ID, processed.ToProcess[x].Parcelas, processed.ToProcess[x].Valor, processed.ToProcess[x].Descricao)
	}
	log.Printf("in-line at Repository: %s for %d items\n", time.Since(now).String(), len(processed.ToProcess))
	log.Println()
	return processed, nil
}

func ProcessChannel(chTransaction chan *Transaction) error {
	processed := 0
	now := time.Now()
	data = NewTransactions(20)

	for ch := 0; ch < len(data.ToProcess); ch++ {
		chTransaction <- data.ToProcess[ch]
		log.Printf("channel at Repository ==> ID: %d | Parcela: %d | Valor: %f | Descrição: %s\n", data.ToProcess[processed].ID, data.ToProcess[processed].Parcelas, data.ToProcess[processed].Valor, data.ToProcess[processed].Descricao)
		processed++
	}
	close(chTransaction)

	log.Printf("channel at Repository: %s for %d items\n", time.Since(now).String(), processed)
	log.Println()
	return nil
}
