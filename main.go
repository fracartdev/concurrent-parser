package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// parsing del CSV to JSON, per ora simulo con un timeout
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// avvio di 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// invio i jobs ai worker avviati in precedenza
	// poi chiudo il channel per indicare di avere finito
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// risultato
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
