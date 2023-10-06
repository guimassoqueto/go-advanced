package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

// simula um trabalho que vai demorar randomicamente entre 1 e 5 segundos para ocorrer
func doWork() int {
	randomTime := 1 + rand.Intn(5)
	fmt.Printf("sleeping for: %d seconds\n", randomTime)
	time.Sleep(time.Duration(randomTime) * time.Second)
	return randomTime
}


func main() {
	// inicia um canal que recebe um inteiro, o segundo parametro é opcional e representa
	// o tamanho do buffer de inteiros
	dataChan := make(chan int, 2)
	processes := 5

	go func() {
		var wg sync.WaitGroup

		// executa 5 processos concorrentes (cada um em uma thread)
		for i := 0; i < processes; i++ {
			wg.Add(1)
			go func() {
				// a funcao dorme por um periodo randomico de segundos e devolve 
				// esse valor para o canal
				dataChan <- doWork()
				wg.Done()
			}()
		}
		wg.Wait()

		// fecha o canal
		close(dataChan)
	}()
	
	// a thread principal (executando o programa), vai receber em tempo real 
	// os inteiros recebidos nas threads paralelas, vai dormir e avisar quando terminar.
	// note que aqui a funcao sleep nao executa concorrentemente, apenas uma de cada vez.
	for n := range dataChan {
		time.Sleep(time.Duration(n) * time.Second)
		fmt.Printf("sleeped for: %d seconds...\n", n)
	}
}