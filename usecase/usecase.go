package usecase

import (
	"log"
	"time"

	"github.com/renatospaka/poc-channel/repository"
)

func ProcessInLine() error {
	now := time.Now()
	log.Println("in-line")

	toProcess, _ := repository.ProcessInLine()
	newTP := repository.Transactions{}
	newTP.ToProcess = append(newTP.ToProcess, toProcess.ToProcess...)
	
	log.Printf("in-line at Usecase: %s for %d items\n", time.Since(now).String(), len(newTP.ToProcess))
	return nil
}