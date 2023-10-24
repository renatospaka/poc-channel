package usecase

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/renatospaka/poc-channel/repository"
)

func ProcessInLine() (int, time.Duration) {
	log.Println("in-line at usecase")
	now := time.Now()
	newTP := repository.Transactions{}
	
	toProcess, _ := repository.ProcessInLine()
	for x := 0; x < len(toProcess.ToProcess); x++ {
		newTP.ToProcess = append(newTP.ToProcess, toProcess.ToProcess[x])
		log.Printf("in-line at usecase ==> ID: %d | Parcelas: %d | Valor: %f | Descrição: %s\n", toProcess.ToProcess[x].ID, toProcess.ToProcess[x].Parcelas, toProcess.ToProcess[x].Valor, toProcess.ToProcess[x].Descricao)
	}

	// log.Printf("in-line at usecase: %s for %d items\n", time.Since(now).String(), len(newTP.ToProcess))
	return len(newTP.ToProcess), time.Since(now)
}

func ProcessChannel() (int, time.Duration) {
	log.Println("channel at usecase")
	now := time.Now()
	var processed int32 = 0
	
	toProcess := make(chan *repository.Transaction)
	go func(){
		defer close(toProcess)
		repository.ProcessChannel(toProcess)
	}()

	wg := sync.WaitGroup{}
	for p := 0; p<5; p++ {
		wg.Add(1)
		go func(){
			defer wg.Done()

			for t := range toProcess {
				atomic.AddInt32(&processed, 1)
				log.Printf("channel at usecase ==> ID: %d | Parcelas: %d | Valor: %f | Descrição: %s\n", t.ID, t.Parcelas, t.Valor, t.Descricao)
			}
		}()
	}
	wg.Wait()


	// log.Printf("channel at usecase: %s for %d items\n", time.Since(now).String(), processed)
	return int(processed), time.Since(now)
}
