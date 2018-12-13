package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("waiting for sleep...")
	err = cmd.Wait()
	log.Printf("Done: %v\n", err)
}
