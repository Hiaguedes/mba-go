package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

// problema basico de concorrencia

// pra testar rode 1000 conexoes com 100 simultaneas como ab
// ab -n 10000 -c 100 http://localhost:3000/

var number uint64 = 0 // podemos ter concorrencia no uso dessa variavel, mesma thread pra mesma variavel

// rodando com go run -race 25-multithread/concurrency/* - race pode me apontar problemas de race condition
func main() {

	// m := sync.Mutex{}  // solucao com mutex

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// m.Lock()
		// number++
		// m.Unlock()

		atomic.AddUint64(&number, 1) // adiciona de maneira a nao perder a referencia pra n chamadas de maneira concorrente

		time.Sleep(300 * time.Millisecond)
		fmt.Fprintf(w, "Hello User number %d!\n", number)
	})

	
	log.Println("Rodando na porta 3000")
	http.ListenAndServe(":3000", nil)
}