package main

import (
	"../../runner"
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")
	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())

	// Start 中通过Select从多个channel中取值
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(1)
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor- Task#%d\n", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
