package usecase

import (
	"log"
	"time"

	"github.com/renatospaka/poc-channel/repository"
)

func ProcessInLine() error {
	log.Println("in-line at usecase")
	now := time.Now()
	toProcess, _ := repository.ProcessInLine()
	newTP := repository.Transactions{}

	for x := 0; x < len(toProcess.ToProcess); x++ {
		newTP.ToProcess = append(newTP.ToProcess, toProcess.ToProcess[x])
		log.Printf("in-line at usecase ==> ID: %d | Parcelas: %d | Valor: %f | Descrição: %s\n", toProcess.ToProcess[x].ID, toProcess.ToProcess[x].Parcelas, toProcess.ToProcess[x].Valor, toProcess.ToProcess[x].Descricao)
	}

	log.Printf("in-line at usecase: %s for %d items\n", time.Since(now).String(), len(newTP.ToProcess))
	return nil
}

func ProcessChannel() error {
	log.Println("channel at usecase")
	now := time.Now()
	processed := 0
	toProcess := make(chan *repository.Transaction)

	go repository.ProcessChannel(toProcess)
	for t := range toProcess {
		processed++
		log.Printf("channel at usecase ==> ID: %d | Parcelas: %d | Valor: %f | Descrição: %s\n", t.ID, t.Parcelas, t.Valor, t.Descricao)
	}

	log.Printf("channel at usecase: %s for %d items\n", time.Since(now).String(), processed)
	return nil
}
