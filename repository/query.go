package repository

import (
	"log"
	// "runtime"
	"time"
)

var data *Transactions

func ProcessInLine() (*Transactions, error) {
	processed := &Transactions{}
	now := time.Now()
	data = NewTransactions(75000)

	for x := 0; x < len(data.ToProcess); x++ {
		processed.ToProcess = append(processed.ToProcess, data.ToProcess[x])
		// log.Printf("in-line at Repository ==> processadores: %d | NumGoroutine: %d |  ID: %d | Parcelas: %d | Valor: %f | Descrição: %s\n", runtime.NumCPU(), runtime.NumGoroutine(), processed.ToProcess[x].ID, processed.ToProcess[x].Parcelas, processed.ToProcess[x].Valor, processed.ToProcess[x].Descricao)
	}
	log.Printf("in-line at Repository: %s for %d items\n", time.Since(now).String(), len(processed.ToProcess))
	log.Println()
	return processed, nil
}

func ProcessChannel(chTransaction chan<- *Transaction) error {
	processed := 0
	now := time.Now()
	data = NewTransactions(75000)

	for ch := 0; ch < len(data.ToProcess); ch++ {
		// log.Printf("channel at Repository ==> processadores: %d | NumGoroutine: %d | ID: %d | Parcelas: %d | Valor: %f | Descrição: %s\n", runtime.NumCPU(), runtime.NumGoroutine(), data.ToProcess[processed].ID, data.ToProcess[processed].Parcelas, data.ToProcess[processed].Valor, data.ToProcess[processed].Descricao)
		chTransaction <- data.ToProcess[ch]
		processed++
	}
	// close(chTransaction)

	log.Printf("channel at Repository: %s for %d items\n", time.Since(now).String(), processed)
	log.Println()
	return nil
}
