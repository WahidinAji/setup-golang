package main

import "time"

func main() {
	go func() {
		startServer()
		print("\n")
		commands()
	}()
	time.Sleep(1 * time.Second)
}
