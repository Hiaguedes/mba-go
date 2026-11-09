package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func main() {

	// Habilitar trace do GC
	// debug.SetGCPercent(-1) // desativa o gc automatico
	// ajusta o percentual do GC para 100% (padrão é 100)
	// debug.SetGCPercent(300)

	debug.SetMemoryLimit(100 * 1204 * 1024) // desativa o limite de memória do GC

	allocateMemory := func(size int) []byte {
		return make([]byte, size)
	}

	// Alocando memoria pra observar o comportamento do GC
	for i := 0; i < 10; i++ {
		_ = allocateMemory(20 * 1024 * 1024)          // aloca 20MB de memória
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
		fmt.Printf("TotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
		fmt.Printf("Sys = %v MiB\n", m.Sys/1024/1024)
		fmt.Printf("NumGC = %v\n", m.NumGC)
		fmt.Printf("Lookups = %v\n", m.Lookups)
		fmt.Printf("Mallocs = %v\n", m.Mallocs)
		fmt.Printf("Frees = %v\n", m.Frees)
		fmt.Printf("HeapAlloc = %v MiB\n", m.HeapAlloc/1024/1024)
		fmt.Printf("HeapSys = %v MiB\n", m.HeapSys/1024/1024)
		fmt.Printf("HeapReleased = %v MiB\n", m.HeapReleased/1024/1024)
		fmt.Printf("HeapInUse = %v MiB\n", m.HeapInuse/1024/1024)
		fmt.Printf("HeapIdle = %v MiB\n", m.HeapIdle/1024/1024)
		fmt.Printf("HeapObjects = %v\n", m.HeapObjects)
		fmt.Println("===============================")
		time.Sleep(1 * time.Second) // aloca memoria a cada 1s
	}


}

/* 
quanto mais roda o gc mais memoria e tempo ele consome, por isso é importante ajustar o GOGC para um valor adequado para a aplicação, caso contrário o GC pode se tornar um gargalo de performance.
quanto maior o memory limit, mais memoria o GC vai consumir, por isso é importante ajustar o memory limit para um valor adequado para a aplicação, caso contrário o GC pode se tornar um gargalo de performance.
*/