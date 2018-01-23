package main

import (
	"log"
	"os/exec"
	"time"
)

func main() {
	log.Println("Start Logging Load Average")
	for {
		output, err := exec.Command("uptime").Output()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(output))
		time.Sleep(1 * time.Second)
	}
}
