package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	waitGroup()
	fmt.Printf("\n")
	mutex()
}

// Garatir que diferentes goroutines não altere o mesmo dados, usando o lock
func mutex() {
	var m sync.Mutex

	i := 0

	for range 1_000 {
		go func() {
			m.Lock()
			// Somente essa goroutine esta alteradando a variavel i
			i++
			m.Unlock()
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(i)
}

func waitGroup() {
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
