package usecase

import (
	"log"
	"time"

	"github.com/renatospaka/poc-channel/repository"
)

func ProcessInLine() error {
	now := time.Now()
	log.Println("in-line at usecase")

	toProcess, _ := repository.ProcessInLine()
	newTP := repository.Transactions{}
	for x := 0; x<len(toProcess.ToProcess);x++ {
		newTP.ToProcess = append(newTP.ToProcess, toProcess.ToProcess[x])
		log.Printf("in-line at Usecase ==> ID: %d | Parcela: %d | Valor: %f | Descrição: %s\n", toProcess.ToProcess[x].ID, toProcess.ToProcess[x].Parcelas, toProcess.ToProcess[x].Valor, toProcess.ToProcess[x].Descricao)
	}
	
	log.Printf("in-line at Usecase: %s for %d items\n", time.Since(now).String(), len(newTP.ToProcess))
	return nil
}

func ProcessChannel() error {
	processed := 0
	now := time.Now()
	toProcess := make(chan *repository.Transaction)
	log.Println("channel at usecase")

	go repository.ProcessChannel(toProcess)
	for t := range toProcess {
		processed++
		
		log.Printf("channel at Usecase ==> ID: %d | Parcela: %d | Valor: %f | Descrição: %s\n", t.ID, t.Parcelas, t.Valor, t.Descricao)
		// if processed < 2 {
		// 	log.Printf("ID: %d\n", t.ID)
		// }		
	}
	
	log.Printf("channel at Usecase: %s for %d items\n", time.Since(now).String(), processed)
	return nil
}