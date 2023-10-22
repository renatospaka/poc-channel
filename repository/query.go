package repository

import (
	"log"
	"time"
)

var data *Transactions

func init() {
}

func ProcessInLine() (*Transactions, error) {
	processed := &Transactions{}
	now := time.Now()
	data = NewTransactions(3000)

	processed.ToProcess = append(processed.ToProcess, data.ToProcess...)
	log.Printf("in-line at Repository: %s for %d items\n", time.Since(now).String(), len(processed.ToProcess))
	return processed, nil
}