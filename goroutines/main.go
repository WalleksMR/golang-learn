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

	fmt.Printf("\n")

	fmt.Printf("channel\n\n")
	channel := make(chan int)

	go setList(channel)
	// ler do channer
	channelRead := <-channel
	fmt.Println("channelRead", channelRead)

	// ler do channer pelo range
	for value := range channel {
		fmt.Println("range", value)
	}

}

// Channel
func setList(channel chan int) {

	for value := range 10 {
		// escrever no channel
		channel <- value
	}
	// Precisa fechar o canal, pois nao vai ter nenhum outra goroutines atribuindo a esse canal.
	close(channel)
}

// Garatir que diferentes goroutines não altere o mesmo dados, usando o lock
func mutex() {
	fmt.Printf("mutex\n\n")

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
	time.Sleep(time.Second * 1)
	fmt.Println(i)
}

func waitGroup() {
	fmt.Printf("waitGroup\n\n")
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
