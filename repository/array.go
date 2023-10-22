package repository

import (
	"log"
	"math/rand"
	"strconv"
)

type Transaction struct {
	ID int
	Valor float32
	Parcelas int
	Descricao string
}

func newTransaction(qty int) *Transaction {
	t := &Transaction{
		ID:        randomInt(10*qty, 10000*qty),
		Valor:     randomFloat32(float32(1000*qty+(1/qty)), float32(10000000*qty+(1/(qty*10)))),
		Parcelas:  randomInt(1, 24),
	}
	t.Descricao = randomStringFromSel("Parcela", "Acordo", "Pedido", "Contrato", "Item", "Solicitação", "Segmento", "Requisição", "Chamado", "Passagem", "Reserva", "Erro", "Processo", "Ação", "Procedimento") + " #" + strconv.Itoa(t.ID)
	return t
}

type Transactions struct {
	ToProcess []*Transaction
}

func NewTransactions(qty int) *Transactions {
	toProcess := make([]*Transaction, qty)
	for q := 0; q<qty ; q++ {
		t := newTransaction(qty)
		toProcess[q] = t
	}

	log.Printf("Tamanho: %d\v", len(toProcess))

	return &Transactions{
		ToProcess: toProcess,
	}
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomStringFromSel(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}
