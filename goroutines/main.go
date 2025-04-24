package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// Quantidade de processo que ele vai esperar, e isso vai diminuindo conforme o numero de Done for sendo chamado
	wg.Add(3)
	go callDatabase(&wg)
	go callAPI(&wg)
	go processInternal(&wg)
	wg.Wait()
}

func callDatabase(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Printf("callDatabase\n")
	wg.Done()
}

func callAPI(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Printf("callAPI\n")
	wg.Done()
}

func processInternal(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Printf("processInternal\n")
	wg.Done()
}
