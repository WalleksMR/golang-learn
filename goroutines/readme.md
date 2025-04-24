As **goroutines** são uma das principais características do Go para lidar com concorrência. Elas são unidades leves de execução gerenciadas pelo runtime do Go, permitindo que múltiplas operações sejam executadas de forma assíncrona e eficiente. Vamos explorar como as goroutines funcionam com base no código fornecido, destacando três conceitos importantes: **WaitGroup**, **Mutex** e **Channels**.

## 1. WaitGroup: Sincronização de Goroutines

O `sync.WaitGroup` é usado para aguardar que um conjunto de goroutines termine sua execução antes de prosseguir.

**Exemplo no Código:**

```go
func waitGroup() {
    var wg sync.WaitGroup
    wg.Add(3) // Define 3 goroutines para aguardar

    go callDatabase(&wg)
    go callAPI(&wg)
    go processInternal(&wg)

    wg.Wait() // Bloqueia até que todas as goroutines finalizem
}
```

- **Funcionamento**:

  - `wg.Add(3)`: Indica que o WaitGroup deve esperar 3 goroutines.

  - Cada goroutine chama `wg.Done()` ao finalizar, decrementando o contador.
  - `wg.Wait()`: Bloqueia a thread principal até que o contador chegue a zero.

- **Uso Prático:** Útil para coordenar tarefas paralelas (como chamadas a APIs, bancos de dados, etc.) e garantir que todas terminem antes de continuar.

## 2. Mutex: Controle de Acesso a Recursos Compartilhados

O `sync.Mutex` (mutual exclusion lock) é usado para evitar **condições de corrida** _(race condition)_ quando múltiplas goroutines acessam um recurso compartilhado.

**Exemplo no Código:**

```go
func mutex() {
    var m sync.Mutex
    i := 0

    for range 1_000 {
        go func() {
            m.Lock() // Bloqueia o acesso à variável i
            i++     // Operação segura
            m.Unlock()
        }()
    }
    time.Sleep(time.Second * 1) // (!) Não é a melhor prática
    fmt.Println(i) // Deve imprimir 1000
}
```

- **Funcionamento**:

  - `m.Lock()`: Garante que apenas uma goroutine por vez acesse o código bloqueado.

  - `m.Unlock()`: Libera o acesso para outras goroutines.

- **Problema no Exemplo**: O uso de `time.Sleep` para aguardar as goroutines é frágil. O ideal seria usar um `WaitGroup`.

## 3. Channels: Comunicação entre Goroutines

**Channels** são canais de comunicação seguros para troca de dados entre goroutines. Eles garantem sincronização e evitam condições de corrida.

**Exemplo no Código:**

```go
func main() {
    channel := make(chan int)
    go setList(channel) // Produz valores no channel

    // Consumidor 1: Lê um valor
    channelRead := <-channel
    fmt.Println("channelRead", channelRead)

    // Consumidor 2: Lê valores restantes com range
    for value := range channel {
        fmt.Println("range", value)
    }
}

func setList(channel chan int) {
    for value := range 10 { // (!) Erro: range 10 não é válido em Go
        channel <- value // Envia valor para o channel
    }
    close(channel) // Fecha o channel após enviar todos os valores
}
```

- **Funcionamento:**

  - **Produtor:** setList envia valores para o channel (channel <- value).

  - **Consumidor:** A função principal lê os valores (<-channel e range channel).

  - `close(channel)`: Indica que não há mais valores a serem enviados.

- **Problema no Exemplo:** A expressão `range 10` é inválida. O correto seria `for i := 0; i < 10; i++`.

**Channel Buffer**

TODO: add

### Pontos-Chave:

**Goroutines**:

- São iniciadas com `go func()`. Executam de forma concorrente.

- Consomem poucos recursos comparadas a threads tradicionais.

**WaitGroup**:

- Coordena múltiplas goroutines, garantindo que todas terminem.

**Mutex**:

- Protege dados compartilhados de acesso concorrente.

**Channels**:

- Permitem comunicação segura entre goroutines.

- Devem ser fechados quando não há mais dados a enviar (`close(channel)`).
