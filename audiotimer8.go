package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

func main() {
	for {
		err := runProgram()
		if err != nil {
			log.Printf("Program failed with error: %v, restarting...\n", err)
			time.Sleep(1 * time.Second)
		}
	}
}

func runProgram() error {
	startTime := time.Now()
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			elapsed := time.Since(startTime)
			fileName := fmt.Sprintf("count_%d.txt", int(elapsed.Seconds()))
			err = ioutil.WriteFile(fileName, []byte(strconv.Itoa(int(elapsed.Seconds()))), 0644)
			if err != nil {
				log.Printf("Failed to create file: %v", err)
				return
			}
			time.Sleep(10 * time.Second)
		}
	}()

	cmd := exec.Command("go", "run", "C:\\Users\\computer\\Music\\audiodynamic08\\audiod8.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	wg.Wait()
	return err
}
