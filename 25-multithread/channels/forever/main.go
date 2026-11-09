// package main

// func main() {
// 	forever := make(chan bool)

// 	<-forever // esperando canal ficar cheio, segura o processo do main por enquanto

// 	/*

// 		possivel deadlock

// 		fatal error: all goroutines are asleep - deadlock!

// 			goroutine 1 [chan receive]:
// 			main.main()
// 					C:/Users/hiago/mba-go/25-multithread/channels/forever/main.go:6 +0x25
// 			exit status 2
// 	*/
// }

package main

func main() {
	forever := make(chan bool)

	go func(){
		for i := range(10){
			println(i)
		}
		forever <- true // preencho o valor do canal pra nao dar deadlock e liberar go routine
	}()

	<-forever
}