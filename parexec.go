package main

import (
	"runtime"
	"sync"
)

type ParExecutor func(string)

func worker(f ParExecutor, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		f(job)
	}
}

func ExecPar(wg *sync.WaitGroup, work []string, f ParExecutor) {
	numCPUs := runtime.NumCPU()
	wg.Add(numCPUs)

	// Gesamtproblem
	numResults := len(work)

	jobs := make(chan string, numResults)

	// Starte so viele Worker wie CPUs da sind
	for i := 0; i < numCPUs; i++ {
		go worker(f, jobs, wg)
	}

	// befuelle Daten
	for i := 0; i < numResults; i++ {
		jobs <- work[i]
	}

	close(jobs)
}

func ExecParWait(work []string, f ParExecutor) {
	var wg sync.WaitGroup
	ExecPar(&wg, work, f)
	wg.Wait()
}
